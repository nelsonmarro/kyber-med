package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")
	rect := canvas.NewRectangle(color.NRGBA{R: 255, G: 0, B: 0, A: 255})
	top := container.NewHBox(rect)
	left := canvas.NewText("left", color.White)
	middle := canvas.NewText("content", color.White)
	content := container.NewBorder(top, nil, left, nil, middle)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
