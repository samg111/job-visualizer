package buildwidgets

import (
	"fmt"
	"job-visualizer/pkg/jobdata/filter"
	"job-visualizer/pkg/mapping"
	"job-visualizer/pkg/shared"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func BuildMainButtons(jobs []shared.JobData) (*widget.Button, *widget.Button, *widget.Button) {
	refreshButton := widget.NewButton("Click to refresh list of jobs to original", func() {
		handleJobRefresh(jobs)
	})
	filterButton := widget.NewButton("Click to filter the jobs", func() {
		handleJobFilter(jobs)
	})
	selectedDetailsButton := widget.NewButton("Click to display selected job details", func() {
		shared.Window.DetailsWidget.SetText(shared.Window.SelectedJobDetails)
	})

	return refreshButton, filterButton, selectedDetailsButton
}

func BuildStartButtons(window fyne.Window, inputFileLabel *widget.Label) *widget.Button {
	inputFileButton := widget.NewButton("Select Input Files", func() {
		fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			shared.CheckErrorWarn(err)
			if reader == nil {
				println("user cancelled file selection")
				return
			}
			defer reader.Close()
			selectedFile := reader.URI().Path()
			inputFileLabel.SetText(fmt.Sprintf("Selected file: %s", selectedFile))
		}, window)
		fileDialog.Show()
	})
	return inputFileButton
}

func BuildLabel(text string, boldBool bool, italicBool bool) *widget.Label {
	return widget.NewLabelWithStyle(text, fyne.TextAlignCenter,
		fyne.TextStyle{Bold: boldBool, Italic: italicBool})
}

func BuildRemoteCheckbox() *widget.Check {
	remoteCheckbox := widget.NewCheck("Remote Work: check for yes, uncheck for all", func(checked bool) {
		if checked {
			shared.Window.Filters.WorkFromHomeEntry = true
		} else {
			shared.Window.Filters.WorkFromHomeEntry = false
		}
	})
	return remoteCheckbox
}

func handleJobRefresh(jobs []shared.JobData) {
	removeActiveFilters()
	filteredJobs := filter.FilterJobs(jobs)
	mapping.GenerateMap(filteredJobs)
	shared.Window.FilteredJobs = &filteredJobs
	refreshEntries()
}

func handleJobFilter(jobs []shared.JobData) {
	filteredJobs := filter.FilterJobs(jobs)
	// jobs = jobdata.GetJobData(jobs)
	mapping.GenerateMap(filteredJobs)
	shared.Window.FilteredJobs = &filteredJobs
}

func removeActiveFilters() {
	shared.Window.Filters.KeywordEntry = ""
	shared.Window.Filters.LocationEntry = ""
	shared.Window.Filters.MinSalaryEntry = ""
	shared.Window.Filters.WorkFromHomeEntry = false
}

func refreshEntries() {
	shared.Window.KeywordEntryWidget.SetText("")
	shared.Window.LocationEntryWidget.SetText("")
	shared.Window.MinSalaryEntryWidget.SetText("")
	shared.Window.KeywordEntryWidget.SetPlaceHolder("Enter keyword filter here")
	shared.Window.LocationEntryWidget.SetPlaceHolder("Enter location filter here")
	shared.Window.MinSalaryEntryWidget.SetPlaceHolder("Enter minimum salary filter here")
}
