package maps

import (
	"alnoor/gotdd/utils"
	"testing"
)

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

		utils.AssertError(t, err, NotFoundOnDectionary)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dectionary := Dectionary{}
		err := dectionary.Add("test", "this is test")
		want := "this is test"

		utils.AssertNoError(t, err)

		assertDefinition(t, dectionary, "test", want)
	})

	t.Run("add definition to existing word", func(t *testing.T) {
		dectionary := Dectionary{"test": "this is test"}
		err := dectionary.Add("test", "this is test")
		want := "this is test"

		assertDefinition(t, dectionary, "test", want)
		utils.AssertError(t, err, WordAlreadyExist)
	})

}

func assertDefinition(t testing.TB, dectionary Dectionary, word, want string) {
	t.Helper()
	got, err := dectionary.Search(word)

	if err != nil {
		t.Fatal("should find the word", err)
	}

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
