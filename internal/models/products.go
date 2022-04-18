package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName      string  `json:"product_name"`
	CategoryID       uint    `json:"category_id"`
	ProductStock     int     `json:"product_stock"`
	ProductPrice     float64 `json:"product_price"`
	ProductStockCode string  `json:"product_stock_code"`
	ProductType      string  `json:"product_type,omitempty"`
	ProductColor     string  `json:"product_color,omitempty"`
	ProductSize      string  `json:"product_size,omitempty"`
	ProductMaterial  string  `json:"product_material,omitempty"`
	Gender           string  `json:"gender,omitempty"`
	Category         *Category
}
