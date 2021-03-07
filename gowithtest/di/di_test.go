package main

import (
	"bytes"
	"os"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Julia")

	got := buffer.String()
	want := "Hello, Julia"

	if got != want {
		t.Errorf("want %q, but got %q", want, got)
	}
}

func ExampleGreet() {
	Greet(os.Stdout, "Julia")
	// Output: Hello, Julia
}
