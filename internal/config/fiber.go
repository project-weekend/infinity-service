package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/infinity/infinity-service/internal/common"
	"github.com/infinity/infinity-service/server/config"
)

func NewFiber(config *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      config.Name,
		ErrorHandler: NewErrorHandler(),
		Prefork:      config.Prefork,
	})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		if serviceErr := common.AsServiceError(err); serviceErr != nil {
			return ctx.Status(serviceErr.HTTPStatus).JSON(serviceErr)
		}

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
