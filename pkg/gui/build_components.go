package gui

import (
	// "job-visualizer/pkg/gui"
	// "job-visualizer/pkg/gui/job_data"
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
		GetJobData(jobs)
		// openWebpage()
		refreshEntries()
	})
	displayPane := container.NewVBox(getJobsButton, mapButton)
	return displayPane
}

func BuildFilterComponents() (*fyne.Container, *widget.Check) {
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

// func getJobData(jobs []structs.JobData) {
// 	// Filter jobs
// 	jobs = filterJobs(jobs)
// 	// jobs = assignLatLongs(jobs)
// 	// geoplotMap := createGeoplotMap(jobs)
// 	// window.Server = createHttpServer(geoplotMap)
// 	window.JobDataGui = &jobs
// }

// func filterJobs(jobs []structs.JobData) []structs.JobData {
// 	filters := window.Filters
// 	if filters.KeywordEntry != "" || filters.LocationEntry != "" || filters.MinSalaryEntry != "" ||
// 		filters.WorkFromHomeEntry {
// 		var filteredJobs []structs.JobData
// 		for _, job := range jobs {
// 			filteredJobs = filterIndividualJob(job, filteredJobs)
// 		}
// 		return filteredJobs
// 	}
// 	return jobs
// }

// func filterIndividualJob(job structs.JobData, filteredJobs []structs.JobData) []structs.JobData {
// 	filters := window.Filters
// 	filterMatch := true
// 	if filters.KeywordEntry != "" {
// 		//fmt.Printf("keyword entered: %s", filters.KeywordEntry)
// 		filterMatch = FilterKeyword(job, filters.KeywordEntry)
// 	}
// 	if filters.LocationEntry != "" && filterMatch {
// 		//fmt.Printf("location entered: %s", filters.LocationEntry)
// 		filterMatch = FilterLocation(job, filters.LocationEntry)
// 	}
// 	if filters.MinSalaryEntry != "" && filterMatch {
// 		//fmt.Printf("min salary entered: %s", filters.MinSalaryEntry)
// 		filterMatch = FilterMinSalary(job, filters.MinSalaryEntry)
// 	}
// 	if filters.WorkFromHomeEntry && filterMatch {
// 		//fmt.Println("work from home filter applied")
// 		filterMatch = FilterWorkFromHome(job)
// 	}
// 	if filterMatch {
// 		filteredJobs = append(filteredJobs, job)
// 	}
// 	return filteredJobs
// }

// func FilterKeyword(job structs.JobData, filterInput string) bool {
// 	filterMatch := false
// 	filter := strings.ToLower(filterInput)
// 	jobTitle := strings.ToLower(job.JobTitle)
// 	companyName := strings.ToLower(job.CompanyName)
// 	description := strings.ToLower(job.Description)
// 	qualifications := strings.ToLower(job.Qualifications)
// 	if strings.Contains(jobTitle, filter) || strings.Contains(companyName, filter) ||
// 		strings.Contains(description, filter) || strings.Contains(qualifications, filter) {
// 		filterMatch = true
// 	}
// 	return filterMatch
// }

// func FilterLocation(job structs.JobData, filterInput string) bool {
// 	filterMatch := false
// 	jobLocation := strings.ToLower(job.Location)
// 	filter := strings.ToLower(filterInput)
// 	if strings.Contains(jobLocation, filter) {
// 		filterMatch = true
// 	}
// 	return filterMatch
// }

// func FilterMinSalary(job structs.JobData, filter string) bool {
// 	filterMatch := false
// 	salary, _ := strconv.Atoi(job.Salary)
// 	minSalary, err := strconv.Atoi(filter)
// 	checkErrorWarn(err)
// 	if salary > minSalary {
// 		filterMatch = true
// 	}
// 	return filterMatch
// }

// func FilterWorkFromHome(job structs.JobData) bool {
// 	filterMatch := false
// 	if job.WorkFromHome == "Yes" {
// 		filterMatch = true
// 	}
// 	return filterMatch
// }

// func getJobData(jobs []structs.JobData) {
// 	window.JobDataGui = &jobs
// }

// func GetJobData(jobs []structs.JobData) *[]structs.JobData {
// 	// jobs = filterJobs(jobs)
// 	// jobs = assignLatLongs(jobs)
// 	// geoplotMap := createGeoplotMap(jobs)
// 	// window.Server = createHttpServer(geoplotMap)
// 	return &jobs
// }

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
