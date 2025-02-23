package main

import (
	"fyne.io/fyne/v2/app"

	ui "github.com/nelsonmarro/kyber-med/internal/UI/auth"
)

func main() {
	a := app.New()
	mainWindow := a.NewWindow("Login")

	loginPage := ui.NewLoginPage()

	mainWindow.SetContent(loginPage.CreateLoginPage())
	mainWindow.ShowAndRun()
}
