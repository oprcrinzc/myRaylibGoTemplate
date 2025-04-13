package src

func Call(score int) {
	if score == 0 {
		cfg := new(Config).LoadConfig()
		CreateWindow(cfg.Window.Width, cfg.Window.Height, cfg.Fps, cfg.Window.Title, cfg.Window.Flag)
	}
}
