package scene

import (
	"oprc_core/src/global"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Intro() {
	rl.BeginTextureMode(global.Rt2CurrentScene)
	rl.ClearBackground(rl.Black)
	// rl.OpenURL("https://github.com/gen2brain/raylib-go")
	rl.EndTextureMode()
}
