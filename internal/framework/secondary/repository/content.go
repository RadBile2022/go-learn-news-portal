package repository

import (
	"context"
	"fmt"
	"github.com/RadBile2022/go-learn-news-portal/internal/core/entity"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/handling"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/pagination"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
)

type Content interface {
	FindAllContents(ctx context.Context, q *pagination.Queries, pg *pagination.Pagination) ([]*entity.Content, error)
	FindContentByID(ctx context.Context, e *entity.Content) error
	CreateContent(ctx context.Context, e *entity.Content) error
	UpdateContent(ctx context.Context, e *entity.Content) error
	DeleteContent(ctx context.Context, e *entity.Content) error
}

func NewContent(db *gorm.DB) Content {
	return &contentRepository{db: db}
}

type contentRepository struct {
	db *gorm.DB
}

func (r *contentRepository) FindAllContents(ctx context.Context, q *pagination.Queries, pg *pagination.Pagination) ([]*entity.Content, error) {
	var err error
	var modelContents []*entity.Content
	var totalRows int64

	order := fmt.Sprintf("%s %s", q.OrderBy, q.OrderType)
	offset := (pg.Page - 1) * pg.Limit
	status := ""
	if q.Status != "" {
		status = q.Status
	}

	sqlMain := r.db.Preload(clause.Associations).
		Where("title ilike ? OR excerpt ilike ? OR description ilike ?", "%"+q.Search+"%", "%"+q.Search+"%", "%"+q.Search+"%").
		Where("status LIKE ?", "%"+status+"%")

	if q.CategoryID > 0 {
		sqlMain = sqlMain.Where("category_id =?", q.CategoryID)
	}

	err = sqlMain.Model(&modelContents).Count(&totalRows).Error
	if err != nil {
		return nil, handling.NewHttpError500(err)
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(pg.Limit)))

	err = sqlMain.
		Order(order).
		Limit(pg.Limit).
		Offset(offset).
		Find(&modelContents).Error
	if err != nil {
		return nil, handling.NewHttpError500(err)
	}

	//pg.Limit = pg.Limit

	pg.TotalRows = totalRows
	pg.TotalPages = totalPages
	pg.Rows = modelContents

	return modelContents, nil
}

func (r *contentRepository) FindContentByID(ctx context.Context, e *entity.Content) error {
	err := r.db.Where("id = ?", e.ID).Preload(clause.Associations).First(&e).Error
	if err != nil {
		return handling.NewHttpError500(err)
	}

	return nil
}

func (r *contentRepository) CreateContent(ctx context.Context, e *entity.Content) error {
	err := r.db.Create(&e).Error
	if err != nil {
		return handling.NewHttpError500(err)
	}

	return nil
}

func (r *contentRepository) UpdateContent(ctx context.Context, e *entity.Content) error {
	err := r.db.Where("id = ?", e.ID).Updates(&e).Error
	if err != nil {
		return handling.NewHttpError500(err)
	}

	return nil
}

func (r *contentRepository) DeleteContent(ctx context.Context, e *entity.Content) error {
	err := r.db.Where("id = ?", e.ID).Delete(&e).Error
	if err != nil {
		return handling.NewHttpError500(err)
	}

	return nil
}
