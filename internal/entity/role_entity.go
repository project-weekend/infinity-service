package entity

import "time"

type Role string

const (
	Role_Admin    Role = "admin"
	Role_Maker    Role = "maker"
	Role_Checker  Role = "checker"
	Role_Approver Role = "approver"
)

type Roles struct {
	ID        int       `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (a *Roles) TableName() string {
	return "roles"
}
