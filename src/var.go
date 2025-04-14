package src

import (
	t "oprc_core/src/types"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func CreateWindow(width, height, fps int32, title string, configFlag uint32) {
	rl.SetConfigFlags(configFlag)
	rl.InitWindow(width, height, title)
	rl.SetTargetFPS(fps)
}

var rt2Post rl.RenderTexture2D
var rt2CurrentScene rl.RenderTexture2D

var cfg *t.Config = new(t.Config)
var WIDTH int32
var HEIGHT int32
var ISFULLSCREEN bool = false

var state int = 1

func LoadRT2d() {
	rt2Post = rl.LoadRenderTexture(WIDTH, HEIGHT)
	rt2CurrentScene = rl.LoadRenderTexture(WIDTH, HEIGHT)
}
func UnLoadRT2d() {
	rl.UnloadRenderTexture(rt2Post)
	rl.UnloadRenderTexture(rt2CurrentScene)
}
func ReLoadRT2d() {
	UnLoadRT2d()
	LoadRT2d()
}

var FontMed128 rl.Font
var FontMed64 rl.Font
var FontMed32 rl.Font
var FontMed16 rl.Font

var FontPacks map[string]t.FontPack = make(map[string]t.FontPack)

func LoadFont() {
	FontMed128 = rl.LoadFontEx("./assets/fonts/Prompt/Prompt-Medium.ttf", 128, nil, 0)
	FontMed64 = rl.LoadFontEx("./assets/fonts/Prompt/Prompt-Medium.ttf", 64, nil, 0)
	FontMed32 = rl.LoadFontEx("./assets/fonts/Prompt/Prompt-Medium.ttf", 32, nil, 0)
	FontMed16 = rl.LoadFontEx("./assets/fonts/Prompt/Prompt-Medium.ttf", 16, nil, 0)

	FontPacks["med128"] = t.FontPack{Font: FontMed128, Size: 128, Spacing: 4}
	FontPacks["med64"] = t.FontPack{Font: FontMed64, Size: 64, Spacing: 4}
	FontPacks["med32"] = t.FontPack{Font: FontMed32, Size: 32, Spacing: 4}
	FontPacks["med16"] = t.FontPack{Font: FontMed16, Size: 16, Spacing: 4}

	for _, f := range FontPacks {
		rl.SetTextureFilter(f.Font.Texture, rl.TextureFilterMode(2))
	}
}
func UnLoadFont() {
	rl.UnloadFont(FontMed128)
	rl.UnloadFont(FontMed64)
	rl.UnloadFont(FontMed32)
	rl.UnloadFont(FontMed16)
}

var Seq *t.SequenceFunc = new(t.SequenceFunc)
