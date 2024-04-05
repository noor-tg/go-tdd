package maps

import "testing"

func TestSearch(t *testing.T) {
	dectionary := map[string]string{"test": "this is just a test"}

	got := Search(dectionary, "test")
	assertStrings(t, got)
}

func assertStrings(t testing.TB, got string) {
	want := "this is just a test"

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
