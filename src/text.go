package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func ZiDrawText(text string, posX int32, posY int32, size int32, color rl.Color) {
	rl.DrawTextEx(state.font, text, rl.Vector2{X: float32(posX), Y: float32(posY)}, float32(size), 4, color)
}
