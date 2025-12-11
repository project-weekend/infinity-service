package entity

import "time"

type ProductCategory struct {
	ID           int       `gorm:"primary_key;auto_increment"`
	CategoryCode string    `gorm:"column:category_code"`
	Name         string    `gorm:"column:name"`
	Description  string    `gorm:"column:description"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (a *ProductCategory) TableName() string {
	return "product_categories"
}
