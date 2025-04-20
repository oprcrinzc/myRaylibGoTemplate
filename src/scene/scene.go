package scene

import (
	"oprc_core/src/global"
	gs "oprc_core/src/scene/game"
)

var W int32
var H int32
var STATE *int = new(int)
var FONTPACKS map[string]global.FontPack = make(map[string]global.FontPack)

func FromState(state *int, w, h int32) {
	STATE = state
	W = w
	H = h
	switch *state {
	case 0:
		Intro()
	case 1:
		MainMenu()
	case 2:
		gs.Call(global.PlayerA.GameScene)

	}
}
