package excel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/spf13/cast"
)

func convertNumberToAlphabet(i int) string {
	var Letters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	row := i
	result := Letters[row%26]
	row = row / 26
	for row > 0 {
		row = row - 1
		result = Letters[row%26] + result
		row = row / 26
	}
	return result
}

func CreateFileExcel(columnsTitle []string, data [][]interface{}) *excelize.File {
	xlsx := excelize.NewFile()

	for i, value := range columnsTitle {
		xlsx.SetCellValue("Sheet1", convertNumberToAlphabet(i)+"1", value)
	}

	// arrColWidth := []int{}

	for i, arr := range data {
		for ii, value := range arr {
			xlsx.SetCellValue("Sheet1", convertNumberToAlphabet(ii)+cast.ToString(i+2), value)
		}
	}

	return xlsx
}
