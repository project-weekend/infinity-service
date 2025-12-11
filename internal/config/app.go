package config

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/valkey-io/valkey-go"
	"gorm.io/gorm"

	"github.com/infinity/infinity-service/handler"
	"github.com/infinity/infinity-service/handler/routes"
	"github.com/infinity/infinity-service/internal/repository"
	"github.com/infinity/infinity-service/internal/repository/cache"
	"github.com/infinity/infinity-service/internal/repository/db"
	"github.com/infinity/infinity-service/internal/service/productcategory"
	"github.com/infinity/infinity-service/internal/service/user"
	"github.com/infinity/infinity-service/server/config"
	"github.com/infinity/infinity-service/server/middleware"
)

type AppBootstrap struct {
	Config    *config.Config
	Logger    *slog.Logger
	DB        *gorm.DB
	Cache     valkey.Client
	Validate  *validator.Validate
	AppEngine *fiber.App
}

func Bootstrap(app *AppBootstrap) {
	// Apply CORS middleware globally
	corsMiddleware := middleware.NewCORS(app.Config)
	app.AppEngine.Use(corsMiddleware)

	// setup mysql repository
	userRepository := db.NewUserRepository(app.Logger)
	sessionRepository := db.NewSessionRepository(app.Logger)
	mySqlProductCategoryRepository := db.NewMySqlProductCategoryRepository(app.Logger, app.DB)

	// setup product category repository with optional caching
	var productCategoryRepository repository.ProductCategoryRepository
	if app.Config.Valkey.Enabled {
		valkeyRepo := cache.NewCacheProductCategoryRepository(app.Logger, app.Cache, mySqlProductCategoryRepository)
		productCategoryRepository = &valkeyRepo
	} else {
		productCategoryRepository = mySqlProductCategoryRepository
	}

	// setup service
	userService := user.NewUserService(app.Config, app.Logger, app.DB, userRepository, sessionRepository)
	productCategoryService := productcategory.NewProductCategoryService(app.Config, app.Logger, app.DB, productCategoryRepository)

	// setup handler
	userHandler := handler.NewUserHandler(app.Logger, app.Validate, userService)
	productCategoryHandler := handler.NewProductCategoryHandler(app.Logger, app.Validate, productCategoryService)

	// setup middleware
	authMiddleware := middleware.NewAuth(userService)

	routeConfig := routes.RouteConfig{
		AppEngine:              app.AppEngine,
		UserHandler:            userHandler,
		ProductCategoryHandler: productCategoryHandler,
		Middleware:             authMiddleware,
	}

	routeConfig.Setup()
}
