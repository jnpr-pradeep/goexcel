package bom

type BOM struct {
	Name     string
	Sections []Section
}

// Section represents one section, or subset, of line items in a bill of
// materials.
type Section struct {
	Name  string
	Type  string
	Items []LineItem
}

// LineItem represents one line item in a bill of materials.
type LineItem struct {
	Quantity int
	Item     Item
}

type Item struct {
	SKU   string
	Price int16
}

// NewLineItem creates a new LineItem for a specified quantity of Devices.
func NewLineItem(quantity int, sku string, price int16) *LineItem {
	return &LineItem{
		Quantity: quantity,
		Item:     Item{SKU: sku, Price: price},
	}
}
