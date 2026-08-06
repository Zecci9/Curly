package router

import (
	"github.com/Zecci9/curly/backend/internal/database"
	"github.com/Zecci9/curly/backend/internal/handler"
	"github.com/Zecci9/curly/backend/internal/middleware"
	"github.com/Zecci9/curly/backend/internal/repository"
	"github.com/Zecci9/curly/backend/internal/service"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.Default()

	r.Use(middleware.RequestID())

	// 创建用户模块依赖

	userRepo := repository.NewUserRepository(
		database.DB,
	)

	userService := service.NewUserService(
		userRepo,
	)

	userHandler := handler.NewUserHandler(
		userService,
	)

	api := r.Group("/api/v1")

	api.GET(
		"/health",
		handler.Health,
	)

	api.POST(
		"/users/register",
		userHandler.Register,
	)

	return r
}
