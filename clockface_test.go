package main

import (
	"math"
	"testing"
	"time"

	clockface "github.com/forderation/go-test/clockface"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := clockface.Point{X: 150, Y: 150 - 90}
	got := clockface.SecondHand(tm)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}

}

func TestSecondsInRadians(t *testing.T) {

	cases := []struct {
		time  time.Time
		angle float64
	}{
		{clockface.SimpleTime(0, 0, 30), math.Pi},
		{clockface.SimpleTime(0, 0, 0), 0},
		{clockface.SimpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{clockface.SimpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clockface.SecondsInRadians(c.time)
			if got != c.angle {
				t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestSecondHandVector(t *testing.T) {
	cases := []struct {
		time  time.Time
		point clockface.Point
	}{
		{
			clockface.SimpleTime(0, 0, 30),
			clockface.Point{X: 0, Y: -1},
		},
		{
			clockface.SimpleTime(0, 0, 45),
			clockface.Point{X: -1, Y: 0},
		},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clockface.SecondHandPoint(c.time)
			if got != c.point {
				t.Fatalf("Wanted %v Point, but got %v", c.point, got)
			}
		})
	}
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
