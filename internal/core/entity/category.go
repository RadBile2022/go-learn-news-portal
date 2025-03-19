package entity

import "time"

type Category struct {
	ID          int64      `gorm:"id" json:"id"`
	Title       string     `gorm:"title" json:"title"`
	Slug        string     `gorm:"slug" json:"slug"`
	CreatedByID int64      `gorm:"created_by_id" json:"created_by_id"`
	User        User       `gorm:"foreignKey:CreatedByID" json:"user"`
	CreatedAt   time.Time  `gorm:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"updated_at" json:"updated_at"`
}
