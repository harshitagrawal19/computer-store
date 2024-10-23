package utils

import (
	"computer-store/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

// LoadPricingRules loads pricing rules from a JSON file.
func LoadPricingRules(filename string) (models.PricingRules, error) {
	// Read the Json file
	fileBytes, err := os.Open(filename)
	if err != nil {
		return models.PricingRules{}, errors.New("could not open file: " + err.Error())
	}
	data, err := io.ReadAll(fileBytes)
	if err != nil {
		return models.PricingRules{}, errors.New("could not read file: " + err.Error())
	}

	var rules models.PricingRules
	err = json.Unmarshal(data, &rules)
	if err != nil {
		return models.PricingRules{}, errors.New("could not unmarshal json: " + err.Error())
	}

	return rules, nil
}

// LoadProducts loads product catalog from a JSON file.
func LoadProducts(filename string) (map[string]models.Item, error) {

	fileBytes, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("could not open file: " + err.Error())
	}
	data, err := io.ReadAll(fileBytes)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}

	var products []models.Item
	err = json.Unmarshal(data, &products)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal json: %v", err)
	}

	productMap := make(map[string]models.Item)
	for _, product := range products {
		productMap[product.SKU] = product
	}

	return productMap, nil
}
