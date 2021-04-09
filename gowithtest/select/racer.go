package racer

import (
	"fmt"
	"net/http"
	"time"
)

const (
	defaultTimeout = 10 * time.Second
)

// Racer return a ConfigurableRacer with default timeout
func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, defaultTimeout)
}

// ConfigurableRacer return the fastest result from a and b
func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}

}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		defer close(ch)

		http.Get(url)
	}()
	return ch
}
