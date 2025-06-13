package main

import (
	"job-visualizer/pkg/database"
	"job-visualizer/pkg/excel"
	"job-visualizer/pkg/gui"
	"job-visualizer/pkg/jobdata"
	"job-visualizer/pkg/structs"

	_ "modernc.org/sqlite"
)

func main() {
	file := excel.OpenExcelFile()
	rows := excel.GetAllRows(file)
	allJobData := jobdata.ProcessRows(rows, []structs.JobData{})

	jobsDatabase := database.CreateDatabase()
	database.SetupDatabase(jobsDatabase)
	database.WriteToDatabase(jobsDatabase, allJobData)

	// for i, job := range allJobData {
	// 	fmt.Printf("Job %d:\n", i+1)
	// 	fmt.Printf("  Location: %s\n", job.Location)
	// 	fmt.Printf("  Job Title: %s\n", job.JobTitle)
	// 	fmt.Printf("  Company Name: %s\n", job.CompanyName)
	// 	fmt.Printf("  Job Id: %s\n", job.JobId)
	// 	fmt.Printf("  Country: %s\n", job.Country)
	// 	fmt.Printf("  Description: %s\n", job.Description)
	// 	fmt.Printf("  Date Posted: %s\n", job.DatePosted)
	// 	fmt.Printf("  Salary: %s\n", job.Salary)
	// 	fmt.Printf("  Work From Home: %s\n", job.WorkFromHome)
	// 	fmt.Printf("  Qualifications: %s\n", job.Qualifications)
	// 	fmt.Printf("  Links: %s\n", job.Links)
	// 	fmt.Println()
	// }

	gui.CreateGui(allJobData)
}
