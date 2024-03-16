package camera

import (
	"github.com/RugiSerl/physics/app/math"
	"github.com/RugiSerl/physics/app/values"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	CAMERA_ACCELERATION = 10
	CAMERA_ZOOM_AMOUNT  = 10
)

type Camera2D struct {
	rl.Camera2D
	speed     math.Vec2
	zoomSpeed float32
}

func NewCamera() *Camera2D {
	c := new(Camera2D)
	c.Offset = rl.NewVector2(0, 0)
	c.Target = rl.NewVector2(0, 0)
	c.speed = math.NewVec2(0, 0)
	c.zoomSpeed = 0

	return c

}

// met à jour la caméra pour visualiser le jeu et appliquer les transformations de cette dernière
func (c *Camera2D) UpdateCamera() {

	//déplacement éventuel de la caméra
	c.speed = c.speed.Scale(0.7)
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		c.speed.X -= CAMERA_ACCELERATION
	}
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		c.speed.X += CAMERA_ACCELERATION
	}
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		c.speed.Y -= CAMERA_ACCELERATION
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		c.speed.Y += CAMERA_ACCELERATION
	}
	// g.cameraMomentum estD la vitesse de la caméra, qui augmente lorsque l'utilisateur déplace la caméra, et diminue à chaque frame
	c.Target = math.FromRL(c.Target).Add(c.speed.Scale(values.Dt / float64(c.Zoom))).ToRL()

	//décalage de la caméra, pour que la cible, c'est-à-dire les coordonnées de la caméra, se trouve au milieu de l'écran
	c.Offset = rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2)

	//met à jour le zoom de la caméra

	c.zoomSpeed *= 0.8
	c.zoomSpeed += rl.GetMouseWheelMove() * CAMERA_ZOOM_AMOUNT * c.Zoom

	c.Zoom += c.zoomSpeed

	if c.Zoom < 1 { //1 est le minimum
		c.Zoom = 1
	}

}
