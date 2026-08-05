package router

import (
	"github.com/Zecci9/curly/backend/internal/handler"
	"github.com/Zecci9/curly/backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.Default()
	r.Use(middleware.RequestID())
	r.GET("/health", handler.Health)

	return r

}
