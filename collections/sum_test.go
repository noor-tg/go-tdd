package collections

import (
	"testing"

	"github.com/MarvinJWendt/testza"
)

func TestSum(t *testing.T) {

	t.Run("sum total to 15", func(t *testing.T) {
		input := [5]int{1, 2, 3, 4, 5}
		sum := Sum(input)

		expected := 15

		testza.AssertEqual(t, expected, sum)

	})

	t.Run("sum total to 10", func(t *testing.T) {
		input := [5]int{1, 2, 3, 4, 0}
		sum := Sum(input)

		expected := 10

		testza.AssertEqual(t, expected, sum)
	})
}

func init() {
	testza.SetShowStartupMessage(false)
}
