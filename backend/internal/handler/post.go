package handler

import (
	"github.com/Zecci9/curly/backend/internal/response"
	"github.com/Zecci9/curly/backend/internal/service"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	Service *service.PostService
}

func NewPostHandler(
	s *service.PostService,
) *PostHandler {

	return &PostHandler{
		Service: s,
	}

}

// 创建文章

func (h *PostHandler) Create(c *gin.Context) {

	type CreateRequest struct {
		Title string `json:"title"`

		Content string `json:"content"`
	}

	var req CreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		response.Error(
			c,
			40000,
			"参数错误",
		)

		return

	}

	userIDValue, _ := c.Get("user_id")

	userID := uint(userIDValue.(float64))

	post, err := h.Service.CreatePost(
		req.Title,
		req.Content,
		userID,
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
		post,
	)

}
func (h *PostHandler) List(c *gin.Context) {

	posts, err := h.Service.GetPosts()

	if err != nil {

		response.Error(
			c,
			50000,
			err.Error(),
		)

		return
	}

	response.Success(
		c,
		posts,
	)

}
