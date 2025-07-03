package buildcontainers

import (
	"fmt"
	"job-visualizer/pkg/gui/build/buildwidgets"
	"job-visualizer/pkg/shared"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildStartContainer(window fyne.Window, startButton *widget.Button, progressBar *widget.ProgressBar) *fyne.Container {
	startLabel := buildwidgets.BuildLabel("Welcome to job-visualizer, choose your input files and output file location",
		true, false)
	inputFileLabel := buildwidgets.BuildLabel("No files selected", false, false)
	inputFileButton := buildwidgets.BuildStartButtons(window, inputFileLabel)
	// progressBar := widget.NewProgressBar()
	// progressBar.SetValue(0)

	return container.NewVBox(startLabel, inputFileLabel, inputFileButton, startButton, progressBar)
}

func BuildLeftSplit(jobs []shared.JobData) *container.Split {
	createJobList()
	refreshButton, filterButton, selectedDetailsButton := buildwidgets.BuildMainButtons(jobs)
	selectedDetailsLabel := buildwidgets.BuildLabel("Select a job and click button to display details",
		true, false)

	filterContainer := buildFilterContainer()

	jobScroll := container.NewScroll(shared.Window.ListWidget)
	filterVBox := container.NewVBox(refreshButton, filterContainer, filterButton)
	selectedDetailsContainer := container.NewBorder(
		selectedDetailsLabel,
		selectedDetailsButton,
		nil,
		nil,
		jobScroll)
	leftSplit := container.NewVSplit(filterVBox, selectedDetailsContainer)
	return leftSplit
}

func BuildRightSplit() *fyne.Container {
	detailsLabel := buildwidgets.BuildLabel("Select a job to display details", true, false)
	detailsLabel.Wrapping = fyne.TextWrapWord
	shared.Window.DetailsWidget = detailsLabel
	rightPane := container.NewVBox(shared.Window.DetailsWidget)
	return rightPane
}

func createJobList() {
	getDataLen := func() int {
		if shared.Window.FilteredJobs == nil {
			return 0
		}
		return len(*shared.Window.FilteredJobs)
	}

	updateListItem := func(itemNum widget.ListItemID, listItem fyne.CanvasObject) {
		itemName := (*shared.Window.FilteredJobs)[itemNum].CompanyName
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
	jobData := *window.FilteredJobs
	job := jobData[i]
	formattedDetails := fmt.Sprintf("Company Name:\n%s\n\nJob Title:\n%s\n\nLocation:\n%s\n\nDate Posted:"+
		"\n%s\n\nSalary:\n%d\n\nWork From Home:\n%s\n\nQualifications:\n%s\n\nLinks:\n%s\n\n",
		job.CompanyName, job.JobTitle, job.Location, job.DatePosted, job.Salary, job.WorkFromHome, job.Qualifications,
		job.Links)
	return formattedDetails
}

func buildFilterContainer() *fyne.Container {
	keywordContainer := buildKeywordContainer()
	locationContainer := buildLocationContainer()
	minSalaryContainer := buildMinSalaryContainer()
	remoteCheckbox := buildwidgets.BuildRemoteCheckbox()
	filterContainer := container.NewVBox(keywordContainer, locationContainer, minSalaryContainer, remoteCheckbox)
	return filterContainer
}
