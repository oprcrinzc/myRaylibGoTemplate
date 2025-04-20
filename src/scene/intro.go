package scene

import (
	"oprc_core/src/global"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var raylibgoLogo *rl.Texture2D = new(rl.Texture2D)

var introTime float32 = 7
var introTimeCount float32
var logoAlpha float32

func introLoad() {
	*raylibgoLogo = rl.LoadTexture("./assets/images/raylibgo.png")
}

func Intro() {
	if raylibgoLogo.Width <= 0 {
		introLoad()
	}
	s := float32(1.0)
	rl.BeginTextureMode(global.Rt2CurrentScene)
	rl.ClearScreenBuffers()
	rl.DrawTextureEx(*raylibgoLogo,
		rl.NewVector2(
			float32(global.WIDTH-raylibgoLogo.Width)/2,
			float32(global.HEIGHT-raylibgoLogo.Height)/2),
		0, s, rl.NewColor(255, 255, 255, uint8(logoAlpha)))
	mt := rl.MeasureTextEx(global.FontMed32, "Made with RaylibGo", 32, 4)
	rl.DrawTextEx(global.FontMed32, "Made with RaylibGo",
		rl.NewVector2(
			(float32(global.WIDTH)-mt.X)/2,
			(float32(global.HEIGHT)-mt.Y)/2+float32(raylibgoLogo.Height)/1.7),
		32, 4, rl.NewColor(245, 245, 245, uint8(logoAlpha)))

	rl.EndTextureMode()

	introTimeCount += 1 * rl.GetFrameTime()
	if logoAlpha > 255 {
		logoAlpha = 255
	} else {
		logoAlpha += 40 * rl.GetFrameTime()
	}
	if logoAlpha > 255 {
		logoAlpha = 255
	}

	if introTimeCount >= introTime {
		global.State = 1
	}
}
