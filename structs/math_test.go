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
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("calc area for rectangle", func(t *testing.T) {
		shape := Rectangle{10.0, 10.0}
		checkArea(t, shape, 100.0)
	})

	t.Run("calc area for circle", func(t *testing.T) {
		shape := Circle{10.0}
		checkArea(t, shape, 314.1592653589793)
	})
}
