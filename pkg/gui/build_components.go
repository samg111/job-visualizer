package gui

import (
	"job-visualizer/pkg/structs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func buildTopLeftComponents(jobs []structs.JobData) *fyne.Container {
	mapButton := widget.NewButton("Click to open/refresh map", func() {})
	getJobsButton := widget.NewButton("Click get the unfiltered jobs or refresh list of jobs/filters to original", func() {
		removeActiveFilters()
		GetJobData(jobs)
		// openWebpage()
		refreshEntries()
	})
	displayPane := container.NewVBox(getJobsButton, mapButton)
	return displayPane
}

func buildFilterComponents() (*fyne.Container, *widget.Check) {
	keywordContainer := buildKeywordContainer()
	locationContainer := buildLocationContainer()
	minSalaryContainer := buildMinSalaryContainer()
	filterContainer := container.NewVBox(keywordContainer, locationContainer, minSalaryContainer)
	remoteCheckbox := buildRemoteCheckbox()
	return filterContainer, remoteCheckbox
}

func removeActiveFilters() {
	Window.Filters.KeywordEntry = ""
	Window.Filters.LocationEntry = ""
	Window.Filters.MinSalaryEntry = ""
	Window.Filters.WorkFromHomeEntry = false
}

func refreshEntries() {
	Window.KeywordEntryWidget.SetText("")
	Window.LocationEntryWidget.SetText("")
	Window.MinSalaryEntryWidget.SetText("")
	Window.KeywordEntryWidget.SetPlaceHolder("Enter keyword filter here")
	Window.LocationEntryWidget.SetPlaceHolder("Enter location filter here")
	Window.MinSalaryEntryWidget.SetPlaceHolder("Enter minimum salary filter here")
}

func buildKeywordContainer() *fyne.Container {
	Window.KeywordEntryWidget = widget.NewEntry()
	Window.KeywordEntryWidget.SetPlaceHolder("Enter keyword filter here")
	keywordButton := widget.NewButton("Click to apply keyword", func() {
		Window.Filters.KeywordEntry = Window.KeywordEntryWidget.Text
	})
	keywordContainer := container.NewGridWithColumns(2, Window.KeywordEntryWidget, keywordButton)
	return keywordContainer
}

func buildLocationContainer() *fyne.Container {
	Window.LocationEntryWidget = widget.NewEntry()
	Window.LocationEntryWidget.SetPlaceHolder("Enter location filter here")
	locationButton := widget.NewButton("Click to apply location", func() {
		Window.Filters.LocationEntry = Window.LocationEntryWidget.Text
	})
	locationContainer := container.NewGridWithColumns(2, Window.LocationEntryWidget, locationButton)
	return locationContainer
}

func buildMinSalaryContainer() *fyne.Container {
	Window.MinSalaryEntryWidget = widget.NewEntry()
	Window.MinSalaryEntryWidget.SetPlaceHolder("Enter minimum salary filter here")
	minSalaryButton := widget.NewButton("Click to apply minimum salary", func() {
		Window.Filters.MinSalaryEntry = Window.MinSalaryEntryWidget.Text
	})
	minSalaryContainer := container.NewGridWithColumns(2, Window.MinSalaryEntryWidget, minSalaryButton)
	return minSalaryContainer
}

func buildRemoteCheckbox() *widget.Check {
	remoteCheckbox := widget.NewCheck("Remote Work: check for yes, uncheck for all", func(checked bool) {
		if checked {
			Window.Filters.WorkFromHomeEntry = true
		} else {
			Window.Filters.WorkFromHomeEntry = false
		}
	})
	return remoteCheckbox
}
