package gui

import (
	"fmt"

	"job-visualizer/pkg/structs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	// "github.com/skratchdot/open-golang/open"
)

func BuildWindow(mainWindow fyne.Window, jobs []structs.JobData) {
	leftPane := buildLeftPane(jobs)
	rightPane := buildRightPane()
	contentPane := container.NewHSplit(leftPane, rightPane)

	mainWindow.SetContent(contentPane)
}

func buildLeftPane(jobs []structs.JobData) *container.Split {
	window.ListWidget = widget.NewList(getDataLength, createListItem, updateListItem)
	window.ListWidget.OnSelected = func(i int) {
		window.SelectedJobDetails = formatJobDetails(i)
	}

	topContainer := BuildTopLeftComponents(jobs)
	// filterContainer, remoteCheckbox := buildFilterComponents()
	// dataButton := widget.NewButton("Click to filter the jobs", func() {
	// 	getJobData(jobs)
	// 	// openWebpage()
	// })
	topPane := container.NewVBox(topContainer)
	bottomPane := container.NewVBox(window.ListWidget)
	// topPane := container.NewVBox(topLeftContainer, filterContainer, remoteCheckbox, dataButton)
	leftPane := container.NewVSplit(topPane, bottomPane)
	return leftPane
}

func getDataLength() int {
	if window.JobDataGui == nil {
		return 0
	}
	return len(*window.JobDataGui)
}

func createListItem() fyne.CanvasObject {
	return widget.NewLabel("list items here")
}

func updateListItem(itemNum widget.ListItemID, listItem fyne.CanvasObject) {
	itemName := (*window.JobDataGui)[itemNum].CompanyName
	listItem.(*widget.Label).SetText(itemName)
}

func formatJobDetails(i int) string {
	jobData := *window.JobDataGui
	job := jobData[i]
	formattedDetails := fmt.Sprintf("Company Name:\n\t%s\n\nJob Title:\n\t%s\n\nLocation:\n\t%s\n\nDate Posted:"+
		"\n\t%s\n\nSalary:\n\t%s\n\nWork From Home:\n\t%s\n\nQualifications:\n\t%s\n\nLinks:\n\t%s\n\n",
		job.CompanyName, job.JobTitle, job.Location, job.DatePosted, job.Salary, job.WorkFromHome, job.Qualifications,
		job.Links)
	return formattedDetails
}

// func buildTopLeftComponents(jobs []structs.JobData) *fyne.Container {
// 	mapButton := widget.NewButton("Click to open/refresh map", func() {})
// 	getJobsButton := widget.NewButton("Click get unfiltered jobs or refresh list of jobs/filters to original", func() {
// 		removeActiveFilters()
// 		getJobData(jobs)
// 		// openWebpage()
// 		refreshEntries()
// 	})
// 	displayPane := container.NewVBox(getJobsButton, mapButton)
// 	return displayPane
// }

// func removeActiveFilters() {
// 	window.Filters.KeywordEntry = ""
// 	window.Filters.LocationEntry = ""
// 	window.Filters.MinSalaryEntry = ""
// 	window.Filters.WorkFromHomeEntry = false
// }

// func getJobData(jobs []structs.JobData) {
// 	window.JobDataGui = GetJobData(jobs)
// }

// func refreshEntries() {
// 	window.KeywordEntryWidget.SetText("")
// 	window.LocationEntryWidget.SetText("")
// 	window.MinSalaryEntryWidget.SetText("")
// 	window.KeywordEntryWidget.SetPlaceHolder("Enter keyword filter here")
// 	window.LocationEntryWidget.SetPlaceHolder("Enter location filter here")
// 	window.MinSalaryEntryWidget.SetPlaceHolder("Enter minimum salary filter here")
// }

// func openWebpage() {
// 	url := fmt.Sprintf("http://localhost:8080/%d", serverCount)
// 	err := open.Run(url)
// 	checkErrorWarn(err)
// }

func buildRightPane() *fyne.Container {
	detailsButton := widget.NewButton("Click to display selected job details", func() {
		window.DetailsWidget.SetText(window.SelectedJobDetails)
	})
	detailsLabel := widget.NewLabelWithStyle("Select a job to display details", fyne.TextAlignLeading,
		fyne.TextStyle{Bold: false, Italic: false})
	detailsLabel.Wrapping = fyne.TextWrapWord
	window.DetailsWidget = detailsLabel
	rightPane := container.NewVBox(detailsButton, window.DetailsWidget)
	return rightPane
}

// func buildFilterComponents() (*fyne.Container, *widget.Check) {
// 	keywordContainer := buildKeywordContainer()
// 	locationContainer := buildLocationContainer()
// 	minSalaryContainer := buildMinSalaryContainer()
// 	filterContainer := container.NewVBox(keywordContainer, locationContainer, minSalaryContainer)
// 	remoteCheckbox := buildRemoteCheckbox()
// 	return filterContainer, remoteCheckbox
// }

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

// func buildMinSalaryContainer() *fyne.Container {
// 	window.MinSalaryEntryWidget = widget.NewEntry()
// 	window.MinSalaryEntryWidget.SetPlaceHolder("Enter minimum salary filter here")
// 	minSalaryButton := widget.NewButton("Click to apply minimum salary", func() {
// 		window.Filters.MinSalaryEntry = window.MinSalaryEntryWidget.Text
// 	})
// 	minSalaryContainer := container.NewGridWithColumns(2, window.MinSalaryEntryWidget, minSalaryButton)
// 	return minSalaryContainer
// }

// func buildLocationContainer() *fyne.Container {
// 	window.LocationEntryWidget = widget.NewEntry()
// 	window.LocationEntryWidget.SetPlaceHolder("Enter location filter here")
// 	locationButton := widget.NewButton("Click to apply location", func() {
// 		window.Filters.LocationEntry = window.LocationEntryWidget.Text
// 	})
// 	locationContainer := container.NewGridWithColumns(2, window.LocationEntryWidget, locationButton)
// 	return locationContainer
// }

// func buildKeywordContainer() *fyne.Container {
// 	window.KeywordEntryWidget = widget.NewEntry()
// 	window.KeywordEntryWidget.SetPlaceHolder("Enter keyword filter here")
// 	keywordButton := widget.NewButton("Click to apply keyword", func() {
// 		window.Filters.KeywordEntry = window.KeywordEntryWidget.Text
// 	})
// 	keywordContainer := container.NewGridWithColumns(2, window.KeywordEntryWidget, keywordButton)
// 	return keywordContainer
// }
