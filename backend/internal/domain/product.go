package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	SKU         string  `json:"sku" gorm:"unique"`
	Description string  `json:"description"`
	CostPrice   float64 `json:"cost_price"`
	SalePrice   float64 `json:"sale_price"`
	Quantity    uint
	CategoryID  uint
}
