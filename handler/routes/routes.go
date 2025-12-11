package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/infinity/infinity-service/handler"
)

type RouteConfig struct {
	AppEngine              *fiber.App
	UserHandler            *handler.UserHandler
	ProductCategoryHandler *handler.ProductCategoryHandler
	Middleware             fiber.Handler
}

func (r *RouteConfig) Setup() {
	r.RegisterRoutes()
	r.RegisterProtectedRoutes()
}

func (r *RouteConfig) RegisterRoutes() {
	// Public routes (no authentication required)
	r.AppEngine.Post("/api/v1/user/login", r.UserHandler.Login)
	r.AppEngine.Post("/api/v1/product-category/:id", r.ProductCategoryHandler.Get)
	r.AppEngine.Delete("/api/v1/product-category/:id", r.ProductCategoryHandler.Delete)
}

func (r *RouteConfig) RegisterProtectedRoutes() {
	// Create protected group with authentication middleware
	protected := r.AppEngine.Group("/api/v1", r.Middleware)

	// Protected user routes
	protected.Post("/user", r.UserHandler.Register)
	protected.Post("/user/me", r.UserHandler.CurrentUser)

	protected.Post("/product-category", r.ProductCategoryHandler.Create)
	protected.Get("/product-categories", r.ProductCategoryHandler.List)
	//protected.DeleteByID("/product-category/:id", r.ProductCategoryHandler.Get)
}
