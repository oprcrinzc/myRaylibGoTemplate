package gs

import (
	"oprc_core/src/global"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func insideStore() {
	rl.BeginTextureMode(global.Rt2CurrentScene)
	rl.ClearScreenBuffers()
	rl.ClearBackground(rl.SkyBlue)
	rl.DrawTextEx(global.FontMed32, "inside", rl.NewVector2(60, 60), 32, 4, rl.RayWhite)
	rl.EndTextureMode()
}
