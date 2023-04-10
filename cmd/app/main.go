package main

import (
	"github.com/caviarman/shop_caviar/internal/app"
	"github.com/caviarman/shop_caviar/internal/config"
	"github.com/caviarman/shop_caviar/internal/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	logger.Init()

	var cfg config.Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		logger.Critical(err, "cleanenv.ReadEnv")
	}

	err = app.Run(&cfg)
	if err != nil {
		logger.Critical(err, "app.Run")
	}
}
