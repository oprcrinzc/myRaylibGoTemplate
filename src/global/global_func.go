package global

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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

func LoadFont() {
	FontMed128 = rl.LoadFontEx("./assets/fonts/Prompt/Prompt-Medium.ttf", 128, nil, 0)
	FontMed64 = rl.LoadFontEx("./assets/fonts/Prompt/Prompt-Medium.ttf", 64, nil, 0)
	FontMed32 = rl.LoadFontEx("./assets/fonts/Prompt/Prompt-Medium.ttf", 32, nil, 0)
	FontMed16 = rl.LoadFontEx("./assets/fonts/Prompt/Prompt-Medium.ttf", 16, nil, 0)

	FontPacks["med128"] = FontPack{Font: FontMed128, Size: 128, Spacing: 2}
	FontPacks["med64"] = FontPack{Font: FontMed64, Size: 64, Spacing: 2}
	FontPacks["med32"] = FontPack{Font: FontMed32, Size: 32, Spacing: 2}
	FontPacks["med16"] = FontPack{Font: FontMed16, Size: 16, Spacing: 2}

	for _, f := range FontPacks {
		rl.GenTextureMipmaps(&f.Font.Texture)
		rl.SetTextureFilter(f.Font.Texture, rl.FilterBilinear)
	}
}
func UnLoadFont() {
	rl.UnloadFont(FontMed128)
	rl.UnloadFont(FontMed64)
	rl.UnloadFont(FontMed32)
	rl.UnloadFont(FontMed16)
}

func CreateWindow(width, height, fps int32, title string, configFlag uint32) {
	rl.SetConfigFlags(configFlag)
	rl.InitWindow(width, height, title)
	rl.SetTargetFPS(fps)
}

func LoadRT2d() {
	Rt2Post = rl.LoadRenderTexture(WIDTH, HEIGHT)
	Rt2CurrentScene = rl.LoadRenderTexture(WIDTH, HEIGHT)
	rl.SetTextureFilter(Rt2Post.Texture, rl.FilterAnisotropic16x)
	rl.SetTextureFilter(Rt2CurrentScene.Texture, rl.FilterAnisotropic16x)
}
func UnLoadRT2d() {
	rl.UnloadRenderTexture(Rt2Post)
	rl.UnloadRenderTexture(Rt2CurrentScene)
}
func ReLoadRT2d() {
	UnLoadRT2d()
	LoadRT2d()
}

func ToggleFullscreenLogic() {
	if rl.IsKeyPressed(rl.KeyF11) {
		WIDTH = int32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))
		HEIGHT = int32(rl.GetMonitorHeight(rl.GetCurrentMonitor()))
		rl.ToggleFullscreen()
		ISFULLSCREEN = !ISFULLSCREEN
	}
}

func ReLoadRT2dLogic() {
	if ISFULLSCREEN {
		WIDTH = int32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))
		HEIGHT = int32(rl.GetMonitorHeight(rl.GetCurrentMonitor()))
	} else {
		WIDTH = int32(rl.GetScreenWidth())
		HEIGHT = int32(rl.GetScreenHeight())
	}
	ReLoadRT2d()
}
