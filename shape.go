package physics

type Shape interface {
	Intersect(Line) []Vector
	GetPosition() Vector
	SetPosition(Vector)
}
