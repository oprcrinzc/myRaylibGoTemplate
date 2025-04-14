package components

import rl "github.com/gen2brain/raylib-go/raylib"

type (
	Button struct {
		Text        string
		Pos         rl.Vector2
		IsHover     bool
		Func        func(self *Button)
		Width       int32
		Height      int32
		Font        rl.Font
		FontSize    int32
		FontSpacing int32
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

	rl.DrawRectangle(int32(b.Pos.X), int32(b.Pos.Y), b.Width, b.Height, rl.SkyBlue)
	ts := rl.MeasureTextEx(b.Font, b.Text, float32(b.FontSize), float32(b.FontSpacing))
	tp := rl.NewVector2(
		b.Pos.X-ts.X/2+float32(b.Width)/2,
		b.Pos.Y-ts.Y/2+float32(b.Height)/2)
	// rl.DrawCircle(int32(tp.X), int32(tp.Y), 5, rl.Red)
	rl.DrawTextEx(b.Font, b.Text, tp, float32(b.FontSize), float32(b.FontSpacing), rl.RayWhite)
}
