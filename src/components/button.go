package components

import (
	"oprc_core/src/global"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type (
	Button struct {
		Text     string
		Pos      rl.Vector2
		IsHover  bool
		Func     func(self *Button)
		Width    int32
		Height   int32
		FontPack global.FontPack
	}
	Drawable interface {
		Draw()
	}
)

func (b *Button) Make(t string, w, h int32, pos rl.Vector2) *Button {
	b.Pos = pos
	b.Text = t
	b.IsHover = false
	b.Width = w
	b.Height = h
	return b
}
func (b *Button) SetFunc(f func(self *Button)) *Button {
	b.Func = f
	return b
}
func (b *Button) Do() *Button {
	if b.Func != nil {
		b.Func(b)
	}
	return b
}
func (b *Button) Draw() {
	MOUSEPOS = rl.GetMousePosition()
	// check hover

	if MOUSEPOS.X >= b.Pos.X && MOUSEPOS.X <= b.Pos.X+float32(b.Width) &&
		MOUSEPOS.Y >= b.Pos.Y && MOUSEPOS.Y <= b.Pos.Y+float32(b.Height) {
		b.IsHover = true
	} else {
		b.IsHover = false
	}

	col := rl.Pink
	if b.IsHover {
		col = rl.Beige
		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			b.Do()
		}
	}
	rl.DrawRectangle(int32(b.Pos.X), int32(b.Pos.Y), b.Width, b.Height, col)

	ts := rl.MeasureTextEx(b.FontPack.Font, b.Text, float32(b.FontPack.Size), float32(b.FontPack.Spacing))
	tp := rl.NewVector2(
		b.Pos.X-ts.X/2+float32(b.Width)/2,
		b.Pos.Y-ts.Y/2+float32(b.Height)/2)
	// rl.DrawCircle(int32(tp.X), int32(tp.Y), 5, rl.Red)
	rl.DrawTextEx(b.FontPack.Font, b.Text, tp, float32(b.FontPack.Size), float32(b.FontPack.Spacing), rl.White)
}
