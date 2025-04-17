package main

import (
	"basic-excel-reader/pkg/excel"
	"basic-excel-reader/pkg/gui"

	"fmt"
)

func main() {
	gui.CreateGui()

	file := excel.OpenExcelFile()
	rows := excel.GetAllRows(file)
	fmt.Println(rows)
}
