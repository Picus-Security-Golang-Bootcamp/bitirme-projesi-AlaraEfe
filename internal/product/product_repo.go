package product

import (
	"github.com/AlaraEfe/BasketService/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) Migration() {
	p.db.AutoMigrate(&models.Product{})
}

func (r *ProductRepository) create(p *models.Product) (*models.Product, error) {
	zap.L().Debug("product.repo.create", zap.Reflect("product", p))

	if err := r.db.Preload("Category").Create(p).Error; err != nil {
		zap.L().Error("product.repo.Create failed to create product", zap.Error(err))
		return nil, err
	}

	p1, err := r.getByID(p.ID)
	if err != nil {
		zap.L().Error("product.repo.Create failed to create product", zap.Error(err))
		return nil, err
	}

	return p1, nil
}

func (r *ProductRepository) getAll() (*[]models.Product, int) {
	zap.L().Debug("product.repo.getAll")

	var ps = &[]models.Product{}
	var count int64

	r.db.Preload("Category").Find(&ps).Count(&count)

	return ps, int(count)

}

func (r *ProductRepository) getByID(id uint) (*models.Product, error) {
	zap.L().Debug("product.repo.getByID", zap.Reflect("id", id))

	var product = &models.Product{}
	if result := r.db.Preload("Category").First(&product, id); result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (r *ProductRepository) SearchByProductName(productName string) (*[]models.Product, error) {
	zap.L().Debug("product.repo.SearchByProductName", zap.Reflect("name", productName))
	var products = &[]models.Product{}
	if result := r.db.Preload("Category").Where("Product_Name ILIKE ?", "%"+productName+"%").Find(&products); result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (r *ProductRepository) SearchByProductID(productID uint) (*models.Product, error) {
	zap.L().Debug("product.repo.SearchByProductID", zap.Reflect("id", productID))
	var product = &models.Product{}
	if result := r.db.Preload("Category").Where("ID = ?", productID).First(&product); result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (r *ProductRepository) deleteProductByID(productID uint) error {
	zap.L().Debug("product.repo.delete", zap.Reflect("id", productID))
	product, err := r.getByID(productID)
	if err != nil {
		return nil
	}
	if result := r.db.Delete(&product); result.Error != nil {
		return nil
	}

	return nil
}

func (r *ProductRepository) updateProductByID(p *models.Product) (*models.Product, error) {
	zap.L().Debug("product.repo.updateProductByID", zap.Reflect("product", p))

	if result := r.db.Save(&p); result.Error != nil {
		return nil, result.Error
	}
	product, err := r.getByID(p.ID)
	if err != nil {
		return nil, err
	}

	return product, nil
}
