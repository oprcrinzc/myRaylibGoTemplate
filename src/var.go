package src

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type (
	WindowConfig struct {
		Flag   uint32 `toml:"flag"`
		Width  int32  `toml:"width"`
		Height int32  `toml:"height"`
		Title  string `toml:"title"`
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
		Debug      bool         `toml:"debug"`
		GameCursor bool         `toml:"gameCursor"`
		Fullscreen bool         `toml:"fullscreen"`
		Sound      SoundConfig  `toml:"sound"`
		Fps        int32        `toml:"fps"`
		Window     WindowConfig `toml:"window"`
		KeyMap     KeyMapConfig `toml:"keymap"`
		Lang       string       `toml:"lang"`
	}
)

type SequenceFunc []func()

func (s *SequenceFunc) Add(f func()) *SequenceFunc {
	*s = append(*s, f)
	return s
}
func (s *SequenceFunc) Run() *SequenceFunc {
	for _, f := range *s {
		f()
	}
	return s
}
func (s *SequenceFunc) Size() (i int) {
	for range *s {
		i++
	}
	return
}
func (s *SequenceFunc) String() (t string) {
	t += fmt.Sprintf("+ Sequence --- %d\n", s.Size())
	return
}

func CreateWindow(width, height, fps int32, title string, configFlag uint32) {
	rl.SetConfigFlags(configFlag)
	rl.InitWindow(width, height, title)
	rl.SetTargetFPS(fps)
}
func (c *Config) LoadConfig() {
	var config Config
	d, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// d = d[:len(d)-3]
	_, err = toml.DecodeFile(d+"/config.toml", &config)
	if err != nil {
		panic(err)
	}
	*c = config
	fmt.Println("Loaded: Config")
	if c.Debug {
		fmt.Println(config)
	}
}

var rt2Post rl.RenderTexture2D
var rt2CurrentScene rl.RenderTexture2D

var cfg *Config = new(Config)
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

var Seq *SequenceFunc = new(SequenceFunc)
