package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to names", func(t *testing.T) {
		got := Hello("Alex", "English")
		want := "Hello, Alex"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello world if there is no name", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Greek", func(t *testing.T) {
		got := Hello("Alex", "Greek")
		want := "Geia sou Alex"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Alex", "French")
		want := "Bonjour, Alex"

		assertCorrectMessage(t, got, want)
	})

}
