package main

import (
	"job-visualizer/pkg/excel"
	// "job-visualizer/pkg/gui"

	"testing"
)

func TestReadExcel(testing *testing.T) {
	expectedRows := 300
	expectedColumns := 2

	rows := readExcel()
	rowsLen := len(rows)
	columnsLen := len(rows[0])
	if rowsLen < expectedRows || columnsLen < expectedColumns {
		testing.Errorf("reading data from excel file\nrows returned: %d, expected: %d\ncolumns returned: %d, expected: %d",
			rowsLen, expectedRows, expectedColumns, expectedColumns)
	}
}

func readExcel() [][]string {
	file := excel.OpenExcelFile()
	rows := excel.GetAllRows(file)
	return rows
}
