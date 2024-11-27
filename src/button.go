package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ButtonState int

const (
	BUTTON_NOT_INTERACTED ButtonState = iota
	BUTTON_HOVER
	BUTTON_CLICK_LEFT
	BUTTON_CLICK_RIGHT
	BUTTON_DOWN_LEFT
	BUTTON_DOWN_RIGHT
)

func ZiDrawButton(text string, x float32, y float32, sizeX float32, sizeY float32, baseColor rl.Color, textColor rl.Color, hoverColor rl.Color, textsize float32) ButtonState {
	if textsize == -1 {
		textsize = sizeY / 1.5
	}

	rec := rl.Rectangle{X: x, Y: y, Width: sizeX, Height: sizeY}
	isMouseOver := rl.CheckCollisionPointRec(rl.GetMousePosition(), rec)
	var buttonColor rl.Color

	if isMouseOver {
		buttonColor = hoverColor
	} else {
		buttonColor = baseColor
	}

	rl.DrawRectangle(int32(x), int32(y), int32(sizeX), int32(sizeY), buttonColor)
	var textWidth int32 = 0

	for i := 0; i < len(text); i++ {
		textWidth += rl.MeasureText(string(text[i]), int32(textsize))
	}

	textX := x + (sizeX-rl.MeasureTextEx(state.font, text, textsize, 4).X)/2
	textY := y + (sizeY-textsize)/2

	ZiDrawText(text, int32(textX), int32(textY), int32(textsize), textColor)

	if isMouseOver && rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		rec := rl.Rectangle{X: x, Y: y, Width: sizeX, Height: sizeY}
		rl.DrawRectangleLinesEx(rec, 2, textColor)
		return BUTTON_DOWN_LEFT
	}

	if isMouseOver && rl.IsMouseButtonDown(rl.MouseButtonRight) {
		rec := rl.Rectangle{X: x, Y: y, Width: sizeX, Height: sizeY}
		rl.DrawRectangleLinesEx(rec, 2, textColor)
		return BUTTON_DOWN_LEFT
	}

	if isMouseOver && rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		rec := rl.Rectangle{X: x, Y: y, Width: sizeX, Height: sizeY}
		rl.DrawRectangleLinesEx(rec, 2, textColor)
		return BUTTON_CLICK_LEFT
	}

	if isMouseOver && rl.IsMouseButtonReleased(rl.MouseButtonRight) {
		rec := rl.Rectangle{X: x, Y: y, Width: sizeX, Height: sizeY}
		rl.DrawRectangleLinesEx(rec, 2, textColor)
		return BUTTON_CLICK_RIGHT
	}

	if isMouseOver {
		rec := rl.Rectangle{X: x, Y: y, Width: sizeX, Height: sizeY}
		rl.DrawRectangleLinesEx(rec, 2, textColor)
		return BUTTON_NOT_INTERACTED
	}
	return BUTTON_NOT_INTERACTED
}
