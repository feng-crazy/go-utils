package excel

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestExcel_GetSheetCellValue(t *testing.T) {
	e, err := NewExcelUnmarshal("example_hdf_test.xlsx", "device")
	if err != nil {
		logrus.Error(err)
		return
	}

	rows, err := e.Execfile.GetRows(e.SheetName)
	if err != nil {
		logrus.Error(err)
		return
	}

	fmt.Println(rows)

	fmt.Println(e.GetSheetCellValue("I9"))
	fmt.Println(e.GetSheetCellValue("I6"))
	fmt.Println(e.GetSheetCellValue("O9"))
	fmt.Println(e.GetSheetCellValue("O9"))
}
