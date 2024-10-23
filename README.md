## Project Structure

```bash
/computer-store
│
├── /config                  # Configuration files
│   ├── products.json       # JSON file containing product details
│   └── pricing_rules.json          # JSON file containing pricing rules
│
├── /checkout                # Checkout system package
│   ├── checkout.go         # Main checkout logic
│
├── /models                  # Data models
│   └── models.go           # Struct definitions for Items and PricingRules
│
├── /utils                   # Utility functions
│   └── utils.go            # Functions to load JSON files
│
├── go.mod                  # Go module file
└── main.go                 # Main application entry point