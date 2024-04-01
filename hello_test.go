package main

import (
	"testing"

	"github.com/MarvinJWendt/testza"
)

func TestHello(t *testing.T) {
	t.Run("hello to people", func(t *testing.T) {
		got := Hello("noor", "")
		want := "Hello, noor"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'hello, world' when name is empty", func(t *testing.T) {

		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("hello in arabic", func(t *testing.T) {
		got := Hello("النور", "arabic")
		want := "مرحبا, النور"

		assertCorrectMessage(t, got, want)
	})

	t.Run("hello in french", func(t *testing.T) {
		got := Hello("noor", "french")
		want := "Bonjour, noor"

		assertCorrectMessage(t, got, want)
	})

	t.Run("hello in spanish", func(t *testing.T) {
		got := Hello("noor", "spanish")
		want := "Hola, noor"

		assertCorrectMessage(t, got, want)
	})

}
func init() {
	testza.SetShowStartupMessage(false) // Disable the startup message
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		testza.AssertEqual(t, want, got)
		// t.Errorf("got %q want %q", got, want)
	}
}
