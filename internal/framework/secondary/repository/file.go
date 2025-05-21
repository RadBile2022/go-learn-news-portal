package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/RadBile2022/go-learn-news-portal/internal/core/entity"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/handling"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type File interface {
	BaseRepo[File]

	SumSizeByUploaderID(ctx context.Context, uID uint) (int64, error)
	FindByID(ctx context.Context, e *entity.File) error
	FindByName(ctx context.Context, e *entity.File) error
	FindByNames(ctx context.Context, names []string) ([]*entity.File, error)
	Create(ctx context.Context, e ...*entity.File) error
	Delete(ctx context.Context, e *entity.File) error
}

func NewFile(db *gorm.DB) File {
	fac := func(base *baseRepo[File]) File {
		return &file{
			baseRepo: base,
		}
	}
	return newRepository(fac, db)
}

type file struct {
	*baseRepo[File]
}

func (r *file) SumSizeByUploaderID(ctx context.Context, uID uint) (int64, error) {
	var totalSize int64
	if uID == 0 {
		return 0, nil
	}

	err := r.db.WithContext(ctx).Raw("SELECT COALESCE(SUM(size), 0) AS total_usage FROM app_news_files WHERE uploader_id = ? ", uID).Scan(&totalSize).Error
	if err != nil {
		return 0, err
	}

	return totalSize, nil
}

func (r *file) FindByID(ctx context.Context, e *entity.File) error {
	if e.ID == uuid.Nil {
		return handling.NewHttpError404(nil, "file not found")
	}
	if err := r.db.WithContext(ctx).First(e).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return handling.NewHttpError404(err, "file not found")
		}
		return handling.NewHttpError500(err)
	}
	return nil
}

func (r *file) FindByName(ctx context.Context, e *entity.File) error {
	if e.Name == "" {
		return handling.NewHttpError404(nil, "file not found")
	}
	if err := r.db.WithContext(ctx).Where("name = ?", e.Name).First(e).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return handling.NewHttpError404(err, fmt.Sprintf("file with name %s is not found", e.Name))
		}
		return handling.NewHttpError500(err)
	}
	return nil
}

func (r *file) FindByNames(ctx context.Context, names []string) ([]*entity.File, error) {
	if len(names) == 0 {
		return nil, handling.NewHttpError404(nil, "file not found")
	}
	var files []*entity.File
	if err := r.db.WithContext(ctx).Where("name in ?", names).Find(&files).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, handling.NewHttpError404(err, "file not found")
		}
		return nil, handling.NewHttpError500(err)
	}
	return files, nil
}

func (r *file) Create(ctx context.Context, e ...*entity.File) error {
	if err := r.db.WithContext(ctx).Create(e).Error; err != nil {
		return handling.NewHttpError500(err)
	}
	return nil
}

func (r *file) Delete(ctx context.Context, e *entity.File) error {
	if e.ID == uuid.Nil {
		return handling.NewHttpError404(nil, "file not found")
	}
	if err := r.db.WithContext(ctx).Delete(e).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return handling.NewHttpError404(err, "file not found")
		}
		return handling.NewHttpError500(err)
	}
	return nil
}
