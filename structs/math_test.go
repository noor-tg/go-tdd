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

func TestArea(t *testing.T) {
	t.Run("calc area for rectangle", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		got := rect.Area()
		want := 100.0
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
	t.Run("calc area for circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
