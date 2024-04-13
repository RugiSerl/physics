package camera

import (
	"math"

	m "github.com/RugiSerl/physics/app/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	CAMERA_SPEED  = 1000
	ZOOM_AMOUNT   = 1
	CAMERA_SMOOTH = .05
	ZOOM_SMOOTH   = .05
)

type Camera2D struct {
	rl.Camera2D
	targetPosition  m.Vec2
	targetZoom      float32
	logarithmicZoom float32
}

func NewCamera() *Camera2D {
	c := new(Camera2D)
	c.Camera2D = rl.NewCamera2D(rl.NewVector2(0, 0), rl.NewVector2(0, 0), 0, 1)
	c.targetPosition = m.NewVec2(0, 0)
	c.targetZoom = 0

	return c
}

// met à jour la caméra pour visualiser le jeu et appliquer les transformations de cette dernière
func (c *Camera2D) UpdateCamera() {
	var speed float64 = CAMERA_SPEED * float64(rl.GetFrameTime()) / float64(c.Zoom)

	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		c.targetPosition.X -= speed
	}
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		c.targetPosition.X += speed
	}
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		c.targetPosition.Y -= speed
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		c.targetPosition.Y += speed
	}

	c.Target = m.FromRL(c.Target).Add(c.targetPosition.Substract(m.FromRL(c.Target)).Scale(float64(rl.GetFrameTime()) / CAMERA_SMOOTH)).ToRL()
	// décalage de la caméra, pour que la cible, c'est-à-dire les coordonnées de la caméra, se trouve au milieu de l'écran
	c.Offset = rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2)

	// met à jour le zoom de la caméra

	c.targetZoom += rl.GetMouseWheelMove()

	c.logarithmicZoom += (c.targetZoom - c.logarithmicZoom) / ZOOM_SMOOTH * rl.GetFrameTime()

	c.Zoom = float32(math.Exp(float64(c.logarithmicZoom)))
}

func (c *Camera2D) ConvertToWorldCoordinates(coordinates m.Vec2) m.Vec2 {
	return coordinates.Substract(m.FromRL(c.Offset)).Scale(1 / float64(c.Zoom)).Add(m.FromRL(c.Target))
}

func (c *Camera2D) Begin() {
	rl.BeginMode2D(c.Camera2D)
}

func (c *Camera2D) End() {
	rl.EndMode2D()
}
