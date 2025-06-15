package main

import (
	"job-visualizer/pkg/database"
	"job-visualizer/pkg/excel"
	"job-visualizer/pkg/headless"
	"job-visualizer/pkg/jobdata"
	"job-visualizer/pkg/shared"

	_ "modernc.org/sqlite"
)

func main() {
	// headless := flag.Bool("headless", false, "Run in headless mode (no GUI)")
	// flag.Parse()
	isHeadless := headless.CheckCLIArguments()

	file := excel.OpenExcelFile()
	rows := excel.GetAllRows(file)
	allJobData := jobdata.ProcessRows(rows, []shared.JobData{})

	jobsDatabase := database.CreateDatabase()
	database.SetupDatabase(jobsDatabase)
	database.WriteToDatabase(jobsDatabase, allJobData)

	headless.RunGUIorHeadless(isHeadless, allJobData)

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

	// if headless {
	// 	// fmt.Println("headless mode running")
	// 	// fmt.Printf("%-4s | %-22s | %-45s | %-30s | %-10s\n",
	// 	// 	"#", "Location", "Job Title", "Company Name", "Salary")
	// 	// fmt.Println(strings.Repeat("-", 120))

	// 	for i, job := range allJobData {
	// 		if i%100 == 0 {
	// 			fmt.Printf("%-4s | %-22s | %-45s | %-30s | %-10s\n",
	// 				"#", "Location", "Job Title", "Company Name", "Salary")
	// 			fmt.Println(strings.Repeat("-", 120))
	// 		}
	// 		fmt.Printf("%-4d | %-22s | %-45s | %-30s | %-10d\n",
	// 			i+1, job.Location, job.JobTitle, job.CompanyName, job.Salary)
	// 	}
	// } else {
	// 	gui.CreateGui(allJobData)
	// }
	// gui.CreateGui(allJobData)
}
