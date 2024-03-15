package app

import (
	"fmt"

	"github.com/RugiSerl/physics/app/projects/simulation"
	"github.com/RugiSerl/physics/app/values"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	sim *simulation.Simulation
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

	rl.SetTargetFPS(-1)
}

func update() {
	values.UpdateValues()
	rl.BeginDrawing()

	fmt.Println(values.Dt)

	rl.ClearBackground(rl.RayWhite)
	rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
	sim.Update()
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
