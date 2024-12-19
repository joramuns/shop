package repository

import (
	"fmt"
	"github.com/joramuns/shop/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, fmt.Errorf("error creating user in repo: %v", err)
	}
	return user, nil
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("error getting user in repo: %v", err)
	}
	return &user, nil
}
