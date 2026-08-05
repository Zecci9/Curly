package main

import (
	"github.com/Zecci9/curly/backend/internal/config"
	"github.com/Zecci9/curly/backend/internal/router"
)

func main() {

	cfg := config.Load()

	r := router.Setup()

	r.Run(":" + cfg.Port)

}
