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
	global.Cfg.LoadConfig()
}

func setUp() {
	// lww := 0
	if !global.Cfg.Debug {
		rl.SetTraceLogLevel(rl.LogFatal)
	} else {
		fmt.Print("+ --- Raylib --- +\n")
	}
	// |rl.FlagWindowResizable
	global.CreateWindow(global.Cfg.Window.Width, global.Cfg.Window.Height, global.Cfg.Fps, global.Cfg.Window.Title,
		(global.Cfg.Window.Flag&^rl.FlagFullscreenMode)|rl.FlagMsaa4xHint)
	rl.SetWindowMinSize(int(global.Cfg.Window.Width), int(global.Cfg.Window.Height))
	rl.SetExitKey(rl.KeyNull)

	global.WIDTH = int32(rl.GetScreenWidth())
	global.HEIGHT = int32(rl.GetScreenHeight())

	global.LoadFont()
	global.LoadRT2d()

	scene.FONTPACKS = global.FontPacks

	global.Seq.Add(func() {
		if rl.IsKeyPressed(rl.KeyF11) {
			global.WIDTH = int32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))
			global.HEIGHT = int32(rl.GetMonitorHeight(rl.GetCurrentMonitor()))
			rl.ToggleFullscreen()
			global.ISFULLSCREEN = !global.ISFULLSCREEN
		}
	}).Add(func() {
		if global.ISFULLSCREEN {
			global.WIDTH = int32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))
			global.HEIGHT = int32(rl.GetMonitorHeight(rl.GetCurrentMonitor()))
		} else {
			global.WIDTH = int32(rl.GetScreenWidth())
			global.HEIGHT = int32(rl.GetScreenHeight())
		}
		global.ReLoadRT2d()
		// fmt.Printf("w:%v h:%v w/h:%v 16/9:%v e:%t\n", WIDTH, HEIGHT, float32(WIDTH)/float32(HEIGHT), 16.0/9.0, float32(WIDTH)/float32(HEIGHT) == 16.0/9.0)
		// fmt.Printf("%v %v | %v | %v %v \n", int(float32(WIDTH)*lmmm), int(float32(HEIGHT)*lmmm), lmmm, lm, lmm)
	})
}

func gameLoop() {
	defer rl.CloseWindow()
	defer global.UnLoadRT2d()
	for !rl.WindowShouldClose() && !global.WantExit {
		global.Seq.Run()

		rl.BeginTextureMode(global.Rt2CurrentScene)
		scene.FromState(&global.State, global.WIDTH, global.HEIGHT)
		rl.EndTextureMode()

		rl.BeginDrawing()
		rl.ClearScreenBuffers()
		rl.DrawTexturePro(global.Rt2CurrentScene.Texture, rl.NewRectangle(0, 0, float32(global.WIDTH), -float32(global.HEIGHT)),
			rl.NewRectangle(0, 0, float32(global.WIDTH), float32(global.HEIGHT)), rl.NewVector2(float32(global.WIDTH)-float32(global.Rt2CurrentScene.Texture.Width), 0), 0, rl.White)
		rl.EndDrawing()
	}
}
