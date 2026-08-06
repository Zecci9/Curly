package handler

import (
	"github.com/Zecci9/curly/backend/internal/response"
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
		Email    string `json:"email"`
	}

	var req RegisterRequest

	// 接收JSON
	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			40000,
			"参数错误",
		)

		return
	}

	user, err := h.Service.Register(
		req.Username,
		req.Password,
		req.Email,
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
			"id":       user.ID,
			"username": user.Username,
		},
	)

}
func (h *UserHandler) Me(c *gin.Context) {

	userIDValue, exists := c.Get("user_id")

	if !exists {

		response.Error(
			c,
			40100,
			"未登录",
		)

		return

	}

	userID := uint(userIDValue.(float64))

	user, err := h.Service.GetUserByID(userID)

	if err != nil {

		response.Error(
			c,
			40004,
			"用户不存在",
		)

		return

	}

	response.Success(
		c,
		gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	)

}
