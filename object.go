package physics

import "time"

type Object struct {
	Shape
	Velocity Vector `json:"velocity"`
}

func (o *Object) Update(dt time.Duration) {
	x := o.Shape.GetPosition().X + o.Velocity.X*dt.Seconds()
	y := o.Shape.GetPosition().Y + o.Velocity.Y*dt.Seconds()
	o.SetPosition(Vector{x, y})
}
