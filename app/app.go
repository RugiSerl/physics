package app

import (
	"github.com/RugiSerl/physics/app/camera"
	"github.com/RugiSerl/physics/app/projects"
	"github.com/RugiSerl/physics/app/values"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	myProject projects.Project
	myCamera  *camera.Camera2D
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

	rl.InitWindow(int32(rl.GetMonitorWidth(0)), int32(rl.GetMonitorHeight(0)), "Physics")
	myProject = projects.NewProject(projects.PROJECT_SPRING)
	myCamera = camera.NewCamera()

	rl.SetTargetFPS(-1)
}

func update() {
	values.UpdateValues()
	rl.BeginDrawing()

	myCamera.UpdateCamera()
	myCamera.Begin()

	rl.DrawText(myProject.ProvideDescription(), 0, 0, 20, rl.LightGray)
	rl.ClearBackground(rl.RayWhite)

	if rl.IsKeyPressed(rl.KeyKpAdd) {
		values.DtFactor *= 2
	}
	if rl.IsKeyPressed(rl.KeyKpSubtract) {
		values.DtFactor /= 2
	}
	myProject.Update(myCamera)
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
