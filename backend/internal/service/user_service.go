package service

import (
	"errors"

	"github.com/Zecci9/curly/backend/internal/model"
	"github.com/Zecci9/curly/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(
	repo *repository.UserRepository,
) *UserService {

	return &UserService{
		UserRepo: repo,
	}
}

// 注册用户
func (s *UserService) Register(
	username string,
	password string,
	email string,
) (*model.User, error) {

	// 1.检查用户名是否存在

	oldUser, err := s.UserRepo.FindByUsername(username)

	if err == nil && oldUser.ID != 0 {

		return nil, errors.New("用户名已存在")

	}

	// 2.密码加密

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {

		return nil, err

	}

	// 3.创建用户

	user := &model.User{

		Username: username,

		Password: string(hash),

		Email: email,
	}

	err = s.UserRepo.Create(user)

	if err != nil {

		return nil, err

	}

	return user, nil

}
