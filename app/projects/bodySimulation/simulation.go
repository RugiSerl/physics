package simulation

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"github.com/RugiSerl/physics/app/Systems"
	"github.com/RugiSerl/physics/app/camera"
	m "github.com/RugiSerl/physics/app/math"
	"github.com/RugiSerl/physics/app/physicUnit"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	BODY_MASS = float64(1000)
)

type Simulation struct {
	bodies []Systems.Body
	forces []physicUnit.Force2D
}

func Create() *Simulation {
	s := new(Simulation)
	s.bodies = []Systems.Body{}
	s.forces = []physicUnit.Force2D{}

	return s
}

func (s *Simulation) Update(camera *camera.Camera2D) {
	for i, b := range s.bodies {
		s.forces[i] = s.updateForces(b)

	}
	var average m.Vec2 = m.NewVec2(0, 0)
	for i, b := range s.bodies {
		b.UpdatePosition(s.forces[i])
		s.bodies[i] = b.Copy()
		b.Render(0.5)

		average = average.Add(s.bodies[i].Position)

	}
	if len(s.bodies) != 0 {
		average = average.Scale(1 / float64(len(s.bodies)))
		rl.DrawCircleV(average.Scale(0.5).ToRL(), 10, rl.Red)
		fmt.Println(average)
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) || rl.IsKeyDown(rl.KeyLeftShift) {
		s.spawnBody(camera.ConvertToWorldCoordinates(m.FromRL(rl.GetMousePosition())).Scale(2), BODY_MASS)
	} else if rl.IsKeyPressed(rl.KeySpace) {
		s.spawnMany(camera.ConvertToWorldCoordinates(m.FromRL(rl.GetMousePosition())).Scale(2))

	}

	//rl.DrawRectangleV(camera.ConvertToWorldCoordinates(m.FromRL(rl.GetMousePosition())).ToRL(), rl.NewVector2(20, 20), rl.Red)

}

func (s *Simulation) spawnMany(position m.Vec2) {
	for x := float64(-6); x < 6; x++ {
		for y := float64(-6); y < 6; y++ {
			s.spawnBody(m.NewVec2(BODY_MASS*(rand.Float64()-.5), BODY_MASS*(rand.Float64()-.5)).Add(position), BODY_MASS)
		}
	}
}

func (s *Simulation) spawnBody(position m.Vec2, mass float64) {
	s.bodies = append(s.bodies, Systems.Body{Mass: mass, Position: position, Speed: m.NewVec2(0, 0), Acceleration: m.NewVec2(0, 0)})
	s.forces = append(s.forces, m.Vec2{0, 0})
}

func (s *Simulation) ProvideDescription() string {
	return "Simulate the movements of bodies using the Newton formula\nCurrently " + strconv.Itoa(len(s.bodies)) + " bodies"
}

func (s *Simulation) updateForces(b Systems.Body) physicUnit.Force2D {
	var force physicUnit.Force2D = m.NewVec2(0, 0)
	for _, otherBody := range s.bodies {
		if otherBody != b {
			vector := otherBody.Position.Substract(b.Position)
			attraction := vector.Scale(physicUnit.G * otherBody.Mass * b.Mass / math.Pow(vector.GetNorm(), 2))

			force = force.Add(attraction)
		}
	}

	return force
}
