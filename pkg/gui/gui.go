package gui

import (
	"fmt"
	"job-visualizer/pkg/structs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var Window structs.GuiWindow

type GuiData struct {
	mainWindow fyne.Window
	jobs       []structs.JobData
}

func CreateGui(jobs []structs.JobData) {
	// var window structs.GuiWindow
	mainWindow := createGuiWindow()
	// guiWindow = structs.GuiWindow{}
	gui_data := GuiData{
		mainWindow: mainWindow,
		jobs:       jobs,
	}
	buildWindow(gui_data)

	mainWindow.ShowAndRun()
}

func createGuiWindow() fyne.Window {
	application := app.New()
	Window := application.NewWindow("fyne window")
	Window.Resize(fyne.NewSize(1000, 600))
	return Window
}

func checkErrorWarn(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
