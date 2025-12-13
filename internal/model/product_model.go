package model

import (
	"time"
)

type ProductResponse struct {
	ID          int       `json:"id,omitempty"`
	ProductSKU  string    `json:"productSKU,omitempty"`
	Category    string    `json:"category,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       int64     `json:"price,omitempty"`
	Quantity    int64     `json:"quantity,omitempty"`
	Status      string    `json:"status,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}

type CreateProductRequest struct {
	ProductSKU   string `json:"productSKU" validate:"required,max=100"`
	CategoryCode string `json:"categoryCode" validate:"required,max=100"`
	Name         string `json:"name" validate:"required,max=100"`
	Description  string `json:"description" validate:"required,max=250"`
	Price        int64  `json:"price" validate:"required"`
	Quantity     int64  `json:"quantity" validate:"required"`
}

type SearchProductRequest struct {
	UserID     string `json:"-" validate:"required,max=100"`
	ProductSKU string `json:"productSKU" validate:"max=100"`
	Name       string `json:"name" validate:"max=100"`
	Page       int    `json:"page" validate:"required,min=1"`
	Size       int    `json:"size" validate:"required,min=1"`
}

type GetProductRequest struct {
	ID string `json:"-" validate:"required"`
}

type DeleteProductRequest struct {
	UserID string `json:"-" validate:"required,max=100"`
	ID     string `json:"-" validate:"required"`
}
