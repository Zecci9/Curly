package handler

import (
	"github.com/Zecci9/curly/backend/internal/response"
	"github.com/Zecci9/curly/backend/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(
	s *service.AuthService,
) *AuthHandler {

	return &AuthHandler{
		Service: s,
	}

}

// 登录接口
func (h *AuthHandler) Login(c *gin.Context) {

	type LoginRequest struct {
		Username string `json:"username"`

		Password string `json:"password"`
	}

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			40000,
			"参数错误",
		)

		return

	}

	token, err := h.Service.Login(
		req.Username,
		req.Password,
	)

	if err != nil {

		response.Error(
			c,
			40001,
			err.Error(),
		)

		return

	}

	response.Success(
		c,
		gin.H{
			"token": token,
		},
	)

}
