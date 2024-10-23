package models

// Item represents a product in the store.
type Item struct {
	SKU   string  `json:"sku"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
