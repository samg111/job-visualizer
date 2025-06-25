package build

import (
	"job-visualizer/pkg/gui/build/buildcontainers"
	"job-visualizer/pkg/shared"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func BuildStartWindow(startWindow fyne.Window, mainWindow fyne.Window) fyne.Window {
	startContainer := buildcontainers.BuildStartContainer(startWindow, mainWindow)
	startWindow.SetContent(startContainer)
	return startWindow
}

func BuildMainWindow(gui_data shared.GuiData) fyne.Window {
	window := gui_data.MainWindow
	jobs := gui_data.Jobs
	contentPane := buildMainContent(jobs)
	window.SetContent(contentPane)
	return window
}

func buildMainContent(jobs []shared.JobData) *container.Split {
	leftSplit := buildcontainers.BuildLeftSplit(jobs)
	rightSplit := buildcontainers.BuildRightSplit()
	contentPane := container.NewHSplit(leftSplit, rightSplit)

	return contentPane
}
