package springs

import (
	"fmt"

	"github.com/RugiSerl/physics/app/Systems"
	"github.com/RugiSerl/physics/app/camera"
	"github.com/RugiSerl/physics/app/math"
)

type SpringSimulation struct {
	spring Systems.Spring
	A, B   Systems.Body
}

func Create() *SpringSimulation {
	s := new(SpringSimulation)
	s.A = Systems.Body{Mass: 1}
	s.B = Systems.Body{Mass: 1, Position: math.NewVec2(400, 0)}
	s.spring = Systems.NewSpring(.1, &s.A, &s.B)

	return s
}

func (s SpringSimulation) Update(camera *camera.Camera2D) {
	s.spring.ApplyForce()
	s.A.UpdatePosition()
	s.B.UpdatePosition()
	s.A.Render()
	s.B.Render()
	fmt.Println(s.A.Position)

}

func (s *SpringSimulation) ProvideDescription() string {
	return "Messing around with springs"
}
