package entity

import (
	"time"
)

type ProductStatus string

const (
	ProductStatus_Active    ProductStatus = "active"
	ProductStatus_Inactive  ProductStatus = "inactive"
	ProductStaus_Outofstock ProductStatus = "outofstock"
)

type Products struct {
	ID           int              `gorm:"primary_key"`
	ProductSKU   string           `gorm:"column:product_sku"`
	CategoryCode string           `json:"categoryCode"`
	Name         string           `gorm:"column:name"`
	Description  string           `gorm:"column:description"`
	Price        int64            `gorm:"column:price"`
	Quantity     int64            `gorm:"column:quantity"`
	Status       string           `gorm:"column:status"`
	CreatedAt    time.Time        `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time        `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Category     *ProductCategory `gorm:"foreignKey:CategoryCode;references:category_code"`
}

func (a *Products) TableName() string {
	return "products"
}
