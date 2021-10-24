package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	return Point{}
}

func SecondsInRadians(t time.Time) float64 {
	return math.Pi
}

func SecondHandPoint(t time.Time) Point {
	angle := SecondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func SimpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(123, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func RoughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}
