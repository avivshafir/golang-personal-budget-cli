package module1

// Budget stores budget information
type Budget struct {
	// Amount stores the budget amount
	Max float32
	// Currency stores the budget currency
	Items []Item
}

// Item stores item information
type Item struct {
	// Name stores the item name
	Description string
	// Amount stores the item amount
	Price float32
}
