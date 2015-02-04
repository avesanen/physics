package physics

// Line is a struct of a finite line between points a and b
type Line struct {
	A Vector `json:"a"`
	B Vector `json:"b"`
}

// Interect returns the point of intersection another line
func (l1 *Line) Intersect(l2 Line) []Vector {
	var intersections []Vector
	d := (l1.A.X-l1.B.X)*(l2.A.Y-l2.B.Y) - (l1.A.Y-l1.B.Y)*(l2.A.X-l2.B.X)
	if d != 0 {
		xi := ((l2.A.X-l2.B.X)*(l1.A.X*l1.B.Y-l1.A.Y*l1.B.X) - (l1.A.X-l1.B.X)*(l2.A.X*l2.B.Y-l2.A.Y*l2.B.X)) / d
		yi := ((l2.A.Y-l2.B.Y)*(l1.A.X*l1.B.Y-l1.A.Y*l1.B.X) - (l1.A.Y-l1.B.Y)*(l2.A.X*l2.B.Y-l2.A.Y*l2.B.X)) / d
		if !((l1.A.X < xi && l1.B.X < xi) || (l1.A.X > xi && l1.B.X > xi) ||
			(l1.A.Y < yi && l1.B.Y < yi) || (l1.A.Y > yi && l1.B.Y > yi) ||
			(l2.A.X < xi && l2.B.X < xi) || (l2.A.X > xi && l2.B.X > xi) ||
			(l2.A.Y < yi && l2.B.Y < yi) || (l2.A.Y > yi && l2.B.Y > yi)) {
			intersections = append(intersections, Vector{xi, yi})
		}
	}
	return intersections
}
