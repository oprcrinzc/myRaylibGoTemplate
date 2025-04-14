package scene

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var bgMainMenu *rl.Texture2D = new(rl.Texture2D)

func mainMenuLoad() {
	*bgMainMenu = rl.LoadTexture("./assets/images/cover2.jpg")
}

func MainMenu() {
	if bgMainMenu.Width == 0 {
		mainMenuLoad()
	}
	// rl.ClearBackground(rl.Black)
	rl.ClearScreenBuffers()
	// rl.DrawTexture(img, 0, 0, rl.White)
	ss := (float32(W)/float32(bgMainMenu.Width) + float32(H)/float32(bgMainMenu.Height)) / 2

	// fmt.Printf("%v %v %v %v", bgMainMenu.Height, H, W, ss)

	rl.DrawTextureEx(*bgMainMenu, rl.NewVector2((float32(W)-(float32(bgMainMenu.Width)*ss))/2, 0), 0, ss, rl.White)
}
