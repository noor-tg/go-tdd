package collections

import (
	"testing"

	"github.com/MarvinJWendt/testza"
)

func TestSum(t *testing.T) {

	t.Run("sum total to 15", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		sum := Sum(input)

		expected := 15

		testza.AssertEqual(t, expected, sum)

	})

	t.Run("sum total of with deffient input", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		sum := Sum(input)

		expected := 10

		testza.AssertEqual(t, expected, sum)
	})
}

func init() {
	testza.SetShowStartupMessage(false)
}
