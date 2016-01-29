package main

import (
	"math"
	"testing"
)

var perimeterTests = []struct {
	x, y, r, out float64
}{
	{1, 1, 1, 2 * math.Pi},
	{1, 1, 3, 6 * math.Pi},
}

func TestPerimeter(t *testing.T) {
	for _, tt := range perimeterTests {
		c := Circle{x: tt.x, y: tt.y, r: tt.r}
		p := c.perimeter()

		if p != tt.out {
			t.Errorf("perimeter(%f) => %f, want %f", tt.r, p, tt.out)
		}
	}
}

var areaTests = []struct {
	x, y, r, out float64
}{
	{1, 1, 1, math.Pi},
}

func TestArea(t *testing.T) {
	for _, tt := range areaTests {
		c := Circle{x: tt.x, y: tt.y, r: tt.r}
		a := c.area()

		if a != tt.out {
			t.Errorf("perimeter(%f) => %f, want %f", tt.r, a, tt.out)
		}
	}
}
