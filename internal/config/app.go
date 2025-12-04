package config

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/infinity/infinity-service/handler"
	"github.com/infinity/infinity-service/handler/routes"
	repository "github.com/infinity/infinity-service/internal/repository/db"
	"github.com/infinity/infinity-service/internal/service/user"
	"github.com/infinity/infinity-service/server/config"
	"github.com/infinity/infinity-service/server/middleware"
	"gorm.io/gorm"
)

type AppBootstrap struct {
	Config    *config.Config
	Logger    *slog.Logger
	DB        *gorm.DB
	Validate  *validator.Validate
	AppEngine *fiber.App
}

func Bootstrap(app *AppBootstrap) {
	// Apply CORS middleware globally
	corsMiddleware := middleware.NewCORS(app.Config)
	app.AppEngine.Use(corsMiddleware)

	// setup repository
	userRepository := repository.NewUserRepository(app.Logger)
	sessionRepository := repository.NewSessionRepository(app.Logger)

	// setup service
	userService := user.NewUserService(app.Config, app.Logger, app.DB, userRepository, sessionRepository)

	// setuo handler
	userHandler := handler.NewUserHandler(app.Logger, app.Validate, userService)

	// setup middleware
	authMiddleware := middleware.NewAuth(userService)

	routeConfig := routes.RouteConfig{
		AppEngine:   app.AppEngine,
		UserHandler: userHandler,
		Middleware:  authMiddleware,
	}

	routeConfig.Setup()
}
