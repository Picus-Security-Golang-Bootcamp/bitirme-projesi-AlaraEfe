package models

import (
	"gorm.io/gorm"
)

type BookCategory struct {
	gorm.Model
	ID           uint
	CategoryName string
	Books        []BookProduct `gorm:"foreignKey:CategoryID;references:ID"`
}

type ShoeCategory struct {
	gorm.Model
	ID           uint
	CategoryName string
	Shoes        []ShoeProduct `gorm:"foreignKey:CategoryID;references:ID"`
}

type BagCategory struct {
	gorm.Model
	ID           uint
	CategoryName string
	Bags         []BagProduct `gorm:"foreignKey:CategoryID;references:ID"`
}

type ClothCategory struct {
	gorm.Model
	ID           uint
	CategoryName string
	Clothes      []ClothProduct `gorm:"foreignKey:CategoryID;references:ID"`
}

type JeanCategory struct {
	gorm.Model
	ID           uint
	CategoryName string
	Jeans        []JeanProduct `gorm:"foreignKey:CategoryID;references:ID"`
}

type OutwearCategory struct {
	gorm.Model
	ID           uint
	CategoryName string
	Outwears     []OutwearProduct `gorm:"foreignKey:CategoryID;references:ID"`
}

type SportswearCategory struct {
	gorm.Model
	ID           uint
	CategoryName string
	Sportswear   []SportswearProduct `gorm:"foreignKey:CategoryID;references:ID"`
}
