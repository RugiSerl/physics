package values

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	Dt       float64
	DtFactor float64 = 1
)

func UpdateValues() {
	Dt = float64(rl.GetFrameTime()) * DtFactor
}
