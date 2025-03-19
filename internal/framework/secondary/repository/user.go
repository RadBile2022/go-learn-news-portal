package repository

import (
	"context"
	"errors"
	"go-learn-news-portal/internal/core/entity"
	"go-learn-news-portal/library/v1/handling"
	"gorm.io/gorm"
)

type User interface {
	UpdatePassword(ctx context.Context, e *entity.User) error
	FindUserByID(ctx context.Context, e *entity.User) error
	FindUserByEmail(ctx context.Context, e *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

// FindUserByID implements UserRepository.
func (u *userRepository) FindUserByID(ctx context.Context, e *entity.User) error {
	var err error
	err = u.db.Where("id = ?", e.ID).First(&e).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return handling.NewHttpError404(err, "User not Found")
		}
		return handling.NewHttpError500(err)
	}

	return nil
}

// UpdatePassword implements UserRepository.
func (u *userRepository) UpdatePassword(ctx context.Context, e *entity.User) error {
	var err error
	err = u.db.Model(&e).Where("id = ?", e.ID).Update("password", e.Password).Error
	if err != nil {
		return handling.NewHttpError500(err)
	}

	return nil
}

func (a *userRepository) FindUserByEmail(ctx context.Context, e *entity.User) error {
	var err error

	err = a.db.Where("email = ?", e.Email).First(&e).Error
	if err != nil {
		return handling.NewHttpError500(err)
	}

	return nil
}

func NewUser(db *gorm.DB) User {
	return &userRepository{db: db}
}
