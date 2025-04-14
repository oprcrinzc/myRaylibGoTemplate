package scene

var W int32
var H int32
var STATE *int = new(int)

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
