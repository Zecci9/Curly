package main

import (
	"fmt"

	"github.com/Zecci9/curly/backend/internal/config"
	"github.com/Zecci9/curly/backend/internal/database"
	"github.com/Zecci9/curly/backend/internal/logger"
	"github.com/Zecci9/curly/backend/internal/model"
	"github.com/Zecci9/curly/backend/internal/router"
)

func main() {

	fmt.Println("进入main")

	cfg := config.Load()

	fmt.Println("准备连接数据库")

	database.Connect(cfg)

	database.DB.AutoMigrate(
		&model.User{},
		&model.Post{},
	)

	fmt.Println("数据库函数执行结束")

	logger.Init()

	logger.Logger.Println(
		"Curly server starting on port " + cfg.Port,
	)

	r := router.Setup()

	// HTTPS启动
	err := r.RunTLS(
		":8080",
		"server.crt",
		"server.key",
	)

	if err != nil {
		panic(err)
	}

}
