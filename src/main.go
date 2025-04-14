package src

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Call() {
	loadCfg()
	setUp()
	gameLoop()
}

func loadCfg() {
	cfg.LoadConfig()
}

func setUp() {
	if !cfg.Debug {
		rl.SetTraceLogLevel(rl.LogFatal)
	} else {
		fmt.Print("+ --- Raylib --- +\n")
	}
	CreateWindow(cfg.Window.Width, cfg.Window.Height, cfg.Fps, cfg.Window.Title, cfg.Window.Flag|rl.FlagWindowResizable)
	rl.SetExitKey(rl.KeyNull)
	LoadRT2d()

	Seq.Add(func() {
		if rl.IsKeyPressed(rl.KeyF11) {
			rl.ToggleFullscreen()
		}
	})
}

func gameLoop() {
	defer rl.CloseWindow()
	defer UnLoadRT2d()
	for !rl.WindowShouldClose() {
		Seq.Run()

		cfg.Window.Height = int32(rl.GetScreenHeight())
		cfg.Window.Width = int32(rl.GetScreenWidth())
		ReLoadRT2d()

		rl.BeginTextureMode(rt2CurrentScene)
		rl.ClearBackground(rl.SkyBlue)
		rl.DrawText("Hello", 50, 50, 50, rl.Blue)
		rl.EndTextureMode()

		rl.BeginDrawing()
		rl.ClearScreenBuffers()
		rl.DrawTexturePro(rt2CurrentScene.Texture, rl.NewRectangle(0, 0, float32(cfg.Window.Width), -float32(cfg.Window.Height)),
			rl.NewRectangle(0, 0, float32(cfg.Window.Width), float32(cfg.Window.Height)), rl.NewVector2(0, 0), 0, rl.White)
		rl.EndDrawing()
	}
}
