package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type window struct {
	x             float32
	y             float32
	sizeX         float32
	sizeY         float32
	title         string
	isDragging    bool
	dragStopDelay float32
	dragOffsetX   float32
	dragOffsetY   float32
	prevSizeX     float32
	prevSizeY     float32
	init          bool
}

func SummonWIndow(win *window, titlebarColor rl.Color, windowColor rl.Color, windowColorHover rl.Color, titlebarColorHover rl.Color, textColor rl.Color, isResizeable bool) {
	if !win.init {
		win.prevSizeX = win.sizeX
		win.prevSizeY = win.sizeY
		win.init = true
	}

	ZiDrawButton("", win.x, win.y, win.sizeX, win.sizeY, windowColor, rl.White, windowColorHover, -1)
	bar := ZiDrawButton(win.title, win.x, win.y-30, win.sizeX, 30, titlebarColor, rl.White, titlebarColorHover, -1)
	if bar == BUTTON_DOWN_LEFT || win.isDragging {

		if !win.isDragging {
			win.isDragging = true
			win.dragOffsetX = float32(rl.GetMouseX()) - win.x
			win.dragOffsetY = float32(rl.GetMouseY()) - win.y
		}
		if win.isDragging && rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			//fmt.Println("Hi!")
			win.x = float32(rl.GetMouseX()) - win.dragOffsetX
			win.y = float32(rl.GetMouseY()) - win.dragOffsetY
			win.dragStopDelay = win.dragStopDelay - rl.GetFrameTime()
		}
	}
	if rl.IsMouseButtonReleased(rl.MouseButtonLeft) || (win.isDragging && win.dragStopDelay <= 0) {
		win.isDragging = false
		win.dragStopDelay = 0.5
	}
}
