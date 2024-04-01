package iterations

import (
	"testing"

	"github.com/MarvinJWendt/testza"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 'a' 5 times", func(t *testing.T) {

		out := Repeat("a")
		expected := "aaaaa"

		testza.AssertEqual(t, expected, out)
	})
	t.Run("repeat 'b' 5 times", func(t *testing.T) {

		out := Repeat("b")
		expected := "bbbbb"

		testza.AssertEqual(t, expected, out)
	})
}
