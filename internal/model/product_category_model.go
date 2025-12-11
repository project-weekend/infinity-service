package model

import "time"

type ProductCategoryResponse struct {
	ID           int       `json:"id,omitempty"`
	CategoryCode string    `json:"categoryCode,omitempty"`
	Name         string    `json:"name,omitempty"`
	Description  string    `json:"description,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty"`
}

type CreateProductCategoryRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
