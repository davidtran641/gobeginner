package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Julia")
	want := "Hello, Julia"
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
