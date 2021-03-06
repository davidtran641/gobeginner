package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		actual := Hello("Julia")
		want := "Hello, Julia"
		assertMessage(t, want, actual)
	})

	t.Run("saying 'Hello, World' when empty name", func(t *testing.T) {
		assertMessage(t, "Hello, World", Hello(""))
	})

}

func assertMessage(t *testing.T, want, actual string) {

	if want != actual {
		t.Errorf("want %q, actual %q", want, actual)
	}
}
