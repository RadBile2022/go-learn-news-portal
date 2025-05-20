package request

import (
	"context"
	"go-learn-news-portal/internal/core/entity"
	"go-learn-news-portal/library/v1/middleware"
)

type ContentID struct {
	ID int64
}

func (c *ContentID) ToEntity() *entity.Content {
	return &entity.Content{
		ID: c.ID,
	}
}

type ContentForm struct {
	ID          int64
	Title       string `json:"title"`
	Description string `json:"description"`
	Excerpt     string `json:"excerpt"`
	Tags        string `json:"tags"`
	CategoryId  int    `json:"category_id"`
	Status      string `json:"status"`
	Image       string `json:"image"`
}

func (d *ContentForm) ToEntity(ctx context.Context) *entity.Content {
	userID := middleware.GetUserIDFromContext(ctx)
	e := &entity.Content{
		Title:       d.Title,
		Tags:        d.Tags,
		Description: d.Description,
		Status:      d.Status,
		CategoryID:  int64(d.CategoryId),
		Image:       d.Image,
		Excerpt:     d.Excerpt,
		CreatedByID: userID,
	}

	if d.ID != 0 {
		e.ID = d.ID
	}

	return e
}
