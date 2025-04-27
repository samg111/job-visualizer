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
		getJobData(jobs)
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

func getJobData(jobs []structs.JobData) {
	window.JobDataGui = &jobs
}

// func GetJobData(jobs []structs.JobData) *[]structs.JobData {
// 	// jobs = filterJobs(jobs)
// 	// jobs = assignLatLongs(jobs)
// 	// geoplotMap := createGeoplotMap(jobs)
// 	// window.Server = createHttpServer(geoplotMap)
// 	return &jobs
// }

func refreshEntries() {
	window.KeywordEntryWidget.SetText("")
	window.LocationEntryWidget.SetText("")
	window.MinSalaryEntryWidget.SetText("")
	window.KeywordEntryWidget.SetPlaceHolder("Enter keyword filter here")
	window.LocationEntryWidget.SetPlaceHolder("Enter location filter here")
	window.MinSalaryEntryWidget.SetPlaceHolder("Enter minimum salary filter here")
}

func BuildFilterComponents() (*fyne.Container, *widget.Check) {
	keywordContainer := buildKeywordContainer()
	locationContainer := buildLocationContainer()
	minSalaryContainer := buildMinSalaryContainer()
	filterContainer := container.NewVBox(keywordContainer, locationContainer, minSalaryContainer)
	remoteCheckbox := buildRemoteCheckbox()
	return filterContainer, remoteCheckbox
}

func buildKeywordContainer() *fyne.Container {
	window.KeywordEntryWidget = widget.NewEntry()
	window.KeywordEntryWidget.SetPlaceHolder("Enter keyword filter here")
	keywordButton := widget.NewButton("Click to apply keyword", func() {
		window.Filters.KeywordEntry = window.KeywordEntryWidget.Text
	})
	keywordContainer := container.NewGridWithColumns(2, window.KeywordEntryWidget, keywordButton)
	return keywordContainer
}

func buildLocationContainer() *fyne.Container {
	window.LocationEntryWidget = widget.NewEntry()
	window.LocationEntryWidget.SetPlaceHolder("Enter location filter here")
	locationButton := widget.NewButton("Click to apply location", func() {
		window.Filters.LocationEntry = window.LocationEntryWidget.Text
	})
	locationContainer := container.NewGridWithColumns(2, window.LocationEntryWidget, locationButton)
	return locationContainer
}

func buildMinSalaryContainer() *fyne.Container {
	window.MinSalaryEntryWidget = widget.NewEntry()
	window.MinSalaryEntryWidget.SetPlaceHolder("Enter minimum salary filter here")
	minSalaryButton := widget.NewButton("Click to apply minimum salary", func() {
		window.Filters.MinSalaryEntry = window.MinSalaryEntryWidget.Text
	})
	minSalaryContainer := container.NewGridWithColumns(2, window.MinSalaryEntryWidget, minSalaryButton)
	return minSalaryContainer
}

func buildRemoteCheckbox() *widget.Check {
	remoteCheckbox := widget.NewCheck("Remote Work: check for yes, uncheck for all", func(checked bool) {
		if checked {
			window.Filters.WorkFromHomeEntry = true
		} else {
			window.Filters.WorkFromHomeEntry = false
		}
	})
	return remoteCheckbox
}
