package checkout

import (
	"computer-store/models"
)

// Checkout represents a checkout system.
type Checkout struct {
	pricingRules models.PricingRules
	items        map[string]int // Keeps track of item quantities by SKU.
	products     map[string]models.Item
}

// NewCheckout creates a new checkout with given pricing rules.
func NewCheckout(pricingRules models.PricingRules, products map[string]models.Item) *Checkout {
	return &Checkout{
		pricingRules: pricingRules,
		items:        make(map[string]int),
		products:     products,
	}
}

// Scan adds an item to the checkout.
func (co *Checkout) Scan(item models.Item) {
	co.items[item.SKU]++
}

// Total calculates the total price with the applied pricing rules.
func (co *Checkout) Total() float64 {
	total := 0.0

	// Apply "Buy X Get Y Free" rules.
	for _, rule := range co.pricingRules.BuyXGetYFree {
		if quantity, exists := co.items[rule.SKU]; exists {
			// Calculate the number of free items.
			sets := quantity / (rule.Buy + rule.Free)
			freeItems := sets * rule.Free
			chargeableQuantity := quantity - freeItems

			// Add the chargeable quantity to the total.
			total += float64(chargeableQuantity) * co.products[rule.SKU].Price

			// Remove the SKU from further calculation.
			delete(co.items, rule.SKU)
		}
	}

	// Apply bulk discount rules.
	for _, rule := range co.pricingRules.BulkDiscount {
		if quantity, exists := co.items[rule.SKU]; exists {
			if quantity > rule.Threshold {
				// Apply the discounted price.
				total += float64(quantity) * rule.DiscountedPrice
			} else {
				// Apply the regular price.
				total += float64(quantity) * co.products[rule.SKU].Price
			}

			// Remove the SKU from further calculation.
			delete(co.items, rule.SKU)
		}
	}

	// Calculate the remaining items that don't have special rules.
	for sku, quantity := range co.items {
		total += float64(quantity) * co.products[sku].Price
	}

	return total
}
