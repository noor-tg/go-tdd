package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const minutesLineLength = 80
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

func MinuteHand(tm time.Time) Point {
	p := MinuteHandPoint(tm)
	// scale :: to line length
	p = Point{p.X * minutesLineLength, p.Y * minutesLineLength}
	// flip
	p = Point{p.X, -p.Y}
	// translate: to make point coordinates start from center of clock
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}
	return p
}

func SecondHandPoint(time time.Time) Point {
	return angleToPoint(SecondsInRadians(time))
}

func MinuteHandPoint(time time.Time) Point {
	return angleToPoint(MinutesInRadians(time))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func SecondsInRadians(time time.Time) float64 {
	return math.Pi / (30 / float64(time.Second()))
}

func MinutesInRadians(t time.Time) float64 {
	return (SecondsInRadians(t) / 60) + (math.Pi / (30 / float64(t.Minute())))
}
