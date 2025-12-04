package model

import "context"

type contextKey string

const AuthContextKey contextKey = "auth"

type Auth struct {
	ID string
}

// GetAuthFromContext retrieves the authenticated user from context
func GetAuthFromContext(ctx context.Context) (*Auth, bool) {
	auth, ok := ctx.Value(AuthContextKey).(*Auth)
	return auth, ok
}
