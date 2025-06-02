package math

import (
	"fmt"
	"math"
)

// Vec2 represents a 2D vector with X and Y coordinates.
type Vec2 struct {
	X float64 // X coordinate
	Y float64 // Y coordinate
}

// NewVec2 creates a new Vec2 with the given X and Y coordinates.
func NewVec2(x, y float64) *Vec2 {
	return &Vec2{
		X: x,
		Y: y,
	}
}

// Add adds two Vec2 vectors and returns the result.
// NOTE: This does NOT modify the receiver; you must assign the result:
//   v = v.Add(other)
func (v *Vec2) Add(other *Vec2) *Vec2 {
	return &Vec2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

// Subtract subtracts another Vec2 from this Vec2 and returns the result.
// NOTE: This does NOT modify the receiver; you must assign the result.
func (v *Vec2) Subtract(other *Vec2) *Vec2 {
	return &Vec2{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

// Scale scales the Vec2 by a given factor and returns the result.
// NOTE: This does NOT modify the receiver; you must assign the result.
func (v *Vec2) Scale(factor float64) *Vec2 {
	return &Vec2{
		X: v.X * factor,
		Y: v.Y * factor,
	}
}

// Magnitude returns the length of the Vec2 vector.
func (v *Vec2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalize returns a unit vector in the same direction as the Vec2.
func (v *Vec2) Normalize() *Vec2 {
	mag := v.Magnitude()
	if mag == 0 {
		return &Vec2{X: 0, Y: 0} // Avoid division by zero
	}
	return &Vec2{
		X: v.X / mag,
		Y: v.Y / mag,
	}
}

// Dot returns the dot product of this Vec2 and another Vec2.
func (v *Vec2) Dot(other *Vec2) float64 {
	return v.X*other.X + v.Y*other.Y
}

// Distance returns the distance between this Vec2 and another Vec2.
func (v *Vec2) Distance(other *Vec2) float64 {
	return math.Sqrt((v.X-other.X)*(v.X-other.X) + (v.Y-other.Y)*(v.Y-other.Y))
}

// Angle returns the angle in radians between this Vec2 and the positive X-axis.
func (v *Vec2) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

// Rotate rotates the Vec2 by a given angle in radians and returns the result.
func (v *Vec2) Rotate(angle float64) *Vec2 {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return &Vec2{
		X: v.X*cos - v.Y*sin,
		Y: v.X*sin + v.Y*cos,
	}
}

// String returns a string representation of the Vec2.
func (v *Vec2) String() string {
	return "(" + fmt.Sprintf("%.2f", v.X) + ", " + fmt.Sprintf("%.2f", v.Y) + ")"
}

// Equals checks if two Vec2 vectors are equal.
func (v *Vec2) Equals(other *Vec2) bool {
	return v.X == other.X && v.Y == other.Y
}

// Zero returns a zero vector (0, 0).
func Zero() *Vec2 {
	return &Vec2{X: 0, Y: 0}
}

// One returns a unit vector (1, 1).
func One() *Vec2 {
	return &Vec2{X: 1, Y: 1}
}

// Up returns a vector pointing upwards (0, -1).
func Up() *Vec2 {
	return &Vec2{X: 0, Y: -1}
}

// Down returns a vector pointing downwards (0, 1).
func Down() *Vec2 {
	return &Vec2{X: 0, Y: 1}
}

// Left returns a vector pointing to the left (-1, 0).
func Left() *Vec2 {
	return &Vec2{X: -1, Y: 0}
}

// Right returns a vector pointing to the right (1, 0).
func Right() *Vec2 {
	return &Vec2{X: 1, Y: 0}
}
