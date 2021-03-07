package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

type MockSleeper struct {
	Calls int
}

func (s *MockSleeper) Sleep() {
	s.Calls++
}

type CountdownOperationSpy struct {
	Calls []string
}

const (
	sleep = "sleep"
	write = "write"
)

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &MockSleeper{}

	Countdown(buffer, sleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("want %v, but got %v", want, got)
	}

	if sleeper.Calls != 4 {
		t.Errorf("want %d, but got %v", 4, sleeper.Calls)
	}
}

func TestCountdownSpy(t *testing.T) {
	operationSpy := &CountdownOperationSpy{}
	Countdown(operationSpy, operationSpy)

	want := []string{
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}

	if !reflect.DeepEqual(want, operationSpy.Calls) {
		t.Errorf("wanted calls %v, but got %v", want, operationSpy.Calls)
	}
}

func ExampleCountdown() {
	Countdown(os.Stdout, &MockSleeper{})
	// Output: 3
	// 2
	// 1
	// Go!
}
