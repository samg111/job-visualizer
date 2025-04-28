package gui

import (
	"fmt"
	"job-visualizer/pkg/structs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var window structs.GuiWindow

func CreateGui(jobs []structs.JobData) {
	mainWindow := createGuiWindow()
	// window = structs.GuiWindow{}
	BuildWindow(mainWindow, jobs)

	mainWindow.ShowAndRun()
}

func createGuiWindow() fyne.Window {
	application := app.New()
	window := application.NewWindow("fyne window")
	window.Resize(fyne.NewSize(1000, 600))
	return window
}

func checkErrorWarn(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
