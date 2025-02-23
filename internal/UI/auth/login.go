package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func CreateLoginPage() *fyne.Container {
	loginContainer := container.NewVBox(createPageElements())
	return l.loginContainer
}

func createPageElements() fyne.CanvasObject {
	panic("unimplemented")
}
