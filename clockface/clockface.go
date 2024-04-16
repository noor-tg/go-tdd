package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const secondsLineLength = 90
const clockCenterX = 150
const clockCenterY = 150

func SecondHand(tm time.Time) Point {
	p := SecondHandPoint(tm)
	// scale :: to line length
	p = Point{p.X * secondsLineLength, p.Y * secondsLineLength}
	// flip
	p = Point{p.X, -p.Y}
	// translate: to make point coordinates start from center of clock
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}
	return p
}

func SecondHandPoint(time time.Time) Point {
	angle := SecondsInRadians(time)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func SecondsInRadians(time time.Time) float64 {
	return math.Pi / (30 / float64(time.Second()))
}
