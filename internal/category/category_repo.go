package category

import (
	"github.com/AlaraEfe/BasketService/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (b *CategoryRepository) Migration() {
	b.db.AutoMigrate(&models.Category{})
}

func (r *CategoryRepository) create(c *models.Category) (*models.Category, error) {
	zap.L().Debug("category.repo.create", zap.Reflect("category", c))

	if err := r.db.Create(c).Error; err != nil {
		zap.L().Error("category.repo.Create failed to create category", zap.Error(err))
		return nil, err
	}

	return c, nil
}

func (r *CategoryRepository) listCategories() (*[]models.Category, int) {
	zap.L().Debug("category.repo.listCategories")

	var cs = &[]models.Category{}
	var count int64

	r.db.Find(&cs).Count(&count)

	return cs, int(count)

}
