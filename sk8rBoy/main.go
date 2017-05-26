package main

import "github.com/pedroallenrevez/sk8rBoy"

func main() {
	var impl = sk8rBoy.GameManager
	var gman sk8rBoy.GManInterface = impl

	// width height nskaters speed radius r1 r2
	impl.Init(10, 10, 5, 1, 1, 1, 0)

	NumTurns := 100
	for x := 0; x < NumTurns; x++ {
		impl.Update()
	}
}
