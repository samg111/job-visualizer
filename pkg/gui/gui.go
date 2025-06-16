package gui

import (
	"job-visualizer/pkg/gui/build"
	"job-visualizer/pkg/shared"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func CreateGui(jobs []shared.JobData) {
	mainWindow := createGuiWindow()
	gui_data := creatGuiData(mainWindow, jobs)
	build.BuildWindow(gui_data)
	mainWindow.ShowAndRun()
}

func createGuiWindow() fyne.Window {
	application := app.New()
	Window := application.NewWindow("fyne window")
	Window.Resize(fyne.NewSize(1000, 600))
	return Window
}

func creatGuiData(mainWindow fyne.Window, jobs []shared.JobData) shared.GuiData {
	gui_data := shared.GuiData{
		MainWindow: mainWindow,
		Jobs:       jobs,
	}
	return gui_data
}
