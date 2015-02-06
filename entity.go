package physics

import "time"
import "log"

// Entity is a type that has a velocity and a shape.
type Entity struct {
	Shape
	Id       string `json:"id"`
	Position Vector `json:"position"`
	Velocity Vector `json:"velocity"`
}

// NewEntity returns a new Entity with the shape provided.
func NewEntity(s Shape) *Entity {
	e := &Entity{}
	e.Shape = s
	return e
}

// Update moves the Entity according to passed time and velocity of the Entity.
func (e *Entity) Update(dt time.Duration) {
	e.Position.X += e.Velocity.X * dt.Seconds()
	e.Position.Y += e.Velocity.Y * dt.Seconds()
}

// Intersect returns a location vector of all points where given line intersects entity.
func (e *Entity) Intersect(line Line) []Vector {
	// Subtract entity position from line vector, as shape has no position.
	l := Line{
		A: line.A.Sub(e.Position),
		B: line.B.Sub(e.Position),
	}

	// Get intersections from shape, and add entity position to those.
	intersections := e.Shape.Intersect(l)
	for i, k := range intersections {
		intersections[i] = k.Add(e.Position)
	}

	// Sort intersections nearest to line.A first
	for i := 0; i < len(intersections)-1; i++ {
		for j := i + 1; j < len(intersections); j++ {
			v1 := intersections[i].Sub(line.A)
			v2 := intersections[j].Sub(line.A)
			if v1.Length() > v2.Length() {
				intersections[i], intersections[j] = intersections[j], intersections[i]
				log.Println("swap", intersections[i], intersections[j])
			} else {
				log.Println("no swap", intersections[i], intersections[j])
			}
		}
	}

	return intersections
}
