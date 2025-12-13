package entity

import "time"

type UserStatus string

const (
	UserStatus_Active   UserStatus = "active"
	UserStatus_Inactive UserStatus = "inactive"
)

type User struct {
	ID        int       `gorm:"column:id;primaryKey"`
	RoleID    int       `gorm:"column:role_id"`
	UserCode  string    `gorm:"column:user_code"`
	Email     string    `gorm:"column:email"`
	Status    string    `gorm:"column:status"`
	Password  string    `gorm:"column:password"`
	CreatedBy string    `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Role      Roles     `gorm:"foreignKey:role_id;references:id"`
}

func (a *User) TableName() string {
	return "users"
}
