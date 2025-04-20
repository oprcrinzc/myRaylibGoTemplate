package gs

func Call(n string) {
	if n == "inside store" {
		insideStore()
	} else if n == "outside store" {
		outsideStore()
	}
}
