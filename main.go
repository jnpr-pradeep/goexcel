/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

// import "goexcel/cmd"
import (
	"goexcel/pkg/bom"
	"goexcel/pkg/utils"
)

// func main() {

// 	cmd.Execute()
// }

func getBOM() bom.BOM {
	return bom.BOM{Name: "sample.xlsx", Sections: getSections()}
}

func getSections() []bom.Section {
	return []bom.Section{{Name: "section1", Items: getLineItem()}}
}

func getLineItem() []bom.LineItem {
	return []bom.LineItem{{Quantity: 10, Item: bom.Item{SKU: "Device1", Price: 10}},
		{Quantity: 5, Item: bom.Item{SKU: "Device2", Price: 20}}}
}

func main() {

	ew := utils.NewExcelWriter()
	// ew.Headers = map[string]string{"A1": "Section Name", "B1": "Name", "C1": "Price", "D1": "Quantity", "F1": "Total Price"}
	ew.WriteToExcel(getBOM())
}
