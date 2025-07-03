package gui

import (
	"job-visualizer/pkg/database"
	"job-visualizer/pkg/excel"
	"job-visualizer/pkg/gui/build"
	"job-visualizer/pkg/jobdata"
	"job-visualizer/pkg/jobdata/processing"
	"job-visualizer/pkg/mapping"
	"job-visualizer/pkg/shared"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	_ "modernc.org/sqlite"
)

var application fyne.App

// var StartWindow fyne.Window
// var MainWindow fyne.Window

func RunGUIorHeadless(headless bool) {
	if headless {
		println("temp removed headless functionality")
		// for i, job := range allJobData {
		// 	if i%100 == 0 {
		// 		fmt.Printf("%-4s | %-25s | %-55s | %-25s\n",
		// 			"#", "Location", "Job Title", "Company Name")
		// 		fmt.Println(strings.Repeat("-", 120))
		// 	}
		// 	fmt.Printf("%-4d | %-25s | %-55s | %-25s\n",
		// 		i+1, job.Location, job.JobTitle, job.CompanyName)
		// }
	} else {
		createGuiApp()
	}
}

func createGuiApp() {
	application = app.NewWithID("job-visualizer")
	progressBar := widget.NewProgressBar()
	progressBar.SetValue(0)
	startButton := widget.NewButton("Start Application", func() {
		go func() {
			files := excel.OpenExcelFile()
			rows := excel.GetAllRows(files)
			allJobData := jobdata.ProcessRows(rows, []shared.JobData{})
			allJobData = processing.ProcessLatLongs(allJobData, progressBar)
			allJobData = mapping.GenerateMap(allJobData)

			jobsDatabase := database.CreateDatabase()
			database.SetupDatabase(jobsDatabase)
			database.WriteToDatabase(jobsDatabase, allJobData)
			fyne.DoAndWait(func() {
				shared.MainWindow = createGuiWindow("job-visualizer")
				shared.MainWindow.SetOnClosed(func() { application.Quit() })
				shared.MainWindow = build.BuildMainWindow(shared.MainWindow, allJobData)
				shared.StartWindow.Hide()
				shared.MainWindow.Show()
			})
		}()
	})
	shared.StartWindow = createGuiWindow("job-visualizer")
	shared.StartWindow = build.BuildStartWindow(shared.StartWindow, startButton, progressBar)
	shared.StartWindow.ShowAndRun()
}

func createGuiWindow(title string) fyne.Window {
	// application := app.NewWithID("job-visualizer")
	Window := application.NewWindow(title)
	Window.Resize(fyne.NewSize(1000, 600))
	return Window
}
