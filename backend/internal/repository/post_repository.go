package repository

import (
	"github.com/Zecci9/curly/backend/internal/model"

	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(
	db *gorm.DB,
) *PostRepository {

	return &PostRepository{
		DB: db,
	}

}

// 创建文章

func (r *PostRepository) Create(
	post *model.Post,
) error {

	return r.DB.Create(post).Error

}

// 查询文章列表

func (r *PostRepository) FindAll() ([]model.Post, error) {

	var posts []model.Post

	err := r.DB.
		Order("created_at desc").
		Find(&posts).
		Error

	return posts, err

}

// 根据ID查询文章

func (r *PostRepository) FindByID(
	id uint,
) (*model.Post, error) {

	var post model.Post

	err := r.DB.
		First(&post, id).
		Error

	if err != nil {

		return nil, err

	}

	return &post, nil

}
