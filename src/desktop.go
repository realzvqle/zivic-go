package main

import rl "github.com/gen2brain/raylib-go/raylib"

var win WINDOW = WINDOW{x: 30, y: 30, title: "hi", sizeX: 300, sizeY: 400}

func Desktop() {
	SummonWindow(&win, rl.Beige, rl.White, rl.Gray, rl.Brown, rl.White, true)
}
