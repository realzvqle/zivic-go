package main

//import rl "github.com/gen2brain/raylib-go/raylib"

func SetState() {
	switch state.state {
	case 1:
		Desktop()
	case 2:
		// task manager
	case 3:
		SystemPanicInternal()
	default:
		SystemPanic("INVALID STATE")
	}
}
