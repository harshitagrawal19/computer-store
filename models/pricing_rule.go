package models

// BuyXGetYFreeRule represents a "buy X get Y free" pricing rule.
type BuyXGetYFreeRule struct {
	SKU  string `json:"sku"`
	Buy  int    `json:"buy"`
	Free int    `json:"free"`
}

// BulkDiscountRule represents a bulk discount pricing rule.
type BulkDiscountRule struct {
	SKU             string  `json:"sku"`
	Threshold       int     `json:"threshold"`
	DiscountedPrice float64 `json:"discounted_price"`
}

// PricingRules holds the different types of rules.
type PricingRules struct {
	BuyXGetYFree []BuyXGetYFreeRule `json:"buy_x_get_y_free"`
	BulkDiscount []BulkDiscountRule `json:"bulk_discount"`
}
