package jobdata

import (
	"job-visualizer/pkg/jobdata/filter"
	"job-visualizer/pkg/shared"
	"strconv"
)

func ProcessRows(rows [][]string, allJobData []shared.JobData) []shared.JobData {
	for _, row := range rows[1:] {
		job := shared.JobData{}
		job.CompanyName = row[0]
		job.DatePosted = row[1]
		// job.JobId = row[2]
		job.Country = row[3]
		job.Location = row[4]
		job.Salary = calcSalary(row)
		job.JobTitle = row[9]
		allJobData = append(allJobData, job)
	}
	return allJobData
}

func GetJobData(jobs []shared.JobData) {
	jobs = filter.FilterJobs(jobs)
	// jobs = assignLatLongs(jobs)
	// geoplotMap := createGeoplotMap(jobs)
	// shared.window.Server = createHttpServer(geoplotMap)
	shared.Window.JobDataGui = &jobs
}

// func filterJobs(jobs []shared.JobData) []shared.JobData {
// 	filters := shared.Window.Filters
// 	if filters.KeywordEntry != "" || filters.LocationEntry != "" || filters.MinSalaryEntry != "" ||
// 		filters.WorkFromHomeEntry {
// 		var filteredJobs []shared.JobData
// 		for _, job := range jobs {
// 			filteredJobs = filterIndividualJob(job, filteredJobs)
// 		}
// 		return filteredJobs
// 	}
// 	return jobs
// }

// func filterIndividualJob(job shared.JobData, filteredJobs []shared.JobData) []shared.JobData {
// 	filters := shared.Window.Filters
// 	filterMatch := true
// 	if filters.KeywordEntry != "" {
// 		//fmt.Printf("keyword entered: %s", filters.KeywordEntry)
// 		filterMatch = filterKeyword(job, filters.KeywordEntry)
// 	}
// 	if filters.LocationEntry != "" && filterMatch {
// 		//fmt.Printf("location entered: %s", filters.LocationEntry)
// 		filterMatch = filterLocation(job, filters.LocationEntry)
// 	}
// 	if filters.MinSalaryEntry != "" && filterMatch {
// 		//fmt.Printf("min salary entered: %s", filters.MinSalaryEntry)
// 		filterMatch = filterMinSalary(job, filters.MinSalaryEntry)
// 	}
// 	if filters.WorkFromHomeEntry && filterMatch {
// 		//fmt.Println("work from home filter applied")
// 		filterMatch = filterWorkFromHome(job)
// 	}
// 	if filterMatch {
// 		filteredJobs = append(filteredJobs, job)
// 	}
// 	return filteredJobs
// }

// func filterKeyword(job shared.JobData, filterInput string) bool {
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

// func filterLocation(job shared.JobData, filterInput string) bool {
// 	filterMatch := false
// 	jobLocation := strings.ToLower(job.Location)
// 	filter := strings.ToLower(filterInput)
// 	if strings.Contains(jobLocation, filter) {
// 		filterMatch = true
// 	}
// 	return filterMatch
// }

// func filterMinSalary(job shared.JobData, filter string) bool {
// 	filterMatch := false
// 	salary := job.Salary
// 	minSalary, err := strconv.Atoi(filter)
// 	shared.CheckErrorWarn(err)
// 	if salary > minSalary {
// 		filterMatch = true
// 	}
// 	return filterMatch
// }

// func filterWorkFromHome(job shared.JobData) bool {
// 	filterMatch := false
// 	if job.WorkFromHome == "Yes" {
// 		filterMatch = true
// 	}
// 	return filterMatch
// }

func calcSalary(row []string) int {
	maxSalaryString := row[6]
	minSalaryString := row[7]
	hourlyOrYearly := row[8]

	maxSalary, err := strconv.ParseFloat(maxSalaryString, 64)
	shared.CheckError(err)
	minSalary, err := strconv.ParseFloat(minSalaryString, 64)
	shared.CheckError(err)
	salary := int((maxSalary + minSalary) / 2)
	if hourlyOrYearly == "hourly" {
		salary = salary * 40 * 50
	}
	return salary

}
