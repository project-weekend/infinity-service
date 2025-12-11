package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Products struct {
	ID           int              `gorm:"primary_key"`
	CategoryCode string           `json:"categoryCode"`
	ProductIID   string           `gorm:"column:product_iid"`
	ProductSKU   string           `gorm:"column:product_sku"`
	Name         string           `gorm:"column:name"`
	Description  string           `gorm:"column:description"`
	Price        decimal.Decimal  `gorm:"column:price"`
	Quantity     decimal.Decimal  `gorm:"column:quantity"`
	CreatedAt    time.Time        `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time        `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Category     *ProductCategory `gorm:"foreignKey:CategoryCode;references:category_code"`
}

func (a *Products) TableName() string {
	return "products"
}
