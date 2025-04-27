package gui

import (
	"job-visualizer/pkg/structs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	// "github.com/skratchdot/open-golang/open"
)

func BuildTopLeftComponents(jobs []structs.JobData) *fyne.Container {
	mapButton := widget.NewButton("Click to open/refresh map", func() {})
	getJobsButton := widget.NewButton("Click get unfiltered jobs or refresh list of jobs/filters to original", func() {
		removeActiveFilters()
		// getJobData(jobs)
		// openWebpage()
		refreshEntries()
	})
	displayPane := container.NewVBox(getJobsButton, mapButton)
	return displayPane
}

func removeActiveFilters() {
	window.Filters.KeywordEntry = ""
	window.Filters.LocationEntry = ""
	window.Filters.MinSalaryEntry = ""
	window.Filters.WorkFromHomeEntry = false
}

func refreshEntries() {
	window.KeywordEntryWidget.SetText("")
	window.LocationEntryWidget.SetText("")
	window.MinSalaryEntryWidget.SetText("")
	window.KeywordEntryWidget.SetPlaceHolder("Enter keyword filter here")
	window.LocationEntryWidget.SetPlaceHolder("Enter location filter here")
	window.MinSalaryEntryWidget.SetPlaceHolder("Enter minimum salary filter here")
}
