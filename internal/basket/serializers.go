package basket

import (
	"github.com/AlaraEfe/BasketService/internal/api"
	"github.com/AlaraEfe/BasketService/internal/models"
	"gorm.io/gorm"
)

func ResponseToBasketItem(i *api.BasketItem) *models.BasketItem {
	return &models.BasketItem{
		Model:     gorm.Model{ID: uint(i.ID)},
		ProductID: i.Product.ID,
		Quantity:  i.Quantity,
		UserID:    i.User.ID,
	}
}

func BasketItemToResponse(i *models.BasketItem) *api.BasketItem {
	return &api.BasketItem{
		ID:       uint(i.ID),
		Quantity: i.Quantity,
		Product: &api.Product{
			ID: uint(i.ProductID),
		},
		User: &api.User{
			ID: uint(i.UserID),
		},
	}
}

func basketItemsToResponse(bs *[]models.BasketItem) []*api.BasketItem {
	items := make([]*api.BasketItem, 0)
	for _, i := range *bs {
		items = append(items, BasketItemToResponse(&i))
	}
	return items
}
