package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	mainWindow := a.NewWindow("Login")

	mainWindow.SetContent()
	mainWindow.ShowAndRun()
}
