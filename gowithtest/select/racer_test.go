package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	slowServer := makeDelayServer(20 * time.Millisecond)
	fastServer := makeDelayServer(0 * time.Millisecond)

	defer fastServer.Close()
	defer slowServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL

	got, err := Racer(slowURL, fastURL)

	if err != nil {
		t.Errorf("want no error, but got %v", err)
	}
	if got != want {
		t.Errorf("want %v, but got %v", want, got)
	}
}

func TestRacerTimeout(t *testing.T) {
	serverA := makeDelayServer(11 * time.Millisecond)
	serverB := makeDelayServer(12 * time.Millisecond)

	defer serverA.Close()
	defer serverB.Close()

	_, err := ConfigurableRacer(serverA.URL, serverB.URL, 2*time.Millisecond)
	if err == nil {
		t.Errorf("Expected an error but didn't")
	}
}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		rw.WriteHeader(http.StatusOK)
	}))
}
