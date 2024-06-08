package Systems

import (
	m "github.com/RugiSerl/physics/app/math"
	"github.com/RugiSerl/physics/app/physicUnit"
	"github.com/RugiSerl/physics/app/values"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	BODY_SIZE = 10
)

type Body struct {
	Mass         float64
	Position     m.Vec2
	Speed        m.Vec2
	ForceApplied m.Vec2
}

func (b Body) Render() {
	c := 1 - 1/(b.Speed.GetNorm()/1000+1)

	rl.DrawCircleV(b.Position.ToRL(), BODY_SIZE, rl.NewColor(0, uint8(c*255), 0, 255))
}

func (b *Body) ApplyForce(force physicUnit.Force2D) {
	b.ForceApplied = b.ForceApplied.Add(force)
}

func (b *Body) UpdatePosition() {
	acceleration := b.ForceApplied.Scale(1 / b.Mass)
	b.Speed = b.Speed.Add(acceleration.Scale(values.Dt))
	b.Position = b.Position.Add(b.Speed.Scale(values.Dt))
	b.ForceApplied = m.Vec2{}
}

func (b *Body) Copy() Body {
	return *b
}
