package main

import "testing"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

func Area(r Rectangle) float64 {
	return r.Width * r.Height
}

// test function are below

func TestPerimeter(t *testing.T) {
	rectange := Rectangle{10.0, 10.0}
	got := Perimeter(rectange)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Area(rectangle)
	want := 100.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
