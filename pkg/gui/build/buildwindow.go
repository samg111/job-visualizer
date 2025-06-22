package build

import (
	"job-visualizer/pkg/gui/build/buildcontainers"
	"job-visualizer/pkg/shared"

	"fyne.io/fyne/v2/container"
)

func BuildWindow(gui_data shared.GuiData) {
	mainWindow := gui_data.MainWindow
	jobs := gui_data.Jobs
	contentPane := buildMainWindow(jobs)

	mainWindow.SetContent(contentPane)
}

func buildMainWindow(jobs []shared.JobData) *container.Split {
	leftSplit := buildcontainers.BuildLeftSplit(jobs)
	rightSplit := buildcontainers.BuildRightSplit()
	contentPane := container.NewHSplit(leftSplit, rightSplit)

	return contentPane
}
