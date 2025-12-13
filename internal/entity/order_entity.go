package entity

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
)

const (
	FulfillmentStatus_unfulfilled UserStatus = "unfulfilled"
	FulfillmentStatus_fulfilled   UserStatus = "fulfilled"
)

type Order struct {
	ID          int             `gorm:"column:id;primaryKey"`
	OrderCode   string          `gorm:"column:order_code"`
	Amount      decimal.Decimal `gorm:"column:amount"`
	Products    json.RawMessage `gorm:"column:items"`
	Fulfillment string          `gorm:"column:fulfillment"`
	CreatedAt   time.Time       `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time       `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (a *Order) TableName() string {
	return "orders"
}
