package model

import "time"

type UserResponse struct {
	ID        int       `json:"id,omitempty"`
	RoleID    int       `json:"role_id,omitempty"`
	UserCode  string    `json:"user_code,omitempty"`
	Email     string    `json:"email,omitempty"`
	Status    string    `json:"status,omitempty"`
	CreatedBy string    `json:"created_by,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,max=100"`
	RoleID   int    `json:"role_id" validate:"required,max=10"`
	Password string `json:"password" validate:"required,max=100"`
}

type VerifyUserRequest struct {
	Token string `validate:"required,max=100"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type LoginResponse struct {
	UserID int    `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
	Email  string `json:"email,omitempty"`
}

type LogoutRequest struct {
	Token string `json:"token" validate:"required"`
}
