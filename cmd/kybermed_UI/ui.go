package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	ui "github.com/nelsonmarro/kyber-med/internal/ui/auth"
)

func main() {
	a := app.New()
	mainWindow := a.NewWindow("Login")

	loginPage := ui.NewLoginPage()

	mainWindow.SetContent(loginPage.CreateLoginPage())
	mainWindow.Resize(fyne.Size{Width: 600, Height: 400})

	mainWindow.ShowAndRun()
}
