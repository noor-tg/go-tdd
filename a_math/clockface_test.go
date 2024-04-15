package a_math_test

import (
	"alnoor/gotdd/a_math"
	"fmt"
	"math"
	"testing"
	"time"
)

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := a_math.Point{X: 150, Y: 150 + 90}
	got := a_math.SecondHand(tm)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

func TestSocondHandPoint(t *testing.T) {
	testCases := []struct {
		time  time.Time
		point a_math.Point
	}{
		{simpleTime(0, 0, 30), a_math.Point{0, -1}},
		{simpleTime(0, 0, 45), a_math.Point{-1, 0}},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("convert time %v to Point %v", tC.time, tC.point), func(t *testing.T) {
			got := a_math.SecondHandPoint(tC.time)

			// floating point is not accurate in minimum values
			// so check for roughlyEqual value
			if !roughlyEqualPoint(got, tC.point) {
				t.Fatalf("Got %v, Want %v", got, tC.point)
			}
		})
	}

}

func roughlyEqualPoint(a, b a_math.Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalitythreshold = 1e-7
	return math.Abs(a-b) < equalitythreshold
}

// func TestSecondHandAt30Seconds(t *testing.T) {
// 	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)
// 	want := a_math.Point{150, 150 + 90}

// 	got := a_math.SecondHand(tm)

// 	if got != want {
// 		t.Errorf("Got %v, Want %v", got, want)
// 	}

// }

func TestSecondsInRadians(t *testing.T) {
	testCases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("convert seconds %v to angle %v", tC.time, tC.angle), func(t *testing.T) {
			got := a_math.SecondsInRadians(tC.time)

			if got != tC.angle {
				t.Fatalf("Wanted %v, Got %v", tC.angle, got)
			}
		})
	}

}

func simpleTime(i1, i2, i3 int) time.Time {
	return time.Date(312, time.October, 28, i1, i2, i3, 0, time.UTC)
}
