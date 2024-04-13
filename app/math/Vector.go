package math

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Vec2 struct {
	X, Y float64
}

func (v Vec2) ToRL() rl.Vector2 {
	return rl.NewVector2(float32(v.X), float32(v.Y))
}

func FromRL(vec rl.Vector2) Vec2 {
	return NewVec2(float64(vec.X), float64(vec.Y))
}

func NewVec2(x float64, y float64) Vec2 {
	return Vec2{X: x, Y: y}
}

// ---------------
// linear function
func (v Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{X: v.X + v2.X, Y: v.Y + v2.Y}
}

func (v Vec2) Substract(v2 Vec2) Vec2 {
	return Vec2{X: v.X - v2.X, Y: v.Y - v2.Y}
}

func (v Vec2) Scale(scalar float64) Vec2 {
	return Vec2{X: v.X * scalar, Y: v.Y * scalar}
}

func (v Vec2) GetNorm() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// null vector unhandled
func (v Vec2) Normalize() Vec2 {
	return v.Scale(1 / v.GetNorm())
}

// null vector unhandled
func (v Vec2) ScaleToNorm(norm float64) Vec2 {
	return v.Normalize().Scale(norm)
}

// ------------------
// rotation functions
// trigonometric circle
func (v Vec2) GetAngle() float64 {
	a := math.Acos(v.Normalize().X)
	if v.Y > 0 {
		return a
	} else {
		return -a
	}
}

func (v Vec2) Rotate(angle float64) Vec2 {
	x := math.Cos(angle)*v.X - math.Sin(angle)*v.Y
	y := math.Cos(angle)*v.Y + math.Sin(angle)*v.X
	return Vec2{X: x, Y: y}
}
