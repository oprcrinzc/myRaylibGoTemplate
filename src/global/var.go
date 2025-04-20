package global

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var WantExit bool = false

var FontMed128, FontMed64, FontMed32, FontMed16 rl.Font
var FontPacks map[string]FontPack = make(map[string]FontPack)

var Rt2Post, Rt2CurrentScene rl.RenderTexture2D

var Cfg *Config = new(Config)
var WIDTH int32
var HEIGHT int32
var ISFULLSCREEN bool = false

var State int = 0

var Seq *SequenceFunc = new(SequenceFunc)

var PlayerA Player
var SoundA SoundSys
