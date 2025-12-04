package handler

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/infinity/infinity-service/internal/model"
	"github.com/infinity/infinity-service/internal/service"
)

type UserHandler struct {
	Logger      *slog.Logger
	Validate    *validator.Validate
	UserService service.IUserService
}

func NewUserHandler(logger *slog.Logger, validate *validator.Validate, userService service.IUserService) *UserHandler {
	return &UserHandler{
		Logger:      logger,
		Validate:    validate,
		UserService: userService,
	}
}

func (i *UserHandler) Register(ctx *fiber.Ctx) error {
	request := &model.CreateUserRequest{}
	if err := ctx.BodyParser(request); err != nil {
		i.Logger.ErrorContext(ctx.UserContext(), "Failed parse body", "err", err)
		return fiber.ErrBadRequest
	}

	if err := i.Validate.Struct(request); err != nil {
		i.Logger.ErrorContext(ctx.UserContext(), "Failed validate request", "err", err)
		return fiber.ErrBadRequest
	}

	response, err := i.UserService.Create(ctx.UserContext(), request)
	if err != nil {
		i.Logger.ErrorContext(ctx.UserContext(), "Failed create user", "err", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (i *UserHandler) Login(ctx *fiber.Ctx) error {
	request := &model.LoginRequest{}
	if err := ctx.BodyParser(request); err != nil {
		i.Logger.ErrorContext(ctx.UserContext(), "Failed parse body", "err", err)
		return fiber.ErrBadRequest
	}

	if err := i.Validate.Struct(request); err != nil {
		i.Logger.ErrorContext(ctx.UserContext(), "Failed validate request", "err", err)
		return fiber.ErrBadRequest
	}

	response, err := i.UserService.Login(ctx.UserContext(), request)
	if err != nil {
		i.Logger.ErrorContext(ctx.UserContext(), "Failed login user", "err", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.LoginResponse]{Data: response})
}

func (i *UserHandler) CurrentUser(ctx *fiber.Ctx) error {
	auth, ok := model.GetAuthFromContext(ctx.UserContext())
	if !ok {
		return fiber.ErrUnauthorized
	}

	response, err := i.UserService.CurrentUser(ctx.UserContext(), auth.ID)
	if err != nil {
		i.Logger.ErrorContext(ctx.UserContext(), "Failed get current user", "err", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}
