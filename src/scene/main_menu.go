package scene

import (
	"oprc_core/src/components"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var bgMainMenu *rl.Texture2D = new(rl.Texture2D)

func mainMenuLoad() {
	*bgMainMenu = rl.LoadTexture("./assets/images/cover2.jpg")
	// Load()
}

func MainMenu() {
	if bgMainMenu.Width == 0 {
		mainMenuLoad()
	}
	// rl.ClearBackground(rl.Black)
	rl.ClearScreenBuffers()
	// rl.DrawTexture(img, 0, 0, rl.White)
	ss := (float32(W)/float32(bgMainMenu.Width) + float32(H)/float32(bgMainMenu.Height)) / 2
	// ss := (float32(W) / float32(bgMainMenu.Width)) / 1

	// fmt.Printf("%v %v %v %v", bgMainMenu.Height, H, W, ss)
	origin := rl.NewVector2((float32(W)-(float32(bgMainMenu.Width)*ss))/2, 0)

	rl.DrawTextureEx(*bgMainMenu, origin, 0, ss, rl.White)

	btn := components.Button{FontPack: FONTPACKS["med64"]}
	btn.Make("Config", 200, 100, rl.Vector2Add(rl.NewVector2(50, 50), origin)).SetFunc(
		func(self *components.Button) {
			*STATE = 0
		}).Draw()
}
