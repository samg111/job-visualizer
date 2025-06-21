package build

import (
	"job-visualizer/pkg/shared"
)

func BuildWindow(gui_data shared.GuiData) {
	mainWindow := gui_data.MainWindow
	jobs := gui_data.Jobs
	contentPane := buildMainWindow(jobs)

	mainWindow.SetContent(contentPane)
}
