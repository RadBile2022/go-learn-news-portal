package entity

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;default:gen_random_uuid();"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	memo map[string]any `json:"-" gorm:"-"`
}

func (b *BaseModel) HasMemo(key string) bool {
	if b.memo == nil {
		return false
	}
	_, ok := b.memo[key]
	return ok
}

func (b *BaseModel) GetMemo(key string) any {
	if b.memo == nil {
		return nil
	}
	return b.memo[key]
}

func (b *BaseModel) SetMemo(key string, value any) {
	if b.memo == nil {
		b.memo = make(map[string]any)
	}
	b.memo[key] = value
}

type SoftDeleteModel struct {
	BaseModel

	DeletedBy *uint          `json:"deleted_by"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
