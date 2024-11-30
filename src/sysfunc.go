package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func SystemPanic(error string) {
	state.error = error
	state.state = 3
}

func SystemPanicInternal() {
	rl.ClearBackground(rl.Blue)

	ZiDrawText(":(", 30, 10, 300, rl.White)

	ZiDrawText("Zivic has ran into an issue and was forced to shutdown the system", 300, 300, 50, rl.White)
	string := fmt.Sprintf("Error Code: %s", state.error)
	ZiDrawText(string, 300, 360, 50, rl.White)
}
