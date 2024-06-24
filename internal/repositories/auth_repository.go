package repositories

import (
	"go-unit-test/internal/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *models.User) error
	FindUserByUsername(username string) (*models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func (a authRepository) CreateUser(user *models.User) error {
	if err := a.db.Create(user).Error; err != nil {
		return gorm.ErrInvalidData
	}
	return nil
}

func (a authRepository) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := a.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}
