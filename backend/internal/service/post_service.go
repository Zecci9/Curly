package service

import (
	"errors"

	"github.com/Zecci9/curly/backend/internal/model"
	"github.com/Zecci9/curly/backend/internal/repository"
)

type PostService struct {
	PostRepo *repository.PostRepository
}

func NewPostService(
	repo *repository.PostRepository,
) *PostService {

	return &PostService{
		PostRepo: repo,
	}

}

// 创建文章

func (s *PostService) CreatePost(
	title string,
	content string,
	authorID uint,
) (*model.Post, error) {

	if title == "" {

		return nil, errors.New("标题不能为空")

	}

	if content == "" {

		return nil, errors.New("内容不能为空")

	}

	post := &model.Post{

		Title: title,

		Content: content,

		AuthorID: authorID,

		Status: "draft",
	}

	err := s.PostRepo.Create(post)

	if err != nil {

		return nil, err

	}

	return post, nil

}

// 获取文章列表

func (s *PostService) GetPosts() ([]model.Post, error) {

	return s.PostRepo.FindAll()

}

// 获取文章详情

func (s *PostService) GetPostByID(
	id uint,
) (*model.Post, error) {

	return s.PostRepo.FindByID(id)

}
