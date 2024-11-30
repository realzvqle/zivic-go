package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PSTATE struct {
	font  rl.Font
	state uint8
	error string
}

var state PSTATE = PSTATE{state: 1}

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetTraceLogLevel(rl.LogFatal)
	rl.InitWindow(1600, 900, "Zivic Portable Desktop")
	defer rl.CloseWindow()
	//rl.SetTargetFPS(60)
	state.font = rl.LoadFontEx("resources/fonts/Mukta-ExtraBold.ttf", 500, nil, 0)
	for !rl.WindowShouldClose() {
		rl.DrawFPS(10, 10)
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		//SummonWIndow(&win, rl.Beige, rl.White, rl.Gray, rl.Brown, rl.White, true)
		SetState()
		rl.EndDrawing()
	}
}
