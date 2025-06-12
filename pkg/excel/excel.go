package excel

import (
	"log"
	"os"
	"strconv"

	"job-visualizer/pkg/structs"

	"github.com/xuri/excelize/v2"
)

func OpenExcelFile() *excelize.File {
	err := os.Chdir("../..")
	checkError(err)
	file, err := excelize.OpenFile("JobData.xlsx")
	checkError(err)
	return file
}

func GetAllRows(file *excelize.File) [][]string {
	rows, err := file.GetRows("Jobs")
	checkError(err)
	return rows
}

func ProcessRows(rows [][]string, allJobData []structs.JobData) []structs.JobData {
	for _, row := range rows[1:] {
		job := structs.JobData{}
		job.CompanyName = row[0]
		job.DatePosted = row[1]
		job.JobId = row[2]
		job.Country = row[3]
		job.Location = row[4]
		job.Salary = calcSalary(row)
		job.JobTitle = row[9]
		allJobData = append(allJobData, job)
	}
	return allJobData
}

func calcSalary(row []string) string {
	maxSalary := row[6]
	minSalary := row[7]
	hourlyOrYearly := row[8]

	maxSalaryFloat, err := strconv.ParseFloat(maxSalary, 64)
	checkError(err)
	maxSalaryInt := int(maxSalaryFloat)
	minSalaryFloat, err := strconv.ParseFloat(minSalary, 64)
	checkError(err)
	minSalaryInt := int(minSalaryFloat)
	salaryInt := (maxSalaryInt + minSalaryInt) / 2
	if hourlyOrYearly == "hourly" {
		salaryInt = salaryInt * 40 * 50
	}
	salary := strconv.Itoa(salaryInt)
	return salary

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
