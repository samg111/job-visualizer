package gui

import (
	"fmt"
	"job-visualizer/pkg/structs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func buildLeftSplit(jobs []structs.JobData) *container.Split {
	createJobList()
	topContainer := buildTopLeftComponents(jobs)
	filterContainer, remoteCheckbox := buildFilterComponents()
	dataButton := createDataButton(jobs)
	topPane := container.NewVBox(topContainer, filterContainer, remoteCheckbox, dataButton)
	bottomPane := container.NewScroll(Window.ListWidget)
	leftSplit := container.NewVSplit(topPane, bottomPane)
	return leftSplit
}

func buildRightSplit() *fyne.Container {
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

func createJobList() {
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
}

func createListItem() fyne.CanvasObject {
	return widget.NewLabel("list items here")
}

func formatJobDetails(i int, window structs.GuiWindow) string {
	jobData := *window.JobDataGui
	job := jobData[i]
	formattedDetails := fmt.Sprintf("Company Name:\n\t%s\n\nJob Title:\n\t%s\n\nLocation:\n\t%s\n\nDate Posted:"+
		"\n\t%s\n\nSalary:\n\t%d\n\nWork From Home:\n\t%s\n\nQualifications:\n\t%s\n\nLinks:\n\t%s\n\n",
		job.CompanyName, job.JobTitle, job.Location, job.DatePosted, job.Salary, job.WorkFromHome, job.Qualifications,
		job.Links)
	return formattedDetails
}

func createDataButton(jobs []structs.JobData) *widget.Button {
	dataButton := widget.NewButton("Click to filter the jobs", func() {
		GetJobData(jobs)
		// openWebpage()
	})
	return dataButton
}
