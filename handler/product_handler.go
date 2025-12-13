package handler

import (
	"log/slog"
	"math"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/infinity/infinity-service/internal/common"
	"github.com/infinity/infinity-service/internal/model"
	"github.com/infinity/infinity-service/internal/service"
	"github.com/infinity/infinity-service/server/middleware"
)

type ProductHandler struct {
	Logger         *slog.Logger
	Validate       *validator.Validate
	ProductService service.IProductService
}

func NewProductHandler(logger *slog.Logger, validate *validator.Validate, productService service.IProductService) *ProductHandler {
	return &ProductHandler{
		Logger:         logger,
		Validate:       validate,
		ProductService: productService,
	}
}

func (h *ProductHandler) Create(ctx *fiber.Ctx) error {
	request := &model.CreateProductRequest{}
	if err := ctx.BodyParser(request); err != nil {
		h.Logger.ErrorContext(ctx.UserContext(), "Failed parse body", "err", err)
		return fiber.ErrBadRequest
	}

	if err := h.Validate.Struct(request); err != nil {
		h.Logger.ErrorContext(ctx.UserContext(), "Failed validate request", "err", err)
		return fiber.ErrBadRequest
	}

	err := h.ProductService.Create(ctx.UserContext(), request)
	if err != nil {
		h.Logger.ErrorContext(ctx.UserContext(), "Failed create product", "err", err)
		return common.AsServiceError(err)
	}

	return ctx.JSON(model.WebResponse[*model.GenericResponse]{Data: &model.GenericResponse{Success: true}})
}

func (h *ProductHandler) List(ctx *fiber.Ctx) error {
	userID := middleware.GetUser(ctx)
	request := &model.SearchProductRequest{
		UserID:     userID.ID,
		ProductSKU: ctx.Query("productSKU", ""),
		Name:       ctx.Query("name", ""),
		Page:       ctx.QueryInt("page", 1),
		Size:       ctx.QueryInt("size", 10),
	}

	if err := h.Validate.Struct(request); err != nil {
		h.Logger.ErrorContext(ctx.UserContext(), "Failed validate request", "err", err)
		return fiber.ErrBadRequest
	}

	response, total, err := h.ProductService.Search(ctx.UserContext(), request)
	if err != nil {
		h.Logger.ErrorContext(ctx.UserContext(), "Failed list products", "err", err)
		return common.AsServiceError(err)
	}

	return ctx.JSON(model.WebResponse[[]model.ProductResponse]{
		Data: response,
		Paging: &model.PageMetadata{
			Page:      request.Page,
			Size:      request.Size,
			TotalItem: total,
			TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
		},
	})
}

func (h *ProductHandler) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	request := &model.GetProductRequest{
		ID: id,
	}

	if err := h.Validate.Struct(request); err != nil {
		h.Logger.ErrorContext(ctx.UserContext(), "Failed validate request", "err", err)
		return fiber.ErrBadRequest
	}

	response, err := h.ProductService.Get(ctx.UserContext(), request)
	if err != nil {
		h.Logger.ErrorContext(ctx.UserContext(), "Failed get product", "err", err)
		return common.AsServiceError(err)
	}

	return ctx.JSON(model.WebResponse[*model.ProductResponse]{Data: response})
}

func (h *ProductHandler) Delete(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	id := ctx.Params("id")
	request := &model.DeleteProductRequest{
		UserID: auth.ID,
		ID:     id,
	}

	if err := h.Validate.Struct(request); err != nil {
		h.Logger.ErrorContext(ctx.UserContext(), "Failed validate request", "err", err)
		return fiber.ErrBadRequest
	}

	err := h.ProductService.Delete(ctx.UserContext(), request)
	if err != nil {
		h.Logger.ErrorContext(ctx.UserContext(), "Failed delete product", "err", err)
		return common.AsServiceError(err)
	}

	return ctx.JSON(model.WebResponse[*model.GenericResponse]{Data: &model.GenericResponse{Success: true}})
}
