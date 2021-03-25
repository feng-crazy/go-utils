package excel

import (
	"errors"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/sirupsen/logrus"
)

func JoinCellName(colName string, rowNum int) string {
	cellName, err := excelize.JoinCellName(colName, rowNum)
	if err != nil {
		logrus.Error(err)
	}
	return cellName
}

func CellName(colNum, rowNum int) string {
	cellName, err := excelize.CoordinatesToCellName(colNum, rowNum)
	if err != nil {
		logrus.Error(err)
	}
	return cellName
}

func ColNum(colName string) int {
	colNum, err := excelize.ColumnNameToNumber(colName)
	if err != nil {
		logrus.Error(err)
	}
	return colNum
}

func ColName(num int) string {
	name, err := excelize.ColumnNumberToName(num)
	if err != nil {
		logrus.Error(err)
	}
	return name
}

func FillCellError(err interface{}, rowNum int, colName string) error {
	errStr := fmt.Sprintf("error:%+v, in location :%s, colName:%s", err, JoinCellName(colName, rowNum), colName)
	return errors.New(errStr)
}

func FillAxisError(err interface{}, rowNum int, colNum int) error {
	errStr := fmt.Sprintf("error:%+v, in location :%s, colName:%d", err, CellName(colNum, rowNum), colNum)
	return errors.New(errStr)
}
