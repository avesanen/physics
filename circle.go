package physics

import "math"

// Circle is a struct of a circle with Position Vector and Radius float64.
type Circle struct {
	Position Vector  `json:"position"`
	Radius   float64 `json:"radius"`
}

// GetPosition returns Circle's position as a vector
func (c *Circle) GetPosition() Vector {
	return c.Position
}

// SetPosition sets Circle's position to given vector.
func (c *Circle) SetPosition(p Vector) {
	c.Position = p
}

// Intersect returns points of intersection between this and a Line.
func (circle *Circle) Intersect(l Line) []Vector {
	var intersections []Vector

	d := l.B.Sub(l.A)
	f := l.A.Sub(circle.Position)
	a := d.Dot(d)
	b := 2 * f.Dot(d)
	c := f.Dot(f) - circle.Radius*circle.Radius

	disc := b*b - 4*a*c
	if disc >= 0 {
		disc = math.Sqrt(disc)

		t1 := (-b - disc) / (2 * a)
		t2 := (-b + disc) / (2 * a)

		c1 := Vector{l.A.X + d.X*t1, l.A.Y + d.Y*t1}
		c2 := Vector{l.A.X + d.X*t2, l.A.Y + d.Y*t2}

		if t1 >= 0 && t1 <= 1 {
			intersections = append(intersections, c1)
		}

		if t2 >= 0 && t2 <= 1 {
			intersections = append(intersections, c2)
		}
	}
	return intersections
}
