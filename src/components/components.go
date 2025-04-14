package components

import rl "github.com/gen2brain/raylib-go/raylib"

var W int32
var H int32
var mousePos rl.Vector2

var FontMed128 rl.Font
var FontMed64 rl.Font

func Load() {
	FontMed128 = rl.LoadFontEx("./assets/fonts/Prompt/Prompt-Medium.ttf", 128, nil, 0)
	FontMed64 = rl.LoadFontEx("./assets/fonts/Prompt/Prompt-Medium.ttf", 64, nil, 0)
	// FontMed50 = rl.LoadFont("./assets/fonts/Prompt/Prompt-Medium.ttf")
}
func UnLoad() {
	rl.UnloadFont(FontMed128)
	rl.UnloadFont(FontMed64)
}
