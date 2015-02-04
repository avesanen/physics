package physics

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLineIntersection(t *testing.T) {
	Convey("given a line between Vector{0,1} - Vector{2,1}", t, func() {
		line := Line{
			A: Vector{0, 1},
			B: Vector{2, 1},
		}

		Convey("and another between Vector{1,0} - Vector{1,2}", func() {
			line2 := Line{
				A: Vector{1, 0},
				B: Vector{1, 2},
			}
			Convey("they should intersect once at Vector{1,1}", func() {
				intersection := line.Intersect(line2)
				So(len(intersection), ShouldEqual, 1)
				So(intersection[0].X, ShouldEqual, 1)
				So(intersection[0].Y, ShouldEqual, 1)
				t.Log("Line intersections:", intersection)
			})
		})
	})
}
