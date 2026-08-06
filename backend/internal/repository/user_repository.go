package repository

import (
	"github.com/Zecci9/curly/backend/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {

	return &UserRepository{
		DB: db,
	}

}

// 创建用户
func (r *UserRepository) Create(user *model.User) error {

	return r.DB.Create(user).Error

}

// 根据用户名查用户
func (r *UserRepository) FindByUsername(username string) (*model.User, error) {

	var user model.User

	err := r.DB.
		Where("username = ?", username).
		First(&user).
		Error

	return &user, err
}

func (r *UserRepository) FindByID(id uint) (*model.User, error) {

	var user model.User

	err := r.DB.
		First(&user, id).
		Error

	if err != nil {

		return nil, err

	}

	return &user, nil

}
