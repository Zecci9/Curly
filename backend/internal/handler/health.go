package handler

import (
	"github.com/Zecci9/curly/backend/internal/response"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {

	response.Success(c, gin.H{
		"name":   "Curly",
		"status": "ok",
	})

}
