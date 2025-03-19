package repository

import (
	"context"
	"errors"
	"fmt"
	"go-learn-news-portal/internal/core/entity"
	"go-learn-news-portal/library/v1/handling"

	"gorm.io/gorm"
)

type Category interface {
	FindAllCategories(ctx context.Context) ([]*entity.Category, error)
	FindCategoryByID(ctx context.Context, e *entity.Category) error
	CreateCategory(ctx context.Context, e *entity.Category) error
	UpdateCategoryByID(ctx context.Context, e *entity.Category) error
	DeleteCategory(ctx context.Context, e *entity.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

// CreateCategory implements Category.
func (c *categoryRepository) CreateCategory(ctx context.Context, e *entity.Category) error {
	var (
		err       error
		countSlug int64
	)

	err = c.db.Table("categories").Where("slug = ?", e.Slug).Count(&countSlug).Error
	if err != nil {
		return handling.NewHttpError500(err)
	}

	//countSlug = countSlug + 1 // hasilnya kalau gak ada 0
	//slug := fmt.Sprintf("%s-%d", e.Slug, countSlug)
	//e.Slug = slug

	err = c.db.Create(&e).Error
	if err != nil {
		return handling.NewHttpError500(err)
	}

	return nil
}

// DeleteCategory implements Category.
func (c *categoryRepository) DeleteCategory(ctx context.Context, e *entity.Category) error {
	var (
		err   error
		count int64
	)
	err = c.db.Table("contents").Where("category_id = ?", e.ID).Count(&count).Error
	if err != nil {
		return handling.NewHttpError500(err)
	}

	if count > 0 {
		return handling.NewHttpError(err, 409, "Cannot delete a category that has associated contents", "")
	}

	err = c.db.Where("id = ?", e.ID).Delete(&e).Error
	if err != nil {
		return handling.NewHttpError500(err)
	}

	return nil
}

// EditCategoryByID implements Category.
func (c *categoryRepository) UpdateCategoryByID(ctx context.Context, e *entity.Category) error {
	var (
		err       error
		countSlug int64
	)
	err = c.db.Table("categories").Where("slug = ?", e.Slug).Count(&countSlug).Error
	if err != nil {
		return handling.NewHttpError500(err)
	}

	//countSlug = countSlug + 1
	//slug := e.Slug
	//if countSlug == 0 {
	//	slug = fmt.Sprintf("%s-%d", e.Slug, countSlug)
	//}

	//e.Slug = slug

	fmt.Println(e, "apa dah")
	err = c.db.Where("id = ?", e.ID).Updates(&e).Error
	if err != nil {
		return handling.NewHttpError500(err)
	}

	return nil
}

// FindAllCategories implements Category.
func (c *categoryRepository) FindAllCategories(ctx context.Context) (el []*entity.Category, err error) {
	err = c.db.Order("created_at DESC").Preload("User").Find(&el).Error
	if err != nil {
		return nil, handling.NewHttpError500(err)
	}

	return el, nil
}

// FindCategoryByID implements Category.
func (c *categoryRepository) FindCategoryByID(ctx context.Context, e *entity.Category) error {
	var err error
	err = c.db.Where("id = ?", e.ID).Preload("User").First(&e).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return handling.NewHttpError(err, 404, "Category not found", "")
		}
		return handling.NewHttpError500(err)
	}

	return nil
}

func NewCategory(db *gorm.DB) Category {
	return &categoryRepository{
		db: db,
	}
}
