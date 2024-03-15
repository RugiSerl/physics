package Systems

import (
	m "github.com/RugiSerl/physics/app/math"
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

func (b Body) Render() {
	rl.DrawCircleV(b.Position.ToRL(), BODY_SIZE, rl.Blue)
}
