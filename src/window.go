package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type WINDOW struct {
	x                  float32
	y                  float32
	sizeX              float32
	sizeY              float32
	title              string
	isDragging         bool
	dragStopDelay      float32
	dragOffsetX        float32
	dragOffsetY        float32
	isResizing         bool
	resizeStopDelay    float32
	resizeOffsetX      float32
	resizeOffsetY      float32
	resizeButtonHeight float32
	resizeButtonGap    float32
	prevSizeX          float32
	prevSizeY          float32
	init               bool
}

func handleResize(win *WINDOW, windowColor rl.Color, windowColorHover rl.Color) {
	win.resizeButtonHeight = 30
	win.resizeButtonGap = 0

	resizeButton := ZiDrawButton("", win.x, win.y+win.sizeY+win.resizeButtonGap, win.sizeX, win.resizeButtonHeight, windowColor, rl.White, windowColorHover, -1)
	if resizeButton == BUTTON_DOWN_LEFT || win.isResizing {
		if !win.isResizing && rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			win.isResizing = true
			win.resizeOffsetX = float32(rl.GetMouseX()) - win.sizeX
			win.resizeOffsetY = float32(rl.GetMouseY()) - win.sizeY
		}

		if win.isResizing && rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			win.sizeX = float32(rl.GetMouseX()) - win.resizeOffsetX
			win.sizeY = float32(rl.GetMouseY()) - win.resizeOffsetY
		}
		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			win.isResizing = false
			win.resizeStopDelay = 0.5
		}

		if win.prevSizeX > win.sizeX {
			win.sizeX = win.prevSizeX
		}
		if win.prevSizeY > win.sizeY {
			win.sizeY = win.prevSizeY
		}
	}
}

func handleResizeStop(win *WINDOW) {
	if rl.IsMouseButtonReleased(rl.MouseLeftButton) || (win.isResizing && win.resizeStopDelay <= 0) {
		win.isResizing = false
		win.resizeStopDelay = 0.5
	}
}

func resize(win *WINDOW, windowColor rl.Color, windowColorHover rl.Color) {
	if win.prevSizeX > win.sizeX && win.prevSizeY > win.sizeY {
		handleResizeStop(win)
	} else {
		handleResize(win, windowColor, windowColorHover)
		//fmt.Printf("prevSizeX: %.2f, sizeX: %.2f, prevSizeY: %.2f, sizeY: %.2f\n",
		//	win.prevSizeX, win.sizeX, win.prevSizeY, win.sizeY)
	}

}

func SummonWindow(win *WINDOW, titlebarColor rl.Color, windowColor rl.Color, windowColorHover rl.Color, titlebarColorHover rl.Color, textColor rl.Color, isResizeable bool) {
	if !win.init {
		win.prevSizeX = win.sizeX
		win.prevSizeY = win.sizeY
		win.init = true
	}
	if isResizeable {
		resize(win, windowColor, windowColorHover)
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
