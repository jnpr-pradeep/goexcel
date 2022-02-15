package utils

import (
	"fmt"
	"goexcel/pkg/bom"

	excelize "github.com/xuri/excelize/v2"
)

type ExcelWriter struct {
	Headers map[string]string
	Counter int

	File *excelize.File
}

func NewExcelWriter() *ExcelWriter {
	return &ExcelWriter{Counter: 1, Headers: map[string]string{"A1": "Section Name", "B1": "Name", "C1": "Price", "D1": "Quantity", "E1": "Total Price"}}
}

func (ew *ExcelWriter) setFile(f *excelize.File) {
	ew.File = f
}

func (ew *ExcelWriter) setCounter(c int) {
	ew.Counter = c
}

func (ew *ExcelWriter) createItemEntry(sectionName string, itemName string, itemQty int, itemPrice int16, itemTotalPrice int) {
	ew.File.SetCellValue("Sheet1", fmt.Sprintf("A%v", ew.Counter), sectionName)
	ew.File.SetCellValue("Sheet1", fmt.Sprintf("B%v", ew.Counter), itemName)
	ew.File.SetCellValue("Sheet1", fmt.Sprintf("C%v", ew.Counter), itemQty)
	ew.File.SetCellValue("Sheet1", fmt.Sprintf("D%v", ew.Counter), itemPrice)
	ew.File.SetCellValue("Sheet1", fmt.Sprintf("E%v", ew.Counter), itemTotalPrice)
}

func (ew *ExcelWriter) WriteToExcel(b bom.BOM) {
	f := excelize.NewFile()
	ew.setFile(f)
	for k, v := range ew.Headers {
		// fmt.Println(k, v)
		f.SetCellValue("Sheet1", k, v)
	}
	// Loop through the BOM
	for _, section := range b.Sections {
		// fmt.Println(section.Name, section.Items)
		for _, lineItem := range section.Items {
			ew.Counter += 1
			ew.createItemEntry(section.Name, lineItem.Item.SKU, lineItem.Quantity, lineItem.Item.Price, lineItem.Quantity*int(lineItem.Item.Price))
		}
	}
	// Save the xlshet
	if err := f.SaveAs(b.Name); err != nil {
		fmt.Errorf(err.Error())
	}
}
