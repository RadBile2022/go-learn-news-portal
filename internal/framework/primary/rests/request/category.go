package request

import (
	"context"
	"go-learn-news-portal/internal/core/entity"
	"go-learn-news-portal/library/v1/convert"
	"go-learn-news-portal/library/v1/middleware"
)

type CategoryID struct {
	ID int64
}

func (d *CategoryID) ToEntity() *entity.Category {
	return &entity.Category{
		ID: d.ID,
	}
}

type CategoryForm struct {
	ID    int64
	Title string `json:"title" validate:"required"`
}

func (d *CategoryForm) ToEntity(ctx context.Context) *entity.Category {
	slug := convert.GenerateSlug(d.Title)
	userID := middleware.GetUserIDFromContext(ctx)

	e := &entity.Category{
		Title:       d.Title,
		Slug:        slug,
		CreatedByID: userID,
	}

	if d.ID != 0 {
		e.ID = d.ID
	}

	return e
}
