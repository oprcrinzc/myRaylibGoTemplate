package scene

var W int32
var H int32

func FromState(state int, w, h int32) {
	W = w
	H = h
	switch state {
	case 0:
		Intro()
	case 1:
		MainMenu()

	}
}
