package simulation

import (
	"math"
	"math/rand"

	"github.com/RugiSerl/physics/app/Systems"
	m "github.com/RugiSerl/physics/app/math"
	"github.com/RugiSerl/physics/app/physicUnit"
	"github.com/RugiSerl/physics/app/values"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Simulation struct {
	bodies []Systems.Body
	forces []physicUnit.Force2D
}

func Create() *Simulation {
	s := new(Simulation)

	s.bodies = []Systems.Body{}
	s.forces = make([]m.Vec2, 100)
	for x := float64(0); x < 10; x++ {
		for y := float64(0); y < 10; y++ {
			s.bodies = append(s.bodies, Systems.Body{Mass: 100, Position: m.NewVec2(300+x*100*rand.Float64(), 100+y*100*rand.Float64()), Speed: m.NewVec2(0, 0), Acceleration: m.NewVec2(0, 0)})
		}
	}

	return s
}

func (s *Simulation) Update() {
	for i, b := range s.bodies {
		s.forces[i] = s.updateForces(b)

	}
	var average m.Vec2 = m.NewVec2(0, 0)
	for i, b := range s.bodies {
		s.bodies[i] = updatePosition(b, s.forces[i])
		b.Render()
		average = average.Add(s.bodies[i].Position)
	}
	if len(s.bodies) != 0 {
		average = average.Scale(1 / float64(len(s.bodies)))
		rl.DrawCircleV(average.ToRL(), 10, rl.Red)
	}

}

func (s *Simulation) ProvideDescription() string {
	return "Simulate the movements of bodies using the Newton formula"
}

func (s *Simulation) updateForces(b Systems.Body) physicUnit.Force2D {
	var force physicUnit.Force2D = m.NewVec2(0, 0)
	for _, e := range s.bodies {
		if e != b {
			vector := e.Position.Substract(b.Position)
			attraction := vector.Scale(physicUnit.G * e.Mass * b.Mass / math.Pow(vector.GetNorm(), 2))

			force = force.Add(attraction)
		}
	}

	return force
}

func updatePosition(b Systems.Body, force physicUnit.Force2D) Systems.Body {
	b.Acceleration = force.Scale(1 / b.Mass)
	b.Speed = b.Speed.Add(b.Acceleration.Scale(values.Dt))
	b.Position = b.Position.Add(b.Speed.Scale(values.Dt))

	return b
}
