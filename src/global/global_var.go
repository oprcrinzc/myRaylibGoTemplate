package global

import (
	"oprc_core/src/ot"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var WantExit bool = false

var FontMed128, FontMed64, FontMed32, FontMed16 rl.Font

var FontPacks map[string]ot.FontPack = make(map[string]ot.FontPack)

var Rt2Post rl.RenderTexture2D
var Rt2CurrentScene rl.RenderTexture2D

var Cfg *ot.Config = new(ot.Config)
var WIDTH int32
var HEIGHT int32
var ISFULLSCREEN bool = false

var State int = 1

var Seq *ot.SequenceFunc = new(ot.SequenceFunc)
