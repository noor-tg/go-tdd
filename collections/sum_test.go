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

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{1, 2}, []int{3, 4}, []int{5, 6})
	expected := []int{3, 7, 11}

	testza.AssertEqual(t, expected, actual)
}

func init() {
	testza.SetShowStartupMessage(false)
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum tails of slices correctly", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})

		expected := []int{5, 11}

		testza.AssertEqual(t, expected, actual)
	})

	t.Run("sum tails of slices with one empty", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{4, 5, 6})

		expected := []int{0, 11}

		testza.AssertEqual(t, expected, actual)
	})
}

func TestSumAllHeadNick(t *testing.T) {
	actual := SumAllHeadNick([]int{1, 2, 3}, []int{4, 5, 6})

	expected := []int{3, 9}

	testza.AssertEqual(t, expected, actual)
}
