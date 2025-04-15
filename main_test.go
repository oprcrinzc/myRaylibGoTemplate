package main

import (
	"oprc_core/src/global"
	"testing"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func TestLoadConfig(t *testing.T) {
	cf := global.Config{}
	cf.LoadConfig()
	if cf.Fps <= 0 {
		t.Fail()
	}
}

func TestDistance1(t *testing.T) {
	v1 := rl.NewVector2(0, 0)
	v2 := rl.NewVector2(0, 10)
	d := global.Distance(v1, v2)
	de := 10.0
	if d != de {
		t.Errorf("[%v] distance (%v, %v) and (%v, %v) = %v", de, v1.X, v1.Y, v2.X, v2.Y, d)
	}
}

func TestDistance2(t *testing.T) {
	v1 := rl.NewVector2(0, 10)
	v2 := rl.NewVector2(0, 0)
	d := global.Distance(v1, v2)
	de := 10.0
	if d != de {
		t.Errorf("[%v] distance (%v, %v) and (%v, %v) = %v", de, v1.X, v1.Y, v2.X, v2.Y, d)
	}
}

func TestDistance3(t *testing.T) {
	v1 := rl.NewVector2(0, 0)
	v2 := rl.NewVector2(10, 0)
	d := global.Distance(v1, v2)
	de := 10.0
	if d != de {
		t.Errorf("[%v] distance (%v, %v) and (%v, %v) = %v", de, v1.X, v1.Y, v2.X, v2.Y, d)
	}
}

func TestDistance4(t *testing.T) {
	v1 := rl.NewVector2(10, 0)
	v2 := rl.NewVector2(0, 0)
	d := global.Distance(v1, v2)
	de := 10.0
	if d != de {
		t.Errorf("[%v] distance (%v, %v) and (%v, %v) = %v", de, v1.X, v1.Y, v2.X, v2.Y, d)
	}
}

func TestDistance5(t *testing.T) {
	v1 := rl.NewVector2(0, 0)
	v2 := rl.NewVector2(3, 4)
	d := global.Distance(v1, v2)
	de := 5.0
	if d != de {
		t.Errorf("[%v] distance (%v, %v) and (%v, %v) = %v", de, v1.X, v1.Y, v2.X, v2.Y, d)
	}
}
