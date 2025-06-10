package gui

import (
	"fyne.io/fyne/v2/container"
)

func buildWindow(gui_data GuiData) {
	mainWindow := gui_data.mainWindow
	jobs := gui_data.jobs

	leftPane := buildLeftSplit(jobs)
	rightPane := buildRightSplit()
	contentPane := container.NewHSplit(leftPane, rightPane)

	mainWindow.SetContent(contentPane)
}
