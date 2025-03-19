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
