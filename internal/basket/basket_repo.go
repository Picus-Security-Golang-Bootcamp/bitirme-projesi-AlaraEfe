package basket

import (
	"github.com/AlaraEfe/BasketService/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type basketItemRepository struct {
	db *gorm.DB
}

func NewBasketItemRepository(db *gorm.DB) *basketItemRepository {
	return &basketItemRepository{db: db}
}

func (b *basketItemRepository) Migration() {
	b.db.AutoMigrate(&models.BasketItem{})
}

func (b *basketItemRepository) addToCart(i *models.BasketItem) (*models.BasketItem, error) {
	if err := b.db.Preload("User").Preload("Product").Create(i).Error; err != nil {
		zap.L().Error("basketItem.repo.Create failed to create basketItem", zap.Error(err))
		return nil, err
	}

	return i, nil
}

func (b *basketItemRepository) listCart(id uint) (*[]models.BasketItem, error) {
	zap.L().Debug("basketItem.repo.listCart", zap.Reflect("id", id))

	var items = &[]models.BasketItem{}
	if result := b.db.Where("user_id = ?", id).Find(&items); result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

func (b *basketItemRepository) deleteBasketItemByID(id uint) error {
	zap.L().Debug("basketItem.repo.deleteBasketItemByID", zap.Reflect("id", id))
	item, err := b.GetByID(id)
	if err != nil {
		return nil
	}
	if result := b.db.Delete(&item); result.Error != nil {
		return nil
	}

	return nil
}

func (b *basketItemRepository) GetByID(id uint) (*models.BasketItem, error) {
	zap.L().Debug("basketItem.repo.getByID", zap.Reflect("id", id))

	var item = &models.BasketItem{}
	if result := b.db.First(&item, id); result.Error != nil {
		return nil, result.Error
	}
	return item, nil
}

func (b *basketItemRepository) UpdateBasketItemByID(i *models.BasketItem) (*models.BasketItem, error) {
	zap.L().Debug("basketItem.repo.updateProductByID", zap.Reflect("basketItem", i))

	if result := b.db.Save(&i); result.Error != nil {
		return nil, result.Error
	}
	item, err := b.GetByID(i.ID)
	if err != nil {
		return nil, err
	}

	return item, nil
}
