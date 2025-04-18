package main

import (
	"job-visualizer/pkg/excel"
	"job-visualizer/pkg/gui"

	"fmt"
)

func main() {
	gui.CreateGui()

	file := excel.OpenExcelFile()
	rows := excel.GetAllRows(file)
	fmt.Println(rows)

	// allJobData := excel.ProcessRows(rows, []structs.JobData{})
}
