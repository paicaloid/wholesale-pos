package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `gorm:"foreignKey:CategoryID"`
}

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
