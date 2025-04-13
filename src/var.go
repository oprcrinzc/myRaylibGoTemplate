package src

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type (
	WindowConfig struct {
		Title  string `toml:"title"`
		Width  int32  `toml:"width"`
		Height int32  `toml:"height"`
		Flag   uint32 `toml:"flag"`
	}
	KeyMapConfig struct {
		Up      int `toml:"up"`
		Down    int `toml:"down"`
		Left    int `toml:"left"`
		Right   int `toml:"right"`
		In      int `toml:"in"`
		Out     int `toml:"out"`
		Confirm int `toml:"confirm"`
		Cancel  int `toml:"cancel"`
		Back    int `toml:"back"`
		Forward int `toml:"forward"`
	}
	SoundConfig struct {
		Main        uint8 `toml:"main"`
		Music       uint8 `toml:"music"`
		Action      uint8 `toml:"action"`
		Environment uint8 `toml:"environment"`
		People      uint8 `toml:"people"`
		Animal      uint8 `toml:"animal"`
	}

	Config struct {
		// struct field tags
		Window     WindowConfig `toml:"window"`
		KeyMap     KeyMapConfig `toml:"keymap"`
		Sound      SoundConfig  `toml:"sound"`
		Fps        int32        `toml:"fps"`
		Lang       string       `toml:"lang"`
		GameCursor bool         `toml:"gameCursor"`
		Fullscreen bool         `toml:"fullscreen"`
	}
)

func CreateWindow(width, height, fps int32, title string, configFlag uint32) {
	rl.SetConfigFlags(configFlag)
	rl.InitWindow(width, height, title)
	rl.SetTargetFPS(fps)
}
func (c *Config) LoadConfig() (config Config) {
	d, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// d = d[:len(d)-3]
	_, err = toml.DecodeFile(d+"/config.toml", &config)
	if err != nil {
		panic(err)
	}
	fmt.Println("Loaded: Config")
	fmt.Println(config)
	return
}
