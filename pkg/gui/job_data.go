package gui

import (
	// "job-visualizer/pkg/gui"
	"job-visualizer/pkg/structs"
	"strconv"
	"strings"
	// "github.com/skratchdot/open-golang/open"
)

func GetJobData(jobs []structs.JobData) {
	// Filter jobs
	jobs = filterJobs(jobs)
	// jobs = assignLatLongs(jobs)
	// geoplotMap := createGeoplotMap(jobs)
	// window.Server = createHttpServer(geoplotMap)
	Window.JobDataGui = &jobs
}

func filterJobs(jobs []structs.JobData) []structs.JobData {
	filters := Window.Filters
	if filters.KeywordEntry != "" || filters.LocationEntry != "" || filters.MinSalaryEntry != "" ||
		filters.WorkFromHomeEntry {
		var filteredJobs []structs.JobData
		for _, job := range jobs {
			filteredJobs = filterIndividualJob(job, filteredJobs)
		}
		return filteredJobs
	}
	return jobs
}

func filterIndividualJob(job structs.JobData, filteredJobs []structs.JobData) []structs.JobData {
	filters := Window.Filters
	filterMatch := true
	if filters.KeywordEntry != "" {
		//fmt.Printf("keyword entered: %s", filters.KeywordEntry)
		filterMatch = FilterKeyword(job, filters.KeywordEntry)
	}
	if filters.LocationEntry != "" && filterMatch {
		//fmt.Printf("location entered: %s", filters.LocationEntry)
		filterMatch = FilterLocation(job, filters.LocationEntry)
	}
	if filters.MinSalaryEntry != "" && filterMatch {
		//fmt.Printf("min salary entered: %s", filters.MinSalaryEntry)
		filterMatch = FilterMinSalary(job, filters.MinSalaryEntry)
	}
	if filters.WorkFromHomeEntry && filterMatch {
		//fmt.Println("work from home filter applied")
		filterMatch = FilterWorkFromHome(job)
	}
	if filterMatch {
		filteredJobs = append(filteredJobs, job)
	}
	return filteredJobs
}

func FilterKeyword(job structs.JobData, filterInput string) bool {
	filterMatch := false
	filter := strings.ToLower(filterInput)
	jobTitle := strings.ToLower(job.JobTitle)
	companyName := strings.ToLower(job.CompanyName)
	description := strings.ToLower(job.Description)
	qualifications := strings.ToLower(job.Qualifications)
	if strings.Contains(jobTitle, filter) || strings.Contains(companyName, filter) ||
		strings.Contains(description, filter) || strings.Contains(qualifications, filter) {
		filterMatch = true
	}
	return filterMatch
}

func FilterLocation(job structs.JobData, filterInput string) bool {
	filterMatch := false
	jobLocation := strings.ToLower(job.Location)
	filter := strings.ToLower(filterInput)
	if strings.Contains(jobLocation, filter) {
		filterMatch = true
	}
	return filterMatch
}

func FilterMinSalary(job structs.JobData, filter string) bool {
	filterMatch := false
	salary := job.Salary
	minSalary, err := strconv.Atoi(filter)
	checkErrorWarn(err)
	if salary > minSalary {
		filterMatch = true
	}
	return filterMatch
}

func FilterWorkFromHome(job structs.JobData) bool {
	filterMatch := false
	if job.WorkFromHome == "Yes" {
		filterMatch = true
	}
	return filterMatch
}
