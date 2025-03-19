package entity

import "time"

type Content struct {
	ID          int64      `gorm:"id" json:"id"`
	Title       string     `gorm:"title" json:"title"`
	Excerpt     string     `gorm:"excerpt" json:"excerpt"`
	Description string     `gorm:"description" json:"description"`
	Image       string     `gorm:"image" json:"image"`
	Tags        string     `gorm:"tags" json:"tags"`
	Status      string     `gorm:"status" json:"status"`
	CategoryID  int64      `gorm:"category_id" json:"category_id"`
	CreatedByID int64      `gorm:"created_by_id" json:"created_by_id"`
	User        User       `gorm:"foreignKey:CreatedByID" json:"user"`
	Category    Category   `gorm:"foreignKey:CategoryID" json:"category"`
	CreatedAt   time.Time  `gorm:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"updated_at" json:"updated_at"`
}

//type Content struct {
//	ID          int64      `gorm:"id"`
//	Title       string     `gorm:"title"`
//	Excerpt     string     `gorm:"excerpt"`
//	Description string     `gorm:"description"`
//	Image       string     `gorm:"image"`
//	Tags        string     `gorm:"tags"`
//	Status      string     `gorm:"status"`
//	CategoryID  int64      `gorm:"category_id"`
//	CreatedByID int64      `gorm:"created_by_id"`
//	User        User       `gorm:"foreignKey:CreatedByID"`
//	Category    Category   `gorm:"foreignKey:CategoryID"`
//	CreatedAt   time.Time  `gorm:"created_at"`
//	UpdatedAt   *time.Time `gorm:"updated_at"`
//}

//type ContentEntity struct {
//	ID          int64
//	Title       string
//	Excerpt     string
//	Description string
//	Image       string
//	Tags        []string
//	Status      string
//	CategoryID  int64
//	CreatedByID int64
//	CreatedAt   time.Time
//	Category    CategoryEntity
//	User        UserEntity
//}
