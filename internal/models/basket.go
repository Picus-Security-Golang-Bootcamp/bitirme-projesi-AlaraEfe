package models

import (
	"gorm.io/gorm"
)

type BasketItem struct {
	gorm.Model
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
	UserID    uint `json:"user_id"`
	Product   *Product
}
