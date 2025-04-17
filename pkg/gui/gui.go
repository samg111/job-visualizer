package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func CreateGui() {
	mainWindow := createGuiWindow()
	mainWindow.ShowAndRun()
}

func createGuiWindow() fyne.Window {
	application := app.New()
	window := application.NewWindow("fyne window")
	window.Resize(fyne.NewSize(1000, 600))
	return window
}
