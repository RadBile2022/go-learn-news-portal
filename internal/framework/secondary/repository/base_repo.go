package repository

import (
	"context"

	"gorm.io/gorm"
)

type RepoFactory[T comparable] func(base *baseRepo[T]) T

type BaseRepo[T comparable] interface {
	WithTrx(db *gorm.DB) T
	Custom(ctx context.Context, fn func(db *gorm.DB) *gorm.DB) error
	CustomError(ctx context.Context, fn func(db *gorm.DB) error) error
}

func newRepository[T comparable](fac RepoFactory[T], db *gorm.DB) T {
	b := newBaseRepo[T](db)
	t := fac(b)
	b.fac = fac
	return t
}

func newBaseRepo[T comparable](db *gorm.DB) *baseRepo[T] {
	return &baseRepo[T]{db: db}
}

type baseRepo[T comparable] struct {
	db  *gorm.DB
	fac RepoFactory[T]
}

func (c *baseRepo[T]) WithTrx(db *gorm.DB) T {
	return c.fac(newBaseRepo[T](db))
}

func (b *baseRepo[T]) Custom(ctx context.Context, fn func(db *gorm.DB) *gorm.DB) error {
	if err := fn(b.db.WithContext(ctx)).Error; err != nil {
		return err
	}
	return nil
}

func (c *baseRepo[T]) CustomError(ctx context.Context, fn func(db *gorm.DB) error) error {
	if err := fn(c.db.WithContext(ctx)); err != nil {
		return err
	}
	return nil
}

func (b *baseRepo[T]) softDelete(ctx context.Context, model any, deletedBy uint) error {
	err := b.db.WithContext(ctx).
		Model(model).
		Updates(map[string]any{
			"deleted_by": deletedBy,
			"deleted_at": gorm.Expr("NOW()"),
		}).Error
	if err != nil {
		return err
	}
	return nil
}
