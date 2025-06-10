package gui

import (
	"fyne.io/fyne/v2/container"
)

func buildWindow(gui_data GuiData) {
	mainWindow := gui_data.mainWindow
	jobs := gui_data.jobs

	leftSplit := buildLeftSplit(jobs)
	rightSplit := buildRightSplit()
	contentPane := container.NewHSplit(leftSplit, rightSplit)

	mainWindow.SetContent(contentPane)
}
