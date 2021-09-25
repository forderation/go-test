package main

import (
	"math"
	"testing"
)

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (t Triangle) Area() float64 {
	return (t.Height * t.Base) / 2
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

type Shape interface {
	Area() float64
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
	t.Run("Test Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("Test Circle", func(t *testing.T) {
		circle := Circle{Radius: 10}
		got := circle.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

}

type ComposeTestArea struct {
	Name    string
	Shape   Shape
	HasArea float64
}

func TestComposeArea(t *testing.T) {
	areaTest := []ComposeTestArea{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}
	for _, testObj := range areaTest {
		t.Run(testObj.Name, func(t *testing.T) {
			got := testObj.Shape.Area()
			t.Helper() // helper in this case in to help our debug know better location of error
			if got != testObj.HasArea {
				t.Errorf("got %.2f want %.2f", got, testObj.HasArea)
			}
		})
	}
}
