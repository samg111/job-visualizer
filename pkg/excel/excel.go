package excel

import (
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

func OpenExcelFile() *excelize.File {
	os.Chdir("../..")
	file, err := excelize.OpenFile("JobData.xlsx")
	checkError(err)
	return file
}

func GetAllRows(file *excelize.File) [][]string {
	rows, err := file.GetRows("Jobs")
	checkError(err)
	return rows
}

// func ProcessRows(rows [][]string, allJobData []structs.JobData) []structs.JobData {
// 	for _, row := range rows[1:] {
// 		job := structs.JobData{}
// 		job.CompanyName = row[0]
// 		job.DatePosted = row[1]
// 		job.Location = row[4]
// 		// job.Salary = CalcSalary(row)
// 		job.JobTitle = row[9]
// 		allJobData = append(allJobData, job)
// 	}
// 	return allJobData
// }

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
