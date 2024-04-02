package collections

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum total to 15", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		sum := Sum(input)

		expected := 15

		if !reflect.DeepEqual(sum, expected) {
			t.Errorf("got %v want %v", sum, expected)
		}
	})

	t.Run("sum total of with deffient input", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		sum := Sum(input)

		expected := 10

		if !reflect.DeepEqual(sum, expected) {
			t.Errorf("got %v want %v", sum, expected)
		}
	})
}

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{1, 2}, []int{3, 4}, []int{5, 6})
	expected := []int{3, 7, 11}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSum := func(t testing.TB, actual, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v want %v", actual, expected)
		}
	}

	t.Run("sum tails of slices correctly", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})

		expected := []int{5, 11}

		checkSum(t, actual, expected)

	})

	t.Run("sum tails of slices with one empty", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{4, 5, 6})

		expected := []int{0, 11}

		checkSum(t, actual, expected)
	})
}

func TestSumAllHeadNick(t *testing.T) {
	actual := SumAllHeadNick([]int{1, 2, 3}, []int{4, 5, 6})

	expected := []int{3, 9}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
}
