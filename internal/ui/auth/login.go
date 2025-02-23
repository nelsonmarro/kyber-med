package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	imageLogo.Resize(fyne.Size{Width: 64, Height: 64})
	imageLogo.FillMode = canvas.ImageFillOriginal

	elemets := []fyne.CanvasObject{
		layout.NewSpacer(),
		imageLogo,
		layout.NewSpacer(),
	}

	return elemets
}
