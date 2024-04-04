package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("calc Perimeter for rectangle", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		got := Perimeter(rect)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {
	type Sample struct {
		shape Shape
		want  float64
	}
	// slice of struct
	// contains shape and want
	areaTests := []Sample{ // the test data
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	// iterate over test values
	for _, tt := range areaTests {
		// check for area
		checkArea(t, tt.shape, tt.want)
	}
}
