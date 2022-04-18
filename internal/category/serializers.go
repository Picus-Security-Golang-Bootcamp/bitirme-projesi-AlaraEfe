package category

import (
	"github.com/AlaraEfe/BasketService/internal/api"
	"github.com/AlaraEfe/BasketService/internal/models"
	"github.com/AlaraEfe/BasketService/internal/product"
	//"gorm.io/gorm"
)

/*
func responseToCategory(c *api.Category) *models.Category {

	return &models.Category{
		Model:        gorm.Model{ID: uint(c.ID)},
		CategoryName: c.CategoryName,
	}
}
*/
func CategoryToResponse(c *models.Category) *api.Category {

	products := make([]*api.Product, 0)
	for _, p := range c.Products {
		products = append(products, product.ProductToResponse(&p))
	}

	return &api.Category{
		ID:           uint(c.ID),
		CategoryName: c.CategoryName,
		Products:     products,
	}
}

func categoriesToResponse(cs *[]models.Category) []*api.Category {
	categories := make([]*api.Category, 0)
	for _, c := range *cs {
		categories = append(categories, CategoryToResponse(&c))
	}
	return categories
}
