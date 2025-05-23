package scene

import (
	"oprc_core/src/components"
	"oprc_core/src/global"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var bgMainMenu *rl.Texture2D = new(rl.Texture2D)

func mainMenuLoad() {
	*bgMainMenu = rl.LoadTexture("./assets/images/MainMenuBG.JPG")
	// *bgMainMenu = rl.LoadTexture("./assets/images/test.jpg")
	// *bgMainMenu = rl.LoadTexture("./assets/images/cover2.jpg")
	// *bgMainMenu = rl.LoadTexture("./assets/images/Cover-2.jpg")
	// Load()
}
func mainMunuUnload() {
	rl.UnloadTexture(*bgMainMenu)
}

func MainMenu() {
	if bgMainMenu.Width <= 0 {
		mainMenuLoad()
	}
	rl.BeginTextureMode(global.Rt2CurrentScene)
	rl.ClearScreenBuffers()

	// ss := (float32(W)/float32(bgMainMenu.Width) + float32(H)/float32(bgMainMenu.Height)) / 2
	// origin := rl.NewVector2((float32(W)-(float32(bgMainMenu.Width)*ss))/2, 0)

	ss, origin := global.RSO(W, H, bgMainMenu.Width, bgMainMenu.Height)

	rl.DrawTextureEx(*bgMainMenu, origin, 0, ss, rl.White)

	btn1 := components.Button{FontPack: FONTPACKS["med64"]}
	btn1.Make("Enter", 200, 100, rl.Vector2Add(rl.NewVector2(50, float32(H-150)-130), origin)).SetFunc(
		func(self *components.Button) {
			global.State = 2
			if global.PlayerA.GameScene == "" {
				global.PlayerA.GameScene = "inside store"
			}
			mainMunuUnload()
		}).Draw()

	btn := components.Button{FontPack: FONTPACKS["med64"]}
	btn.Make("Exit", 200, 100, rl.Vector2Add(rl.NewVector2(50, float32(H-150)), origin)).SetFunc(
		func(self *components.Button) {
			global.WantExit = true
		}).Draw()

	rl.EndTextureMode()
}
