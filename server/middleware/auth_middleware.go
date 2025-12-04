package middleware

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/infinity/infinity-service/internal/model"
	"github.com/infinity/infinity-service/internal/service/user"
)

func NewAuth(userService *user.UserServiceImpl) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			userService.Logger.WarnContext(ctx.UserContext(), "Missing Authorization header")
			return fiber.ErrUnauthorized
		}
		token := authHeader
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			token = authHeader[7:]
		}

		request := &model.VerifyUserRequest{
			Token: token,
		}

		auth, err := userService.Verify(ctx.UserContext(), request)
		if err != nil {
			userService.Logger.WarnContext(ctx.UserContext(), "Failed verify user", "err", err)
			return fiber.ErrUnauthorized
		}

		userService.Logger.DebugContext(ctx.UserContext(), "User :", "user", auth.ID)

		// Use the AuthContextKey from the model package to ensure consistency
		newCtx := context.WithValue(ctx.UserContext(), model.AuthContextKey, auth)
		ctx.SetUserContext(newCtx)
		return ctx.Next()
	}
}
