package gs

import (
	"oprc_core/src/global"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func outsideStore() {
	rl.BeginTextureMode(global.Rt2CurrentScene)
	rl.ClearScreenBuffers()
	rl.DrawTextEx(global.FontMed32, "outside", rl.NewVector2(60, 60), 32, 4, rl.RayWhite)
	rl.EndTextureMode()
}
