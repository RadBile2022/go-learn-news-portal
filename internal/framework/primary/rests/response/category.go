package response

import (
	"github.com/RadBile2022/go-learn-news-portal/internal/core/entity"
)

type Category struct {
	ID            int64  `json:"id"`
	Title         string `json:"title"`
	Slug          string `json:"slug"`
	CreatedByName string `json:"created_by_name"`
}

func FromEntitiesCategory(entities []*entity.Category) (categories []*Category) {
	for _, v := range entities {
		categories = append(categories, FromEntityCategory(v))
	}
	return categories
}

func FromEntityCategory(e *entity.Category) *Category {
	return &Category{
		ID:            e.ID,
		Title:         e.Title,
		Slug:          e.Slug,
		CreatedByName: e.User.Name,
	}
}
