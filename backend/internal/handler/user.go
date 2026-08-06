package handler

import (
	"net/http"

	"github.com/Zecci9/curly/backend/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(
	s *service.UserService,
) *UserHandler {

	return &UserHandler{
		Service: s,
	}

}

// 注册接口
func (h *UserHandler) Register(c *gin.Context) {

	type RegisterRequest struct {
		Username string `json:"username"`

		Password string `json:"password"`

		Email string `json:"email"`
	}

	var req RegisterRequest

	// 接收JSON

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "参数错误",
			},
		)

		return
	}

	user, err := h.Service.Register(
		req.Username,
		req.Password,
		req.Email,
	)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "注册成功",
			"data": gin.H{
				"id":       user.ID,
				"username": user.Username,
			},
		},
	)

}
