package product

import (
	"github.com/AlaraEfe/BasketService/internal/api"
	"github.com/AlaraEfe/BasketService/internal/models"
	"gorm.io/gorm"
)

func responseToProduct(p *api.Product) *models.Product {
	return &models.Product{
		Model:            gorm.Model{ID: uint(p.ID)},
		ProductName:      p.ProductName,
		CategoryID:       p.Category.ID,
		ProductStock:     p.ProductStock,
		ProductPrice:     p.ProductPrice,
		ProductStockCode: p.ProductStockCode,
		ProductType:      p.ProductType,
		ProductColor:     p.ProductColor,
		ProductSize:      p.ProductSize,
		ProductMaterial:  p.ProductMaterial,
		Gender:           p.Gender,
	}
}

func ProductToResponse(p *models.Product) *api.Product {

	return &api.Product{
		ID:               p.ID,
		ProductName:      p.ProductName,
		ProductStock:     p.ProductStock,
		ProductPrice:     p.ProductPrice,
		ProductStockCode: p.ProductStockCode,
		ProductType:      p.ProductType,
		ProductColor:     p.ProductColor,
		ProductSize:      p.ProductSize,
		ProductMaterial:  p.ProductMaterial,
		Gender:           p.Gender,
		Category: &api.Category{
			ID:           p.CategoryID,
			CategoryName: p.Category.CategoryName,
		},
	}
}

func productsToResponse(ps *[]models.Product) []*api.Product {
	products := make([]*api.Product, 0)
	for _, p := range *ps {
		products = append(products, ProductToResponse(&p))
	}
	return products
}
