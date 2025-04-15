package src

import (
	"fmt"
	"oprc_core/src/global"
	"oprc_core/src/scene"

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
	// lww := 0
	if !cfg.Debug {
		rl.SetTraceLogLevel(rl.LogFatal)
	} else {
		fmt.Print("+ --- Raylib --- +\n")
	}
	// |rl.FlagWindowResizable
	CreateWindow(cfg.Window.Width, cfg.Window.Height, cfg.Fps, cfg.Window.Title,
		(cfg.Window.Flag&^rl.FlagFullscreenMode)|rl.FlagMsaa4xHint)
	rl.SetWindowMinSize(int(cfg.Window.Width), int(cfg.Window.Height))
	rl.SetExitKey(rl.KeyNull)

	WIDTH = int32(rl.GetScreenWidth())
	HEIGHT = int32(rl.GetScreenHeight())

	LoadFont()
	LoadRT2d()

	scene.FONTPACKS = FontPacks

	Seq.Add(func() {
		if rl.IsKeyPressed(rl.KeyF11) {
			WIDTH = int32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))
			HEIGHT = int32(rl.GetMonitorHeight(rl.GetCurrentMonitor()))
			rl.ToggleFullscreen()
			ISFULLSCREEN = !ISFULLSCREEN
		}
	}).Add(func() {
		if ISFULLSCREEN {
			WIDTH = int32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))
			HEIGHT = int32(rl.GetMonitorHeight(rl.GetCurrentMonitor()))
		} else {
			WIDTH = int32(rl.GetScreenWidth())
			HEIGHT = int32(rl.GetScreenHeight())
		}
		ReLoadRT2d()
		// fmt.Printf("w:%v h:%v w/h:%v 16/9:%v e:%t\n", WIDTH, HEIGHT, float32(WIDTH)/float32(HEIGHT), 16.0/9.0, float32(WIDTH)/float32(HEIGHT) == 16.0/9.0)
		// fmt.Printf("%v %v | %v | %v %v \n", int(float32(WIDTH)*lmmm), int(float32(HEIGHT)*lmmm), lmmm, lm, lmm)
	})
}

func gameLoop() {
	defer rl.CloseWindow()
	defer UnLoadRT2d()
	for !rl.WindowShouldClose() && !global.WantExit {
		Seq.Run()

		rl.BeginTextureMode(rt2CurrentScene)
		scene.FromState(&state, WIDTH, HEIGHT)
		rl.EndTextureMode()

		rl.BeginDrawing()
		rl.ClearScreenBuffers()
		rl.DrawTexturePro(rt2CurrentScene.Texture, rl.NewRectangle(0, 0, float32(WIDTH), -float32(HEIGHT)),
			rl.NewRectangle(0, 0, float32(WIDTH), float32(HEIGHT)), rl.NewVector2(float32(WIDTH)-float32(rt2CurrentScene.Texture.Width), 0), 0, rl.White)
		rl.EndDrawing()
	}
}
