package math

import rl "github.com/gen2brain/raylib-go/raylib"

type Rect struct {
	Position, Size Vec2
}

func (r Rect) PointCollision(v Vec2) bool {
	return v.X > r.Position.X && v.X <= r.Position.X+r.Size.X && v.Y > r.Position.Y && v.Y <= r.Position.Y+r.Size.Y
}

func (r Rect) Draw(color rl.Color) {
	rl.DrawRectangleV(r.Position.ToRL(), r.Size.ToRL(), color)
}
