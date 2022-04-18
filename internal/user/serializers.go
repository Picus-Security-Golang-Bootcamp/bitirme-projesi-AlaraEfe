package user

import (
	"github.com/AlaraEfe/BasketService/internal/api"
	basket "github.com/AlaraEfe/BasketService/internal/basket"
	"github.com/AlaraEfe/BasketService/internal/models"
	"gorm.io/gorm"
)

func responseToUser(u *api.User) *models.User {
	items := make([]*models.BasketItem, 0)
	for _, i := range u.Items {
		items = append(items, basket.ResponseToBasketItem(i))
	}
	return &models.User{
		Model:    gorm.Model{ID: uint(u.ID)},
		Email:    u.Email,
		Password: u.Password,
		Role:     u.Role,
		Items:    items,
	}
}

func UserToResponse(u *models.User) *api.User {

	items := make([]*api.BasketItem, 0)
	for _, i := range u.Items {
		items = append(items, basket.BasketItemToResponse(i))
	}
	return &api.User{
		ID:       uint(u.ID),
		Email:    u.Email,
		Password: u.Password,
		Role:     u.Role,
		Items:    items,
	}
}
