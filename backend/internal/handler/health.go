package handler

import (
	"github.com/Zecci9/curly/backend/internal/logger"
	"github.com/Zecci9/curly/backend/internal/response"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {

	id, exists := c.Get("request_id")

	if exists {
		logger.Logger.Println("请求ID:", id)
	} else {
		logger.Logger.Println("请求ID不存在")
	}

	logger.Logger.Println("Health check requested")

	response.Success(c, gin.H{
		"name":    "Curly",
		"status":  "running",
		"version": "v0.0.2",
	})
}
