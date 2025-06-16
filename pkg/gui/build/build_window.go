package build

import (
	"job-visualizer/pkg/shared"

	"fyne.io/fyne/v2/container"
)

func BuildWindow(gui_data shared.GuiData) {
	mainWindow := gui_data.MainWindow
	jobs := gui_data.Jobs

	leftSplit := buildLeftSplit(jobs)
	rightSplit := buildRightSplit()
	contentPane := container.NewHSplit(leftSplit, rightSplit)

	mainWindow.SetContent(contentPane)
}
