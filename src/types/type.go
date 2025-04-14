package types

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

type (
	FontPack struct {
		Font    rl.Font
		Size    int32
		Spacing int32
	}
)
