package jobdata

import (
	"job-visualizer/pkg/jobdata/filter"
	"job-visualizer/pkg/shared"
	"strconv"
)

func ProcessRows(rows [][]string, allJobData []shared.JobData) []shared.JobData {
	for _, row := range rows[1:] {
		if len(row) < 10 {
			continue
		}
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

func GetJobData(jobs []shared.JobData) []shared.JobData {
	jobs = filter.FilterJobs(jobs)
	// jobs = assignLatLongs(jobs)
	// geoplotMap := createGeoplotMap(jobs)
	// shared.window.Server = createHttpServer(geoplotMap)
	shared.Window.JobDataGui = &jobs
	return jobs
}

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
