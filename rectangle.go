package physics

// Rectangle struct implements Entity interface. It is a rectangle
// with two Vectors pointing the corners.
type Rectangle struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// Intersect return points of intersection between the Rectangle and a line.
func (r *Rectangle) Intersect(l Line) []Vector {
	var intersections []Vector
	xMin := -(r.Width / 2)
	xMax := r.Width / 2
	yMin := -(r.Height / 2)
	yMax := r.Height / 2
	bounds := [4]Line{
		Line{Vector{xMin, yMin}, Vector{xMax, yMin}},
		Line{Vector{xMin, yMax}, Vector{xMax, yMax}},
		Line{Vector{xMax, yMin}, Vector{xMax, yMax}},
		Line{Vector{xMin, yMin}, Vector{xMin, yMax}}}
	for _, bound := range bounds {
		intersections = append(intersections, l.Intersect(bound)...)
	}
	return intersections
}
