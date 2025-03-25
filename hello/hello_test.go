package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people in spanish", func(t *testing.T) {
		got := Hello("Aditya", "Spanish")
		want := "Hola, Aditya"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in french", func(t *testing.T) {
		got := Hello("Aditya", "French")
		want := "Bonjour, Aditya"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people without a name", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // t.Helper() is used to assert that this is a helper function for testing
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
