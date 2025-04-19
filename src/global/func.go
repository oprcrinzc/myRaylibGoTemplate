package global

import (
	"math"
	"strconv"

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

func LoadPlayer() {
	PlayerA.Born("Name")
}

func LoadSoundSys() {
	rl.InitAudioDevice()
	SoundA.Born()
}

func DeltaD(a, b float32) float32 {
	if a > b {
		return a / b
	}
	return b / a
}

func RSO(w, h, tw, th int32) (s float32, o rl.Vector2) {
	var x1, x2 float32
	x1 = float32(h) / float32(th)
	x2 = float32(w) / float32(tw)
	if float32(tw)*x1 > float32(w) || float32(th)*x1 > float32(h) {
		s = x2
	}
	if float32(tw)*x2 > float32(w) || float32(th)*x2 > float32(h) {
		s = x1
	} else {
		s = (x1 + x2) / 2
	}

	o = rl.NewVector2(
		(float32((float64(w) - float64(tw)*float64(s))))/2,
		(float32((float64(h) - float64(th)*float64(s))))/2)
	// fmt.Println(" x1:", x1, " x2:", x2, " s:", s, " o:", o)
	return
}

func CompareSliceByte(a, b []byte) bool {
	s1 := ""
	s2 := ""
	for _, i := range a {
		s1 += strconv.Itoa(int(i))
	}
	for _, i := range b {
		s2 += strconv.Itoa(int(i))
	}
	return s1 == s2
}
