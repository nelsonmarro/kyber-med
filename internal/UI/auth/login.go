package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type LoginPage struct {
	loginContainer *fyne.Container
}

func NewLoginPage() *LoginPage {
	return &LoginPage{}
}

func (l *LoginPage) CreateLoginPage() *fyne.Container {
	loginContainer := container.NewVBox(createPageElements()...)
	return loginContainer
}

func createPageElements() []fyne.CanvasObject {
	elemets := []fyne.CanvasObject{
		canvas.NewImageFromFile("../../../assets/plan.png"),
	}

	return elemets
}
