package contexts

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {

	data := make(chan string, 1)

	go func() {
		var result string

		for _, c := range s.response {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(10 * time.Microsecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil

	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	want := "Hello world"
	store := &SpyStore{response: want}
	server := Server(store)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	if response.Body.String() != want {
		t.Errorf("want %s, but got %s", want, response.Body.String())
	}
}

func TestServerCancel(t *testing.T) {
	data := "Hello, world"
	store := &SpyStore{response: data}
	server := Server(store)

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	cancellingCtx, cancel := context.WithCancel(request.Context())
	time.AfterFunc(5*time.Microsecond, cancel)
	request = request.WithContext(cancellingCtx)

	response := &SpyResponseWriter{}

	server.ServeHTTP(response, request)

	if response.written {
		t.Errorf("a response should not have been written")
	}
}

// func assertWasNotCancelled(t *testing.T, store *SpyStore) {
// 	t.Helper()
// 	if store.cancelled {
// 		t.Errorf("Store should not have been cancelled")
// 	}
// }

// func assertWasCancelled(t *testing.T, store *SpyStore) {
// 	t.Helper()
// 	if !store.cancelled {
// 		t.Errorf("Store should be cancelled")
// 	}
// }
