package global

import (
	"fmt"
	"io/fs"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pelletier/go-toml/v2"
)

type (
	InputSystem struct {
	}
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
		Close   int `toml:"close"`
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
	readf, err := fs.ReadFile(os.DirFS(d), "config.toml")
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal(readf, &config)
	if err != nil {
		panic(err)
	}
	*c = config
	if c.Debug {
		fmt.Println("Loaded: Config")
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

const (
	Fuyu = iota
	Haru
	Natsu
	Aki
)
const (
	Sunny = iota
	Cloudy
	Rain
	LightRain
	HeavyRain
	Humid
	Snow
	Thunder
	Fog
	Wind
	Hot
	Cool
	Cold
)
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
const (
	January = iota
	Februay
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

var DayString map[int]string = map[int]string{
	Sunday:    "Sunday",
	Monday:    "Monday",
	Tuesday:   "Tuesday",
	Wednesday: "Wednesday",
	Thursday:  "Thursday",
	Friday:    "Friday",
	Saturday:  "Saturday",
}

var MonthString map[int]string = map[int]string{
	January:   "January",
	Februay:   "Februay",
	March:     "March",
	April:     "April",
	May:       "May",
	June:      "June",
	July:      "July",
	August:    "August",
	October:   "October",
	November:  "November",
	December:  "December",
	September: "September",
}

type (
	Date struct {
		Day          int
		Month        int
		DayFromStart int
		Year         int
	}
	Weather struct {
		Type        int
		Temperature float32
		Humidity    float32
	}
	World struct {
		Season  int
		Weather Weather
		Date    Date
	}
)

type (
	Player struct {
		Name       string
		MouseLayer int32
		hovers     []bool
		CanClick   bool
	}
)

func (p *Player) Born(n string) *Player {
	p.Name = n
	p.MouseLayer = 0
	p.hovers = []bool{}
	p.CanClick = false
	return p
}

func (p *Player) AddHover(b bool) *Player {
	p.hovers = append(p.hovers, b)
	return p
}

func (p *Player) Update() {
	hoverTrueCount := 0
	for _, t := range p.hovers {
		if t {
			hoverTrueCount++
		}
	}
	if hoverTrueCount > 0 && !rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		p.CanClick = true
	} else if hoverTrueCount <= 0 && rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		p.CanClick = false
	}
	if hoverTrueCount > 0 && p.CanClick {
		rl.SetMouseCursor(rl.MouseCursorPointingHand)
	} else {
		rl.SetMouseCursor(rl.MouseCursorDefault)
	}
	p.hovers = []bool{}
}

type (
	Updatable interface{ Update() }
)

type (
	SoundSys struct {
		Sounds map[string]map[string]rl.Sound
		Music  map[string]rl.Music
		sel    string
		played bool
		paused bool
	}
)

func (s *SoundSys) SetVolume() {
	main := float32(Cfg.Sound.Main) / 100.0
	for _, i := range s.Sounds["Action"] {
		rl.SetSoundVolume(i, main*float32(Cfg.Sound.Action)/100.0)
	}
	for _, i := range s.Sounds["Animal"] {
		rl.SetSoundVolume(i, main*float32(Cfg.Sound.Animal)/100.0)
	}
	for _, i := range s.Sounds["Environment"] {
		rl.SetSoundVolume(i, main*float32(Cfg.Sound.Environment)/100.0)
	}
	for _, i := range s.Sounds["People"] {
		rl.SetSoundVolume(i, main*float32(Cfg.Sound.People)/100.0)
	}
	for _, i := range s.Music {
		rl.SetMusicVolume(i, main*float32(Cfg.Sound.Music)/100.0)
	}
}

func (s *SoundSys) Born() *SoundSys {
	s.Sounds = make(map[string]map[string]rl.Sound)
	s.Music = make(map[string]rl.Music)
	s.played = false
	return s
}

func (s *SoundSys) AddSound(g, n, f string) *SoundSys {
	s.Sounds[g][n] = rl.LoadSound(f)
	return s
}

func (s *SoundSys) AddMusic(n, f string) *SoundSys {
	s.Music[n] = rl.LoadMusicStream(f)
	return s
}

func (s *SoundSys) SelectMusic(f string) {
	s.sel = f
	s.played = false
}

func (s *SoundSys) Play(g, n string) {
	rl.PlaySound(s.Sounds[g][n])
}

func (s *SoundSys) Update() {
	if s.Music != nil {
		m := s.Music[s.sel]
		if !s.played {
			rl.PlayMusicStream(m)
			s.played = true
		}
		rl.UpdateMusicStream(m)
	}
}

func (s *SoundSys) ToggleMusic() { // pause play music
	if !s.paused {
		rl.PauseMusicStream(s.Music[s.sel])
	} else {
		rl.ResumeMusicStream(s.Music[s.sel])
	}
}
