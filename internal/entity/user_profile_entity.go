package entity

import "time"

type UserProfile struct {
	ID                int        `gorm:"column:id;primaryKey"`
	UserID            int        `gorm:"column:user_id"`
	ProfileIID        string     `gorm:"column:profile_iid"`
	DisplayName       *string    `gorm:"column:display_name"`
	PronunciationName *string    `gorm:"column:pronunciation_name"`
	Title             *string    `gorm:"column:title"`
	PhoneNumber       *string    `gorm:"column:phone_number"`
	MobilePhone       *string    `gorm:"column:mobile_phone"`
	City              *string    `gorm:"column:city"`
	Country           *string    `gorm:"column:country"`
	EmploymentType    *string    `gorm:"column:employment_type"`
	DateOfBirth       *time.Time `gorm:"column:date_of_birth"`
	AvatarURL         *string    `gorm:"column:avatar_url"`
	Bio               *string    `gorm:"column:bio"`
	CreatedAt         time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt         time.Time  `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	User              User       `gorm:"foreignKey:user_id;references:id"`
}

func (p *UserProfile) TableName() string {
	return "user_profiles"
}
