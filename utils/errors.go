package utils

import "testing"

func AssertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		// panic to stop test
		t.Fatal("did not get an error but wanted one")
	}
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Should not get Error")
	}
}
