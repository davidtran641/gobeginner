package main

import (
	"bytes"
	"os"
	"testing"
)

type MockSleeper struct {
	Calls int
}

func (s *MockSleeper) Sleep() {
	s.Calls++
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

func ExampleCountdown() {
	Countdown(os.Stdout, &MockSleeper{})
	// Output: 3
	// 2
	// 1
	// Go!
}
