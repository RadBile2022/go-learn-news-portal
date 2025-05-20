package entity

import "time"

type User struct {
	ID        int64      `gorm:"id" json:"id"`
	Name      string     `gorm:"name" json:"name"`
	Email     string     `gorm:"email" json:"email"`
	Password  string     `gorm:"password" json:"password"`
	CreatedAt time.Time  `gorm:"created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"updated_at" json:"updated_at"`
}

func (User) TableName() string {
	return "app_news_users"
}

type Users struct {
	ID                 int64     `gorm:"id" json:"id"`
	Name               string    `gorm:"name" json:"name"`
	Email              string    `gorm:"email" json:"email"`
	Password           string    `gorm:"password" json:"password"`
	QuotaFree          int64     `gorm:"quota_free" json:"quota_free"` // 12 files | 64MB
	QuotaFreeUsed      int64     `gorm:"quota_free_used" json:"quota_free_used"`
	QuotaFreeExpiredAt time.Time `gorm:"quota_free_expired_at" json:"quota_free_expired_at"`

	QuotaPaid          int64     `gorm:"quota_paid" json:"quota_paid"` //
	QuotaPaidUsed      int64     `gorm:"quota_paid_used" json:"quota_paid_used"`
	QuotaPaidExpiredAt time.Time `gorm:"quota_paid_expired_at" json:"quota_paid_expired_at"`

	CreatedAt time.Time  `gorm:"created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"updated_at" json:"updated_at"`
}

func (Users) TableName() string {
	return "app_stg_users"
}
