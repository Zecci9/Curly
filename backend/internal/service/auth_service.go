package service

import (
	"errors"

	"github.com/Zecci9/curly/backend/internal/auth"
	"github.com/Zecci9/curly/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func NewAuthService(
	repo *repository.UserRepository,
) *AuthService {

	return &AuthService{
		UserRepo: repo,
	}

}

func (s *AuthService) Login(
	username string,
	password string,
) (string, error) {

	// 1. 查询用户

	user, err := s.UserRepo.FindByUsername(username)

	if err != nil {

		return "", errors.New("用户名不存在")

	}

	if user == nil {

		return "", errors.New("用户名不存在")

	}

	// 2. 验证密码

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {

		return "", errors.New("密码错误")

	}

	// 3. 生成Token

	token, err := auth.GenerateToken(
		user.ID,
		user.Username,
	)

	if err != nil {

		return "", err

	}

	return token, nil

}
