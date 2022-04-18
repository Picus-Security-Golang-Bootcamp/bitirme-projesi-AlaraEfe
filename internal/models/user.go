package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Role     string        `json:"role"`
	Items    []*BasketItem `json:"items,omitempty"`
}
