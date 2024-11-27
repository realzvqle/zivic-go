package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type pstate struct {
	font rl.Font
}

var state pstate

var win window = window{x: 30, y: 30, title: "hi", sizeX: 300, sizeY: 400}

func main() {

	rl.SetTraceLogLevel(rl.LogFatal)
	rl.InitWindow(1600, 900, "Zivic Portable Desktop")
	defer rl.CloseWindow()
	//rl.SetTargetFPS(60)
	state.font = rl.LoadFont("resources/fonts/Mukta-ExtraBold.ttf")
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		//button := ZiDrawButton("Hi", 10, 10, 40, 40, rl.Red, rl.DarkGreen, rl.Brown, -1)
		SummonWIndow(&win, rl.Beige, rl.White, rl.Gray, rl.Brown, rl.White, true)
		//fmt.Printf("X is %d, Y is %d", win.x, win.y)
		// if button == BUTTON_DOWN_LEFT {
		// 	rl.DrawText("h", 90, 90, 10, rl.Blue)
		// }

		rl.EndDrawing()
	}
}
