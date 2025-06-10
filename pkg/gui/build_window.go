package gui

import (
	"fmt"

	// "job-visualizer/pkg/gui/components"
	// "job-visualizer/pkg/gui/job_data"
	// "job-visualizer/pkg/gui/components"
	// "job-visualizer/pkg/gui/job_d"
	"job-visualizer/pkg/structs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	// "github.com/skratchdot/open-golang/open"
)

func buildWindow(gui_data GuiData) {
	// window := Window
	mainWindow := gui_data.mainWindow
	jobs := gui_data.jobs

	leftPane := buildLeftPane(jobs)
	rightPane := buildRightPane()
	contentPane := container.NewHSplit(leftPane, rightPane)

	mainWindow.SetContent(contentPane)
}

func buildLeftPane(jobs []structs.JobData) *container.Split {
	// dataLen:= len(*window.JobDataGui)
	getDataLen := func() int {
		if Window.JobDataGui == nil {
			return 0
		}
		return len(*Window.JobDataGui)
	}

	updateListItem := func(itemNum widget.ListItemID, listItem fyne.CanvasObject) {
		itemName := (*Window.JobDataGui)[itemNum].CompanyName
		listItem.(*widget.Label).SetText(itemName)
	}
	Window.ListWidget = widget.NewList(getDataLen, createListItem, updateListItem)
	Window.ListWidget.OnSelected = func(i int) {
		Window.SelectedJobDetails = formatJobDetails(i, Window)
	}

	topContainer := BuildTopLeftComponents(jobs)
	filterContainer, remoteCheckbox := BuildFilterComponents()
	dataButton := widget.NewButton("Click to filter the jobs", func() {
		GetJobData(jobs)
		// openWebpage()
	})
	topPane := container.NewVBox(topContainer, filterContainer, remoteCheckbox, dataButton)
	bottomPane := container.NewScroll(Window.ListWidget)
	leftPane := container.NewVSplit(topPane, bottomPane)
	return leftPane
}

// func getDataLength() int {
// 	if Window.JobDataGui == nil {
// 		return 0
// 	}
// 	return len(*Window.JobDataGui)
// }

func createListItem() fyne.CanvasObject {
	return widget.NewLabel("list items here")
}

// func updateListItem(window structs.GuiWindow, itemNum widget.ListItemID, listItem fyne.CanvasObject) {
// 	itemName := (*window.JobDataGui)[itemNum].CompanyName
// 	listItem.(*widget.Label).SetText(itemName)
// }

func formatJobDetails(i int, window structs.GuiWindow) string {
	jobData := *window.JobDataGui
	job := jobData[i]
	formattedDetails := fmt.Sprintf("Company Name:\n\t%s\n\nJob Title:\n\t%s\n\nLocation:\n\t%s\n\nDate Posted:"+
		"\n\t%s\n\nSalary:\n\t%s\n\nWork From Home:\n\t%s\n\nQualifications:\n\t%s\n\nLinks:\n\t%s\n\n",
		job.CompanyName, job.JobTitle, job.Location, job.DatePosted, job.Salary, job.WorkFromHome, job.Qualifications,
		job.Links)
	return formattedDetails
}

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
		Window.DetailsWidget.SetText(Window.SelectedJobDetails)
	})
	detailsLabel := widget.NewLabelWithStyle("Select a job to display details", fyne.TextAlignLeading,
		fyne.TextStyle{Bold: false, Italic: false})
	detailsLabel.Wrapping = fyne.TextWrapWord
	Window.DetailsWidget = detailsLabel
	rightPane := container.NewVBox(detailsButton, Window.DetailsWidget)
	return rightPane
}

// func buildRemoteCheckbox() *widget.Check {
// 	remoteCheckbox := widget.NewCheck("Remote Work: check for yes, uncheck for all", func(checked bool) {
// 		if checked {
// 			window.Filters.WorkFromHomeEntry = true
// 		} else {
// 			window.Filters.WorkFromHomeEntry = false
// 		}
// 	})
// 	return remoteCheckbox
// }

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
