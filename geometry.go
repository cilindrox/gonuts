package main

import "math"

func distance(x1, y1, x2, y2 float64) float64 {
	a := x1 - x2
	b := y1 - y2
	return math.Sqrt(a*a + b*b)
}

// Shape interface contains geometry methods
type Shape interface {
	area() float64
	perimeter() float64
}

// Circle Shape type
type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// Rectangle Shape type
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	return r.x1 * r.x2 * r.y1 * r.y2
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (distance(r.x1, 0, r.x2, 0) + distance(0, r.y1, 0, r.y2))
}
