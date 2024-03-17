package app

import (
	"github.com/RugiSerl/physics/app/camera"
	simulation "github.com/RugiSerl/physics/app/projects/bodySimulation"
	"github.com/RugiSerl/physics/app/values"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	sim      *simulation.Simulation
	myCamera *camera.Camera2D
)

func Run() {
	create()
	for !rl.WindowShouldClose() {
		update()

	}
	quit()
}

func create() {

	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(int32(rl.GetMonitorWidth(0)), int32(rl.GetMonitorHeight(0)), "raylib [core] example - basic window")

	sim = simulation.Create()
	rl.ToggleFullscreen()
	myCamera = camera.NewCamera()

	rl.SetTargetFPS(-1)
}

func update() {
	values.UpdateValues()
	rl.BeginDrawing()

	myCamera.UpdateCamera()
	myCamera.Begin()

	rl.DrawText(sim.ProvideDescription(), 190, 200, 20, rl.LightGray)
	rl.ClearBackground(rl.RayWhite)

	sim.Update(myCamera)
	myCamera.End()

	rl.EndDrawing()
}

func quit() {
	rl.CloseWindow()

}

func updateInput() {
	if rl.IsKeyPressed(rl.KeyF11) {
		rl.ToggleFullscreen()
	}
}
