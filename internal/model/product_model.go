package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type ProductResponse struct {
	ProductIID  string          `json:"product_iid,omitempty"`
	ProductSKU  string          `json:"product_sku,omitempty"`
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Price       decimal.Decimal `json:"price,omitempty"`
	CreatedAt   time.Time       `json:"created_at,omitempty"`
	UpdatedAt   time.Time       `json:"updated_at,omitempty"`
}

type CreateProductRequest struct {
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Price       decimal.Decimal `json:"price,omitempty"`
}
