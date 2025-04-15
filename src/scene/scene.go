package scene

import t "oprc_core/src/ot"

var W int32
var H int32
var STATE *int = new(int)
var FONTPACKS map[string]t.FontPack = make(map[string]t.FontPack)

func FromState(state *int, w, h int32) {
	STATE = state
	W = w
	H = h
	switch *state {
	case 0:
		Intro()
	case 1:
		MainMenu()

	}
}
