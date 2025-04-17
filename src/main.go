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
	global.CreateWindow(global.Cfg.Window.Width, global.Cfg.Window.Height,
		global.Cfg.Fps, global.Cfg.Window.Title,
		(global.Cfg.Window.Flag&^rl.FlagFullscreenMode)|rl.FlagMsaa4xHint)
	rl.SetWindowMinSize(int(global.Cfg.Window.Width), int(global.Cfg.Window.Height))
	rl.SetExitKey(rl.KeyNull)

	rl.InitAudioDevice()

	global.WIDTH = int32(rl.GetScreenWidth())
	global.HEIGHT = int32(rl.GetScreenHeight())

	global.LoadFont()
	global.LoadRT2d()
	global.LoadPlayer()

	scene.FONTPACKS = global.FontPacks

	s1 := rl.LoadMusicStream("./assets/audio/bgm0.wav")
	rl.PlayMusicStream(s1)
	global.Seq.
		Add(global.ToggleFullscreenLogic).
		Add(global.ReLoadRT2dLogic).
		Add(global.PlayerA.Update).
		Add(func() { rl.UpdateMusicStream(s1) })

	// rl.LoadAudioStream(441000, 24, 0)
}

func gameLoop() {
	defer rl.CloseWindow()
	defer global.UnLoadRT2d()

	for !rl.WindowShouldClose() && !global.WantExit {
		global.Seq.Run()

		scene.FromState(&global.State, global.WIDTH, global.HEIGHT)

		rl.BeginDrawing()
		rl.ClearScreenBuffers()
		rl.DrawTexturePro(global.Rt2CurrentScene.Texture, rl.NewRectangle(0, 0, float32(global.WIDTH), -float32(global.HEIGHT)),
			rl.NewRectangle(0, 0, float32(global.WIDTH), float32(global.HEIGHT)), rl.NewVector2(float32(global.WIDTH)-float32(global.Rt2CurrentScene.Texture.Width), 0), 0, rl.White)
		rl.EndDrawing()
	}
}
