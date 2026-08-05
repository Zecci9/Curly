package router

import (
	"github.com/Zecci9/curly/backend/internal/handler"
	"github.com/Zecci9/curly/backend/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.Default()

	// CORS跨域配置
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
	}))

	r.Use(middleware.RequestID())

	r.GET("/health", handler.Health)

	return r
}
