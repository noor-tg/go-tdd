package clockface_test

import (
	"alnoor/gotdd/clockface"
	"fmt"
	"math"
	"testing"
	"time"
)

func TestSocondHandPoint(t *testing.T) {
	testCases := []struct {
		time  time.Time
		point clockface.Point
	}{
		{simpleTime(0, 0, 30), clockface.Point{0, -1}},
		{simpleTime(0, 0, 45), clockface.Point{-1, 0}},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("convert time %v to Point %v", tC.time, tC.point), func(t *testing.T) {
			got := clockface.SecondHandPoint(tC.time)

			// floating point is not accurate in minimum values
			// so check for roughlyEqual value
			if !roughlyEqualPoint(got, tC.point) {
				t.Fatalf("Got %v, Want %v", got, tC.point)
			}
		})
	}

}

func TestMinuteHandPoint(t *testing.T) {
	testCases := []struct {
		time  time.Time
		point clockface.Point
	}{
		{simpleTime(0, 30, 0), clockface.Point{0, -1}},
		{simpleTime(0, 45, 0), clockface.Point{-1, 0}},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("convert time %v to Point %v", tC.time, tC.point), func(t *testing.T) {
			got := clockface.MinuteHandPoint(tC.time)

			// floating point is not accurate in minimum values
			// so check for roughlyEqual value
			if !roughlyEqualPoint(got, tC.point) {
				t.Fatalf("Got %v, Want %v", got, tC.point)
			}
		})
	}

}
func TestHourHandPoint(t *testing.T) {
	testCases := []struct {
		time  time.Time
		point clockface.Point
	}{
		{simpleTime(6, 0, 0), clockface.Point{0, -1}},
		{simpleTime(21, 0, 0), clockface.Point{-1, 0}},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("convert time %v to Point %v", tC.time, tC.point), func(t *testing.T) {
			got := clockface.HourHandPoint(tC.time)

			// floating point is not accurate in minimum values
			// so check for roughlyEqual value
			if !roughlyEqualPoint(got, tC.point) {
				t.Fatalf("Got %v, Want %v", got, tC.point)
			}
		})
	}

}

func roughlyEqualPoint(a, b clockface.Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalitythreshold = 1e-7
	return math.Abs(a-b) < equalitythreshold
}

// func TestSecondHandAt30Seconds(t *testing.T) {
// 	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)
// 	want := clockface.Point{150, 150 + 90}

// 	got := clockface.SecondHand(tm)

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
			got := clockface.SecondsInRadians(tC.time)

			if got != tC.angle {
				t.Fatalf("Wanted %v, Got %v", tC.angle, got)
			}
		})
	}

}

func simpleTime(i1, i2, i3 int) time.Time {
	return time.Date(312, time.October, 28, i1, i2, i3, 0, time.UTC)
}

func TestMinutesInRadians(t *testing.T) {
	testCases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("convert seconds %v to angle %v", tC.time, tC.angle), func(t *testing.T) {
			got := clockface.MinutesInRadians(tC.time)

			if got != tC.angle {
				t.Fatalf("Wanted %v, Got %v", tC.angle, got)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, tC := range cases {
		t.Run(fmt.Sprintf("convert hours %v to angle %v", tC.time, tC.angle), func(t *testing.T) {
			got := clockface.HoursInRadians(tC.time)
			if !roughlyEqualFloat64(got, tC.angle) {
				t.Fatalf("Wanted %v radians, but got %v", tC.angle, got)
			}
		})
	}
}
