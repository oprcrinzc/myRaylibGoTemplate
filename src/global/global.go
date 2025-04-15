package global

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var WantExit bool = false

func Delta(a, b float64) float64 {
	if a > b {
		return a - b
	}
	return b - a
}

func Distance(p1, p2 rl.Vector2) (dd float64) {
	dx := Delta(float64(p1.X), float64(p2.X))
	dy := Delta(float64(p1.Y), float64(p2.Y))
	dd = math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
	return
}
