package response

import (
	"go-learn-news-portal/internal/core/entity"
	"strings"
)

type Content struct {
	ID           int64    `json:"id"`
	Title        string   `json:"title"`
	Excerpt      string   `json:"excerpt"`
	Description  string   `json:"description,omitempty"`
	Image        string   `json:"image"`
	Tags         []string `json:"tags,omitempty"`
	Status       string   `json:"status"`
	CategoryID   int64    `json:"category_id,omitempty"`
	CreatedByID  int64    `json:"created_by_id,omitempty"`
	CreatedAt    string   `json:"created_at"`
	CategoryName string   `json:"category_name"`
	Author       string   `json:"author"`
}

func FromEntitiesContent(entities []*entity.Content) (contents []*Content) {
	for _, v := range entities {
		contents = append(contents, FromEntityContent(v))
	}
	return contents
}

func FromEntityContent(e *entity.Content) *Content {
	tags := strings.Split(e.Tags, ",")
	return &Content{
		ID:           e.ID,
		Title:        e.Title,
		Status:       e.Status,
		Excerpt:      e.Excerpt,
		Description:  e.Description,
		Image:        e.Image,
		Tags:         tags,
		CategoryID:   e.Category.ID,
		CreatedByID:  e.User.ID,
		Author:       e.User.Name,
		CreatedAt:    e.CreatedAt.Format("2006-01-02 15:04:05"),
		CategoryName: e.Category.Title,
	}
}
