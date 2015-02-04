package physics

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRectangleIntersection(t *testing.T) {
	Convey("given a rectangle at Vector{1.5, 1.5} with width and height of 1.0", t, func() {
		rect := Rectangle{
			Position: Vector{1.5, 1.5},
			Width:    1.0,
			Height:   1.0,
		}

		Convey("and another between Vector{1,0} - Vector{1,2}", func() {
			line2 := Line{
				A: Vector{1, 0},
				B: Vector{1, 2},
			}

			Convey("they should intersect twice at Vector{1,1} and Vector{1,2}", func() {
				intersection := rect.Intersect(line2)
				So(len(intersection), ShouldEqual, 2)
				t.Log("Rectangle intersections:", intersection)
			})
		})
	})
}
