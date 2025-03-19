package pagination

import (
	"net/http"
	"strconv"
)

type Pagination struct {
	Limit int `json:"limit,omitempty"`

	Page       int `json:"page,omitempty"`
	TotalPages int `json:"total_pages"`

	TotalRows int64       `json:"total_rows"`
	Rows      interface{} `json:"rows"`
}

func NewPagination(r *http.Request) *Pagination {
	limitStr := r.URL.Query().Get("limit")
	pageStr := r.URL.Query().Get("page")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	if limit > 50 {
		limit = 50
	}

	return &Pagination{
		Limit: limit,
		Page:  page,
	}
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}
func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}
func (p *Pagination) GetLimit() int {
	if p.Limit <= 0 {
		p.Limit = 10
	}
	if p.Limit > 50 {
		p.Limit = 50
	}
	return p.Limit
}
