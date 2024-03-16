package simulation

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/RugiSerl/physics/app/Systems"
	m "github.com/RugiSerl/physics/app/math"
	"github.com/RugiSerl/physics/app/physicUnit"
	"github.com/RugiSerl/physics/app/values"
)

type Simulation struct {
	bodies []Systems.Body
	forces []physicUnit.Force2D
}

func Create() *Simulation {
	s := new(Simulation)
	mass := float64(1000)
	s.bodies = []Systems.Body{}
	s.forces = make([]m.Vec2, 100)
	for x := float64(0); x < 10; x++ {
		for y := float64(0); y < 10; y++ {
			s.bodies = append(s.bodies, Systems.Body{Mass: mass, Position: m.NewVec2(mass*rand.Float64(), mass*rand.Float64()), Speed: m.NewVec2(0, 0), Acceleration: m.NewVec2(0, 0)})
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
		b.Render(0.5)
		s.bodies[i] = updatePosition(b, s.forces[i])

		average = average.Add(s.bodies[i].Position)

	}
	if len(s.bodies) != 0 {
		average = average.Scale(1 / float64(len(s.bodies)))
		//rl.DrawCircleV(average.ToRL(), 10, rl.Red)
		fmt.Println(average)
	}

}

func (s *Simulation) ProvideDescription() string {
	return "Simulate the movements of bodies using the Newton formula"
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

func updatePosition(b Systems.Body, force physicUnit.Force2D) Systems.Body {
	b.Acceleration = force.Scale(1 / b.Mass)
	b.Speed = b.Speed.Add(b.Acceleration.Scale(values.Dt))
	b.Position = b.Position.Add(b.Speed.Scale(values.Dt))
	return b
}
