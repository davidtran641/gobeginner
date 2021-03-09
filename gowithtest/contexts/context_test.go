package contexts

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep((10 * time.Millisecond))
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	data := "Hello world"
	store := &SpyStore{response: data}
	server := Server(store)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf("want %s, but got %s", response.Body.String(), data)
	}

	assertWasNotCancelled(t, store)

}

func TestServerFailed(t *testing.T) {
	data := "Hello, world"
	store := &SpyStore{response: data}
	server := Server(store)

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	cancellingCtx, cancel := context.WithCancel(request.Context())
	time.AfterFunc(5*time.Microsecond, cancel)
	request = request.WithContext(cancellingCtx)

	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	assertWasCancelled(t, store)
}

func assertWasNotCancelled(t *testing.T, store *SpyStore) {
	t.Helper()
	if store.cancelled {
		t.Errorf("Store should not have been cancelled")
	}
}

func assertWasCancelled(t *testing.T, store *SpyStore) {
	t.Helper()
	if !store.cancelled {
		t.Errorf("Store should be cancelled")
	}
}
