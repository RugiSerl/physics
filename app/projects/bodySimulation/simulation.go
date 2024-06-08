package simulation

import (
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
	BODY_MASS = float64(100000000000000)
)

type Simulation struct {
	bodies []Systems.Body
	forces []physicUnit.Force2D
}

func Create() *Simulation {
	s := new(Simulation)
	s.bodies = []Systems.Body{}
	s.forces = []physicUnit.Force2D{}

	// s.spawnBody(m.Vec2{0, 0}, BODY_MASS*1000)
	// s.bodies[0].Speed = m.Vec2{0, 0}
	return s
}

func (s *Simulation) Update(camera *camera.Camera2D) {
	for i, b := range s.bodies {
		s.forces[i] = s.updateForces(b)
	}
	for i, b := range s.bodies {
		b.ApplyForce(s.forces[i])
		b.UpdatePosition()
		s.bodies[i] = b.Copy()
		b.Render()

	}

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) || rl.IsKeyDown(rl.KeyLeftShift) {
		s.spawnBody(camera.ConvertToWorldCoordinates(m.FromRL(rl.GetMousePosition())), BODY_MASS)
	} else if rl.IsKeyPressed(rl.KeySpace) {
		s.spawnMany(camera.ConvertToWorldCoordinates(m.FromRL(rl.GetMousePosition())))
	}
}

func (s *Simulation) spawnMany(position m.Vec2) {
	for x := float64(-6); x < 6; x++ {
		for y := float64(-6); y < 6; y++ {
			s.spawnBody(m.NewVec2((rand.Float64()-.5), (rand.Float64()-.5)).Add(position), BODY_MASS)
		}
	}
}

func (s *Simulation) spawnBody(position m.Vec2, mass float64) {
	s.bodies = append(s.bodies, Systems.Body{Mass: mass, Position: position, Speed: m.NewVec2(0, 0), ForceApplied: m.NewVec2(0, 0)})
	s.forces = append(s.forces, m.Vec2{0, 0})
}

func (s *Simulation) ProvideDescription() string {
	return "Simulate the movements of bodies using the Newton formula\nCurrently " + strconv.Itoa(len(s.bodies)) + " bodies"
}

func (s *Simulation) updateForces(b Systems.Body) physicUnit.Force2D {
	var force physicUnit.Force2D = m.NewVec2(0, 0)
	for _, otherBody := range s.bodies {
		if otherBody != b {
			vector := otherBody.Position.Substract(b.Position).Normalize()
			attraction := vector.Scale(physicUnit.G * otherBody.Mass * b.Mass / math.Pow(vector.GetNorm(), 2))

			force = force.Add(attraction)
		}
	}

	return force
}
