package build

import (
	"fmt"
	"job-visualizer/pkg/gui/buttons"
	"job-visualizer/pkg/shared"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func buildLeftSplit(jobs []shared.JobData) *container.Split {
	createJobList()
	topContainer := buildTopLeftComponents(jobs)
	filterContainer, remoteCheckbox := buildFilterComponents()
	dataButton := createDataButton(jobs)
	topPane := container.NewVBox(topContainer, filterContainer, remoteCheckbox, dataButton)
	bottomPane := container.NewScroll(shared.Window.ListWidget)
	leftSplit := container.NewVSplit(topPane, bottomPane)
	return leftSplit
}

func buildRightSplit() *fyne.Container {
	detailsButton := widget.NewButton("Click to display selected job details", func() {
		shared.Window.DetailsWidget.SetText(shared.Window.SelectedJobDetails)
	})
	detailsLabel := widget.NewLabelWithStyle("Select a job to display details", fyne.TextAlignLeading,
		fyne.TextStyle{Bold: false, Italic: false})
	detailsLabel.Wrapping = fyne.TextWrapWord
	shared.Window.DetailsWidget = detailsLabel
	rightPane := container.NewVBox(detailsButton, shared.Window.DetailsWidget)
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

func createDataButton(jobs []shared.JobData) *widget.Button {
	dataButton := widget.NewButton("Click to filter the jobs", func() {
		buttons.HandleJobFilter(jobs)
	})
	return dataButton
}
