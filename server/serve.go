package server

import (
	"fmt"
	"log"

	"github.com/infinity/infinity-service/internal/config"
)

func Serve() {
	appConfig := config.LoadConfig()
	logger := config.NewLogger(appConfig)
	db := config.NewDatabase(appConfig, logger)
	cache := config.NewCache(appConfig, logger)
	validator := config.NewValidator()
	appEngine := config.NewFiber(appConfig)

	config.Bootstrap(&config.AppBootstrap{
		Config:    appConfig,
		Logger:    logger,
		DB:        db,
		Cache:     cache,
		Validate:  validator,
		AppEngine: appEngine,
	})

	addr := fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port)
	logger.Info(fmt.Sprintf("Starting HTTP server on %s", addr))

	if err := appEngine.Listen(addr); err != nil {
		log.Fatal(fmt.Errorf("failed to start http server: %w", err))
	}
}
