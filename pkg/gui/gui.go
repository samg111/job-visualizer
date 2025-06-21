package gui

import (
	"fmt"
	"job-visualizer/pkg/gui/build"
	"job-visualizer/pkg/shared"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func RunGUIorHeadless(headless bool, allJobData []shared.JobData) {
	if headless {
		for i, job := range allJobData {
			if i%100 == 0 {
				fmt.Printf("%-4s | %-25s | %-55s | %-25s\n",
					"#", "Location", "Job Title", "Company Name")
				fmt.Println(strings.Repeat("-", 120))
			}
			fmt.Printf("%-4d | %-25s | %-55s | %-25s\n",
				i+1, job.Location, job.JobTitle, job.CompanyName)
		}
	} else {
		createGui(allJobData)
	}
}

func createGui(jobs []shared.JobData) {
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
