package models

import (
	"gorm.io/gorm"
)

type BookProduct struct {
	gorm.Model
	ProductID        int
	ProductName      string
	BookGenre        string
	BookPageCount    int
	ProductStock     int
	ProductPrice     float64
	ProductStockCode string
	BookISBN         string
	BookAuthorName   string
	CategoryID       uint         `gorm:"primaryKey"`
	ProductCategory  BookCategory `gorm:"foreignKey:CategoryID"`
}

type ShoeProduct struct {
	gorm.Model
	ProductID        int
	ProductName      string
	ShoeColor        string
	ProductStock     int
	ProductPrice     float64
	ProductStockCode string
	ShoeType         string
	ShoeSize         string
	ShoeMaterial     string
	Gender           string
	CategoryID       uint         `gorm:"primaryKey"`
	ProductCategory  ShoeCategory `gorm:"foreignKey:CategoryID"`
}

type BagProduct struct {
	gorm.Model
	ProductID        int
	ProductName      string
	BagColor         string
	ProductStock     int
	ProductPrice     float64
	ProductStockCode string
	BagType          string
	BagMaterial      string
	Gender           string
	CategoryID       uint        `gorm:"primaryKey"`
	ProductCategory  BagCategory `gorm:"foreignKey:CategoryID"`
}

type ClothProduct struct {
	gorm.Model
	ProductID        int
	ProductName      string
	ClothColor       string
	ProductStock     int
	ProductPrice     float64
	ProductStockCode string
	ClothSize        string
	ClothType        string
	ClothMaterial    string
	Gender           string
	CategoryID       uint          `gorm:"primaryKey"`
	ProductCategory  ClothCategory `gorm:"foreignKey:CategoryID"`
}

type JeanProduct struct {
	gorm.Model
	ProductID        int
	ProductName      string
	JeanColor        string
	ProductStock     int
	ProductPrice     float64
	ProductStockCode string
	JeanSize         string
	JeanType         string
	Gender           string
	CategoryID       uint         `gorm:"primaryKey"`
	ProductCategory  JeanCategory `gorm:"foreignKey:CategoryID"`
}

type OutwearProduct struct {
	gorm.Model
	ProductID        int
	ProductName      string
	OutwearColor     string
	ProductStock     int
	ProductPrice     float64
	ProductStockCode string
	OutwearSize      string
	OutwearType      string
	OutwearMaterial  string
	Gender           string
	CategoryID       uint            `gorm:"primaryKey"`
	ProductCategory  OutwearCategory `gorm:"foreignKey:CategoryID"`
}

type SportswearProduct struct {
	gorm.Model
	ProductID          int
	ProductName        string
	SportswearColor    string
	ProductStock       int
	ProductPrice       float64
	ProductStockCode   string
	SportswearSize     string
	SportswearType     string
	SportswearMaterial string
	Gender             string
	CategoryID         uint               `gorm:"primaryKey"`
	ProductCategory    SportswearCategory `gorm:"foreignKey:CategoryID"`
}
