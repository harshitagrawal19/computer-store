package main

import (
	"computer-store/checkout"
	"computer-store/utils"
	"fmt"
)

func main() {
	// Load pricing rules from JSON.
	rules, err := utils.LoadPricingRules("config/pricing_rules.json")
	if err != nil {
		fmt.Println("Error loading pricing rules:", err)
		return
	}

	// Load product catalog from JSON.
	products, err := utils.LoadProducts("config/products.json")
	if err != nil {
		fmt.Println("Error loading products:", err)
		return
	}

	// Create a new checkout.
	co := checkout.NewCheckout(rules, products)

	// Example scenario 1: "atv, atv, atv, vga"
	co.Scan(products["atv"])
	co.Scan(products["atv"])
	co.Scan(products["atv"])
	co.Scan(products["vga"])
	fmt.Printf("Total: $%.2f\n", co.Total()) // Expected: $249.00 (Buy 2, get 1 free for ATV, VGA is $30)

	// Example scenario 2: "atv, ipd, ipd, atv, ipd, ipd, ipd"
	co = checkout.NewCheckout(rules, products)
	co.Scan(products["atv"])
	co.Scan(products["ipd"])
	co.Scan(products["ipd"])
	co.Scan(products["atv"])
	co.Scan(products["ipd"])
	co.Scan(products["ipd"])
	co.Scan(products["ipd"])
	fmt.Printf("Total: $%.2f\n", co.Total()) // Expected: $2718.95 (Bulk discount for 5 iPads)
}
