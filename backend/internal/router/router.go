package router

import (
	"github.com/Zecci9/curly/backend/internal/database"
	"github.com/Zecci9/curly/backend/internal/handler"
	"github.com/Zecci9/curly/backend/internal/middleware"
	"github.com/Zecci9/curly/backend/internal/repository"
	"github.com/Zecci9/curly/backend/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.Default()
	r.Use(cors.New(cors.Config{

		AllowOrigins: []string{
			"*",
		},

		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},

		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
	}))
	r.Use(middleware.RequestID())

	// 创建用户模块依赖

	userRepo := repository.NewUserRepository(
		database.DB,
	)

	postRepo := repository.NewPostRepository(
		database.DB,
	)

	postService := service.NewPostService(
		postRepo,
	)

	postHandler := handler.NewPostHandler(
		postService,
	)
	authService := service.NewAuthService(
		userRepo,
	)

	authHandler := handler.NewAuthHandler(
		authService,
	)
	userService := service.NewUserService(
		userRepo,
	)

	userHandler := handler.NewUserHandler(
		userService,
	)

	api := r.Group("/api/v1")
	authGroup := api.Group("/admin")

	authGroup.Use(
		middleware.Auth(),
	)

	authGroup.GET(
		"/test",
		func(c *gin.Context) {

			c.JSON(
				200,
				gin.H{
					"message": "通过认证",
				},
			)

		},
	)
	api.GET(
		"/health",
		handler.Health,
	)

	api.POST(
		"/users/register",
		userHandler.Register,
	)

	api.POST(
		"/users/login",
		authHandler.Login,
	)

	api.GET(
		"/users/me",
		middleware.Auth(),
		userHandler.Me,
	)
	api.POST(
		"/posts",
		middleware.Auth(),
		postHandler.Create,
	)
	api.GET(
		"/posts",
		postHandler.List,
	)
	api.GET(
		"/posts/:id",
		postHandler.Detail,
	)
	return r

}
