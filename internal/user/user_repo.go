package user

import (
	"github.com/AlaraEfe/BasketService/internal/api"
	"github.com/AlaraEfe/BasketService/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (p *UserRepository) Migration() {
	p.db.AutoMigrate(&models.User{})
}

func (u *UserRepository) GetUser(email *string, password *string) (*models.User, error) {
	zap.L().Debug("user.repo.getUser", zap.Reflect("email", email))

	var user = &models.User{}
	if result := u.db.Where("Email = ? AND Password = ?", email, password).First(&user); result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *UserRepository) GetUserByID(id *uint) (*models.User, error) {
	zap.L().Debug("user.repo.getUserByID", zap.Reflect("id", id))

	var user = &models.User{}
	if result := u.db.Preload("BasketItem").Where("ID = ?", id).First(&user); result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *UserRepository) create(user *api.User) (*models.User, error) {
	zap.L().Debug("user.repo.create", zap.Reflect("user", user))

	us := responseToUser(user)

	if err := u.db.Create(us).Error; err != nil {
		zap.L().Error("user.repo.Create failed to create user", zap.Error(err))
		return nil, err
	}

	return us, nil
}
