package physics

// Shape is an interface for shapes that can return intersection between it and a A-B line.
type Shape interface {
	Intersect(line Line) []Vector
}
