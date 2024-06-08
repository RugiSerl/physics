package simulationOptimised

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/RugiSerl/physics/app/Systems"
	"github.com/RugiSerl/physics/app/camera"
	m "github.com/RugiSerl/physics/app/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	BODY_MASS = float64(1000000000000)
)

type Simulation struct {
	bodies []*Systems.Body
	tree   *QuadTree
}

func Create() *Simulation {
	s := new(Simulation)
	s.bodies = []*Systems.Body{}
	s.tree = NewQuadTree(m.Rect{Position: m.Vec2{X: -TREE_BORDER_SIZE / 2, Y: -TREE_BORDER_SIZE / 2}, Size: m.Vec2{X: TREE_BORDER_SIZE, Y: TREE_BORDER_SIZE}})
	return s
}

func (s *Simulation) Update(camera *camera.Camera2D) {
	s.tree.ShowRegion()
	s.tree.UpdateMass()
	s.tree.UpdateForcesNormal(s.tree.ListChildren())
	s.tree.UpdatePositions(s.tree)
	for _, b := range s.bodies {
		// b.UpdatePosition(s.forces[i])
		// s.bodies[i] = b
		b.Render()
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) || rl.IsKeyDown(rl.KeyLeftShift) {
		s.spawnBody(camera.ConvertToWorldCoordinates(m.FromRL(rl.GetMousePosition())), BODY_MASS)
	} else if rl.IsKeyPressed(rl.KeySpace) {
		s.spawnMany(camera.ConvertToWorldCoordinates(m.FromRL(rl.GetMousePosition())))
	}

	if rl.IsKeyPressed(rl.KeyA) {
		s.tree.PrintTree()
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
	fmt.Println("-------------------------------\nAdding..")

	s.bodies = append(s.bodies, &Systems.Body{Mass: mass, Position: position, Speed: m.NewVec2(0, 0), ForceApplied: m.NewVec2(0, 0)})
	s.tree.Insert(s.bodies[len(s.bodies)-1]) // add body we created
	fmt.Println("Amount of Leafs :", s.tree.AmountOfLeaf())
}

func (s *Simulation) ProvideDescription() string {
	return "Simulate the movements of bodies using the Newton formula and QuadTrees.\nCurrently " + strconv.Itoa(len(s.bodies)) + " bodies"
}
