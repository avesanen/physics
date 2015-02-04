package physics

import (
	"math"
)

// Vector is a struct X,Y float64 values and some mathematical methods
// for vector calculation
type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Length returns the length of vector in float64.
func (a *Vector) Length() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y)
}

// Normalize return a vector with same angle but distance of 1.
func (a *Vector) Normalize() Vector {
	d := a.Length()
	return Vector{a.X / d, a.Y / d}
}

// Cross return a cross product vector from a and b.
func (a *Vector) Cross(b Vector) Vector {
	return Vector{a.X * b.Y, a.Y * b.Y}
}

// Dot return the dot product of vector a and vector b as float64.
func (a *Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y
}

// Add return a vector with vector b added to vector a.
func (a *Vector) Add(b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y}
}

// Sub returns a vector with vector b subtracted from vector a.
func (a *Vector) Sub(b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y}
}

// Mul returns a vector with vector a multiplied by float64.
func (a *Vector) Mul(b float64) Vector {
	return Vector{a.X * b, a.Y * b}
}

// Div returns a vector with vector a divided with vector b.
func (a *Vector) Div(b float64) Vector {
	return Vector{a.X / b, a.Y / b}
}

// Mid returns a midpoint between vectors a and b as a vector.
func (a *Vector) Mid(b Vector) Vector {
	return Vector{(a.X + b.X) / 2, (a.Y + b.Y) / 2}
}

// MulVec returns a vector a*b
func (a Vector) MulVec(b Vector) Vector {
	return Vector{a.X * b.X, a.Y * b.Y}
}

// DivVec returns a vector a/b
func (a Vector) DivVec(b Vector) Vector {
	return Vector{a.X / b.X, a.Y / b.Y}
}
