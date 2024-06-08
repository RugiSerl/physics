package projects

import (
	"github.com/RugiSerl/physics/app/camera"
	simulation "github.com/RugiSerl/physics/app/projects/bodySimulation"
	simulationOptimised "github.com/RugiSerl/physics/app/projects/bodySimulationOptimised"
	"github.com/RugiSerl/physics/app/projects/springs"
)

type Project interface {
	Update(*camera.Camera2D)
	ProvideDescription() string
}

type projectType int

const (
	PROJECT_BODY_SIMULATION projectType = iota
	PROJECT_BODY_SIMULATION_OPTIMISED
	PROJECT_SPRING
)

func NewProject(project projectType) Project {
	var p Project
	switch project {
	case PROJECT_BODY_SIMULATION:
		p = simulation.Create()
	case PROJECT_BODY_SIMULATION_OPTIMISED:
		p = simulationOptimised.Create()
	case PROJECT_SPRING:
		p = springs.Create()
	}
	return p

}
