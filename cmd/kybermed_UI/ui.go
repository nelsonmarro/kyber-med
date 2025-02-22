package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	mainWindow := a.NewWindow("Hello World")

	mainWindow.SetContent(widget.NewLabel("Hello Wolrd!"))
	mainWindow.ShowAndRun()
}
