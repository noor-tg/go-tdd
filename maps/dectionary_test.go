package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dectionary := Dectionary{"test": "this is just a test"}

		got, err := dectionary.Search("test")
		want := "this is just a test"
		if err != nil {
			t.Fatal("should not return error")
		}
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		dectionary := Dectionary{"test": "this is just a test"}

		_, err := dectionary.Search("utest")

		if err == nil {
			t.Fatal("expected to get error")
		}

		assertStrings(t, err.Error(), NotFoundOnDectionary.Error())
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
