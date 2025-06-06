package core

import "math"

type Point struct {
	X float64 // X coordinate of the point
	Y float64 // Y coordinate of the point
}

// NewPoint creates a new Point with the given coordinates.
func NewPoint(x, y float64) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

// Distance calculates the distance between two points.
func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(other.X-p.X, 2) + math.Pow(other.Y-p.Y, 2))
}

// Translate moves the point by the given delta values.
func (p *Point) Translate(dx, dy float64) {
	p.X += dx
	p.Y += dy
}
