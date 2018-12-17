package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func main() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	style := xlsx.NewStyle()

	// fill := *xlsx.NewFill("solid", "00FFFFFF", "00FFFFFF")
	// font := *xlsx.NewFont(20, "Verdana")
	// border := *xlsx.NewBorder("thin", "thin", "thin", "thin")
	font := *xlsx.DefaultFont()
	font.Color = "00FF0000"
	// style.Fill = fill
	style.Font = font
	// style.Border = border

	style.ApplyFill = true
	style.ApplyFont = true
	style.ApplyBorder = true

	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("你你您")
	row = sheet.AddRow()

	cell = row.AddCell()
	cell.Value = "100000"
	cell.SetStyle(style)

	cell = row.AddCell()
	cell.Value = "test"

	err = file.Save("test.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}

}
