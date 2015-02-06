package physics

import "math"

// Circle is a struct of a circle with Radius float64.
type Circle struct {
	Radius float64 `json:"radius"`
}

// Intersect returns points of intersection between this and a Line.
func (circle *Circle) Intersect(l Line) []Vector {
	var intersections []Vector

	d := l.B.Sub(l.A)
	a := d.Dot(d)
	b := 2 * l.A.Dot(d)
	c := l.A.Dot(l.A) - circle.Radius*circle.Radius

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
