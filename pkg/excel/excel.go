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

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
