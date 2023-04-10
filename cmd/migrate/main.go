package main

import (
	"fmt"
	"os"

	"github.com/caviarman/shop_caviar/internal/logger"
	"github.com/caviarman/shop_caviar/internal/migrate"

	"github.com/caviarman/shop_caviar/internal/migrations"
)

var errEnv = fmt.Errorf("env `DATABASE_URL` is empty")

func main() {
	logger.Init()

	url := os.Getenv("DATABASE_URL")
	if url == "" {
		logger.Critical(errEnv, "os.Getenv")
	}

	err := migrate.Run(url, migrations.FS)
	if err != nil {
		logger.Critical(err, "migrate.Run")
	}
}
