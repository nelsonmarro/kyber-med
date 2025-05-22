package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
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
	imageLogo := canvas.NewImageFromFile("assets/plan.png")
	imageLogo.FillMode = canvas.ImageFillOriginal

	elemets := []fyne.CanvasObject{
		imageLogo,
		widget.NewEntry(),
		widget.NewEntry(),
		layout.NewSpacer(),
	}

	return elemets
}
