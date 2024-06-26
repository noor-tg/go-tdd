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

func SecondHandPoint(time time.Time) Point {
	return angleToPoint(SecondsInRadians(time))
}

func MinuteHandPoint(time time.Time) Point {
	return angleToPoint(MinutesInRadians(time))
}

func HourHandPoint(time time.Time) Point {
	return angleToPoint(HoursInRadians(time))
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

func HoursInRadians(t time.Time) float64 {
	return (MinutesInRadians(t) / 12) + (math.Pi / (6 / float64(t.Hour()%12)))
}
