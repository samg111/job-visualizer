package gui

import (
	"job-visualizer/pkg/shared"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// var Window shared.GuiWindow

type GuiData struct {
	mainWindow fyne.Window
	jobs       []shared.JobData
}

func CreateGui(jobs []shared.JobData) {
	mainWindow := createGuiWindow()
	gui_data := creatGuiData(mainWindow, jobs)
	buildWindow(gui_data)
	mainWindow.ShowAndRun()
}

func createGuiWindow() fyne.Window {
	application := app.New()
	Window := application.NewWindow("fyne window")
	Window.Resize(fyne.NewSize(1000, 600))
	return Window
}

func creatGuiData(mainWindow fyne.Window, jobs []shared.JobData) GuiData {
	gui_data := GuiData{
		mainWindow: mainWindow,
		jobs:       jobs,
	}
	return gui_data
}
