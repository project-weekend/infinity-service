package entity

import "time"

type UserSession struct {
	ID          int       `gorm:"column:id;primaryKey"`
	UserID      int       `gorm:"column:user_id"`
	SessionCode string    `gorm:"column:session_code"`
	Token       string    `gorm:"column:token"`
	ExpiresAt   time.Time `gorm:"column:expires_at"`
	IPAddress   string    `gorm:"column:ip_address"`
	UserAgent   string    `gorm:"column:user_agent"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (s *UserSession) TableName() string {
	return "user_sessions"
}

func (s *UserSession) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}
