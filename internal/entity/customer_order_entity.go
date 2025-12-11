package entity

import (
	"encoding/json"
	"time"
)

type CustomerOrder struct {
	ID         int             `gorm:"column:id;primaryKey"`
	OrderID    int             `gorm:"column:order_id"`
	CustomerID int             `gorm:"column:customer_id"`
	Products   json.RawMessage `gorm:"column:products"`
	CreatedAt  time.Time       `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time       `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (a *CustomerOrder) TableName() string {
	return "customer_orders"
}
