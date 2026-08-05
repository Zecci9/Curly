package main

import (
	"github.com/Zecci9/curly/backend/internal/config"
	"github.com/Zecci9/curly/backend/internal/logger"
	"github.com/Zecci9/curly/backend/internal/router"
)

func main() {
	cfg := config.Load()

	logger.Init()
	logger.Logger.Println("Curly server starting on port " + cfg.Port)
	r := router.Setup()

	r.Run(":" + cfg.Port)

}
