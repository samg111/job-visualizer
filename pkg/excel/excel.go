package excel

import (
	"job-visualizer/pkg/shared"

	"github.com/xuri/excelize/v2"
)

func OpenExcelFile() *excelize.File {
	file, err := excelize.OpenFile("JobData.xlsx")
	shared.CheckError(err)
	return file
}

func GetAllRows(file *excelize.File) [][]string {
	rows, err := file.GetRows("Jobs")
	shared.CheckError(err)
	return rows
}
