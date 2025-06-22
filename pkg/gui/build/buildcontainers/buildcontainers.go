package buildcontainers

import (
	"fmt"
	"job-visualizer/pkg/gui/build/buildwidgets"
	"job-visualizer/pkg/shared"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildLeftSplit(jobs []shared.JobData) *container.Split {
	createJobList()
	refreshButton, filterButton, selectedDetailsButton := buildwidgets.BuildMainButtons(jobs)

	filterContainer, remoteCheckbox := buildFilterContainer()

	jobScroll := container.NewScroll(shared.Window.ListWidget)
	filterVBox := container.NewVBox(refreshButton, filterContainer, remoteCheckbox, filterButton)
	selectedDetailsContainer := container.NewBorder(nil, selectedDetailsButton, nil, nil, jobScroll)
	leftSplit := container.NewVSplit(filterVBox, selectedDetailsContainer)
	return leftSplit
}

func BuildRightSplit() *fyne.Container {
	detailsLabel := widget.NewLabelWithStyle("Select a job to display details", fyne.TextAlignLeading,
		fyne.TextStyle{Bold: false, Italic: false})
	detailsLabel.Wrapping = fyne.TextWrapWord
	shared.Window.DetailsWidget = detailsLabel
	rightPane := container.NewVBox(shared.Window.DetailsWidget)
	return rightPane
}

func createJobList() {
	getDataLen := func() int {
		if shared.Window.JobDataGui == nil {
			return 0
		}
		return len(*shared.Window.JobDataGui)
	}

	updateListItem := func(itemNum widget.ListItemID, listItem fyne.CanvasObject) {
		itemName := (*shared.Window.JobDataGui)[itemNum].CompanyName
		listItem.(*widget.Label).SetText(itemName)
	}
	shared.Window.ListWidget = widget.NewList(getDataLen, createListItem, updateListItem)
	shared.Window.ListWidget.OnSelected = func(i int) {
		shared.Window.SelectedJobDetails = formatJobDetails(i, shared.Window)
	}
}

func createListItem() fyne.CanvasObject {
	return widget.NewLabel("list items here")
}

func formatJobDetails(i int, window shared.GuiWindow) string {
	jobData := *window.JobDataGui
	job := jobData[i]
	formattedDetails := fmt.Sprintf("Company Name:\n\t%s\n\nJob Title:\n\t%s\n\nLocation:\n\t%s\n\nDate Posted:"+
		"\n\t%s\n\nSalary:\n\t%d\n\nWork From Home:\n\t%s\n\nQualifications:\n\t%s\n\nLinks:\n\t%s\n\n",
		job.CompanyName, job.JobTitle, job.Location, job.DatePosted, job.Salary, job.WorkFromHome, job.Qualifications,
		job.Links)
	return formattedDetails
}

func buildFilterContainer() (*fyne.Container, *widget.Check) {
	keywordContainer := buildKeywordContainer()
	locationContainer := buildLocationContainer()
	minSalaryContainer := buildMinSalaryContainer()
	filterContainer := container.NewVBox(keywordContainer, locationContainer, minSalaryContainer)
	remoteCheckbox := buildwidgets.BuildRemoteCheckbox()
	return filterContainer, remoteCheckbox
}

// func buildKeywordContainer() *fyne.Container {
// 	shared.Window.KeywordEntryWidget = widget.NewEntry()
// 	shared.Window.KeywordEntryWidget.SetPlaceHolder("Enter keyword filter here")
// 	keywordButton := widget.NewButton("Click to apply keyword", func() {
// 		shared.Window.Filters.KeywordEntry = shared.Window.KeywordEntryWidget.Text
// 	})
// 	keywordContainer := container.NewGridWithColumns(2, shared.Window.KeywordEntryWidget, keywordButton)
// 	return keywordContainer
// }

// func buildLocationContainer() *fyne.Container {
// 	shared.Window.LocationEntryWidget = widget.NewEntry()
// 	shared.Window.LocationEntryWidget.SetPlaceHolder("Enter location filter here")
// 	locationButton := widget.NewButton("Click to apply location", func() {
// 		shared.Window.Filters.LocationEntry = shared.Window.LocationEntryWidget.Text
// 	})
// 	locationContainer := container.NewGridWithColumns(2, shared.Window.LocationEntryWidget, locationButton)
// 	return locationContainer
// }

// func buildMinSalaryContainer() *fyne.Container {
// 	shared.Window.MinSalaryEntryWidget = widget.NewEntry()
// 	shared.Window.MinSalaryEntryWidget.SetPlaceHolder("Enter minimum salary filter here")
// 	minSalaryButton := widget.NewButton("Click to apply minimum salary", func() {
// 		shared.Window.Filters.MinSalaryEntry = shared.Window.MinSalaryEntryWidget.Text
// 	})
// 	minSalaryContainer := container.NewGridWithColumns(2, shared.Window.MinSalaryEntryWidget, minSalaryButton)
// 	return minSalaryContainer
// }

// func buildRemoteCheckbox() *widget.Check {
// 	remoteCheckbox := widget.NewCheck("Remote Work: check for yes, uncheck for all", func(checked bool) {
// 		if checked {
// 			shared.Window.Filters.WorkFromHomeEntry = true
// 		} else {
// 			shared.Window.Filters.WorkFromHomeEntry = false
// 		}
// 	})
// 	return remoteCheckbox
// }
