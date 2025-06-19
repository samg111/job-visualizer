package buttons

import (
	"job-visualizer/pkg/jobdata"
	"job-visualizer/pkg/mapping"
	"job-visualizer/pkg/shared"
)

func HandleJobRefresh(jobs []shared.JobData) {
	removeActiveFilters()
	jobs = jobdata.GetJobData(jobs)
	mapping.GenerateMap(jobs)
	refreshEntries()
}

func HandleJobFilter(jobs []shared.JobData) {
	jobs = jobdata.GetJobData(jobs)
	mapping.GenerateMap(jobs)
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
