package excel

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/sirupsen/logrus"
)

type Excel struct {
	FileName      string
	SheetName     string
	Execfile      *excelize.File
	HeaderStyleID int
	CellStyleID   int
}

func NewExcelMarshal(fileName string, sheetName string) (e Excel) {
	e.FileName = fileName
	e.SheetName = sheetName
	e.Execfile = excelize.NewFile()

	index := e.Execfile.NewSheet(sheetName)

	e.Execfile.SetActiveSheet(index)

	e.HeaderStyleID, _ = e.Execfile.NewStyle(`{"alignment":{"horizontal":"center","Vertical":"center"},"font":{"bold":true}}`)
	e.CellStyleID, _ = e.Execfile.NewStyle(`{"alignment":{"horizontal":"center","Vertical":"center"}}`)
	return
}

func NewExcelUnmarshal(fileName string, sheetName string) (e Excel, err error) {
	e.FileName = fileName
	e.SheetName = sheetName
	e.Execfile, err = excelize.OpenFile(fileName)
	if err != nil {
		logrus.Error(err)
		return e, err
	}

	index := e.Execfile.NewSheet(sheetName)

	e.Execfile.SetActiveSheet(index)
	return
}

func (e *Excel) SetSheetCellValue(axis string, value interface{}) {
	err := e.Execfile.SetCellStyle(e.SheetName, axis, axis, e.CellStyleID)
	if err != nil {
		logrus.Error(err)
	}

	err = e.Execfile.SetCellValue(e.SheetName, axis, value)
	if err != nil {
		logrus.Error(err)
	}
}

func (e *Excel) SetSheetHeaderCellValue(axis string, value interface{}) {
	err := e.Execfile.SetCellStyle(e.SheetName, axis, axis, e.HeaderStyleID)
	if err != nil {
		logrus.Error(err)
	}

	err = e.Execfile.SetCellValue(e.SheetName, axis, value)
	if err != nil {
		logrus.Error(err)
	}
}

func (e *Excel) SetMergeCell(axisBegin, axisEnd string) {
	err := e.Execfile.MergeCell(e.SheetName, axisBegin, axisEnd)
	if err != nil {
		logrus.Error(err)
	}
}

func (e *Excel) GetSheetCellValue(axis string) string {
	value, err := e.Execfile.GetCellValue(e.SheetName, axis)
	if err != nil {
		logrus.Error(err)
	}
	return value
}

func (e *Excel) FileSave() error {
	if err := e.Execfile.SaveAs(e.FileName); err != nil {
		logrus.Error(err.Error())
		return err
	}
	return nil
}
