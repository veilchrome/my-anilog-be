package repository

import (
	"github.com/veilchrome/myanilog-be/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByUsername(username string) (*domain.User, error)
	FindByUsernameOrEmail(identifier string) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByUsernameOrEmail(identifier string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("username = ? OR email = ?", identifier, identifier).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
