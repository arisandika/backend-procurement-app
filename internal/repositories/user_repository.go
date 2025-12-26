package repositories

import (
	"procurement-app/internal/models"

	"gorm.io/gorm"

)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

// FIND BY USERNAME
func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CREATE USER
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}
