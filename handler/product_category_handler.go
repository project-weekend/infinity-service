package handler

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/infinity/infinity-service/internal/model"
	"github.com/infinity/infinity-service/internal/service"
)

type ProductCategoryHandler struct {
	Logger                 *slog.Logger
	Validate               *validator.Validate
	ProductCategoryService service.IProductCategoryService
}

func NewProductCategoryHandler(logger *slog.Logger, validate *validator.Validate, productCategoryService service.IProductCategoryService) *ProductCategoryHandler {
	return &ProductCategoryHandler{
		Logger:                 logger,
		Validate:               validate,
		ProductCategoryService: productCategoryService,
	}
}

func (i *ProductCategoryHandler) Create(ctx *fiber.Ctx) error {
	request := &model.CreateProductCategoryRequest{}
	if err := ctx.BodyParser(request); err != nil {
		i.Logger.ErrorContext(ctx.UserContext(), "Failed parse body", "err", err)
		return fiber.ErrBadRequest
	}

	if err := i.Validate.Struct(request); err != nil {
		i.Logger.ErrorContext(ctx.UserContext(), "Failed validate request", "err", err)
		return fiber.ErrBadRequest
	}

	response, err := i.ProductCategoryService.Create(ctx.UserContext(), request)
	if err != nil {
		i.Logger.ErrorContext(ctx.UserContext(), "Failed create product category", "err", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.GenericResponse]{Data: response})
}

func (i *ProductCategoryHandler) List(ctx *fiber.Ctx) error {
	responses, err := i.ProductCategoryService.List(ctx.UserContext())
	if err != nil {
		i.Logger.ErrorContext(ctx.UserContext(), "Failed to get product category list", "err", err)
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.ProductCategoryResponse]{Data: responses})
}
