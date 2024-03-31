package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("hello to people", func(t *testing.T) {
		got := Hello("noor", "")
		want := "Hello, noor"

		assertStrEqual(t, got, want)
	})
	t.Run("say 'hello, world' when name is empty", func(t *testing.T) {

		got := Hello("", "")
		want := "Hello, World"

		assertStrEqual(t, got, want)
	})
	t.Run("hello in arabic", func(t *testing.T) {
		got := Hello("النور", "arabic")
		want := "مرحبا, النور"

		assertStrEqual(t, got, want)
	})
	t.Run("hello in french", func(t *testing.T) {
		got := Hello("noor", "french")
		want := "Bonjour, noor"

		assertStrEqual(t, got, want)
	})

}

func assertStrEqual(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
