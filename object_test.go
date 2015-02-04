package physics

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestObject(t *testing.T) {
	Convey("given a line between Vector{0,1} - Vector{2,1}", t, func() {
		obj := Object{}

		obj.Shape = Shape(&Circle{
			Radius: 1.0,
		})

		obj.SetPosition(Vector{1.5, 1.5})

		line := Line{
			A: Vector{0, 0},
			B: Vector{3, 3},
		}

		t.Log(obj.Intersect(line))

		obj.Shape = Shape(&Rectangle{
			Position: Vector{1.5, 1.5},
			Width:    1,
			Height:   1,
		})

		t.Log(obj.Intersect(line))

	})
}
