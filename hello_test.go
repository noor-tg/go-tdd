package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("noor")
	want := "Hello, noor"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
