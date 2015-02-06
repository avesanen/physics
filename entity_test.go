package physics

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestObject(t *testing.T) {
	Convey("given a line between Vector{0,1} - Vector{2,1}", t, func() {
		//entity := &Entity{Shape: &Circle{Radius: 1}}
		entity := &Entity{Shape: &Rectangle{Width: 1, Height: 1}, Position: Vector{1, 1}}
		t.Log("testing...", entity.Intersect(Line{A: Vector{-2, -2}, B: Vector{2, 2}}))
		t.Log("testing...", entity.Intersect(Line{A: Vector{2, 2}, B: Vector{-2, -2}}))
	})
}
