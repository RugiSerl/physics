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
	Acceleration m.Vec2
}

func (b Body) Render(scale float64) {
	c := 1 - 1/(b.Speed.GetNorm()/1000+1)

	rl.DrawCircleV(b.Position.Scale(scale).ToRL(), BODY_SIZE*float32(scale), rl.NewColor(0, uint8(c*255), 0, 255))
}

func (b *Body) UpdatePosition(force physicUnit.Force2D) {
	b.Acceleration = force.Scale(1 / b.Mass)
	b.Speed = b.Speed.Add(b.Acceleration.Scale(values.Dt))
	b.Position = b.Position.Add(b.Speed.Scale(values.Dt))
}

func (b *Body) Copy() Body {
	return *b
}
