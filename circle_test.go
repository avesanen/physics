package physics

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCircleIntersection(t *testing.T) {
	Convey("given a line between Vector{0,1} - Vector{2,1}", t, func() {
		circle := Circle{
			Position: Vector{1, 1},
			Radius:   1,
		}

		Convey("and a line between Vector{1,0} - Vector{1,2}", func() {
			line := Line{
				A: Vector{0, 0},
				B: Vector{2, 2},
			}
			Convey("they should intersect once at Vector{1,1}", func() {
				intersection := circle.Intersect(line)
				t.Log("Circle intersections:", intersection)
				So(len(intersection), ShouldEqual, 2)
			})
		})
	})
}
