package components

import rl "github.com/gen2brain/raylib-go/raylib"

var W int32
var H int32
var MOUSEPOS rl.Vector2

func UpdateComponents(w, h int32, mp rl.Vector2) {
	W = w
	H = h
	MOUSEPOS = mp
}
