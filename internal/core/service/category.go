package service

import (
	"context"
	"go-learn-news-portal/internal/framework/primary/rest/request"
	"go-learn-news-portal/internal/framework/primary/rest/response"
	"go-learn-news-portal/internal/framework/secondary/repository"
)

type Category interface {
	FindAllCategories(ctx context.Context) ([]*response.Category, error)
	FindCategoryByID(ctx context.Context, req *request.CategoryID) (*response.Category, error)
	CreateCategory(ctx context.Context, req *request.CategoryForm) error
	UpdateCategoryByID(ctx context.Context, req *request.CategoryForm) error
	DeleteCategory(ctx context.Context, req *request.CategoryID) error
}

func NewCategory(categoryRepo repository.Category) Category {
	return &categoryService{categoryRepository: categoryRepo}
}

type categoryService struct {
	categoryRepository repository.Category
}

func (c *categoryService) FindAllCategories(ctx context.Context) ([]*response.Category, error) {
	categories, err := c.categoryRepository.FindAllCategories(ctx)
	if err != nil {
		return nil, err
	}

	return response.FromEntitiesCategory(categories), nil
}

func (c *categoryService) FindCategoryByID(ctx context.Context, req *request.CategoryID) (*response.Category, error) {
	e := req.ToEntity()

	if err := c.categoryRepository.FindCategoryByID(ctx, e); err != nil {
		return nil, err
	}

	return response.FromEntityCategory(e), nil
}

func (c *categoryService) CreateCategory(ctx context.Context, req *request.CategoryForm) error {
	// terkena aturan unik slug, kalau bisa di find dulu
	return c.categoryRepository.CreateCategory(ctx, req.ToEntity(ctx))
}

func (c *categoryService) UpdateCategoryByID(ctx context.Context, req *request.CategoryForm) error {
	return c.categoryRepository.UpdateCategoryByID(ctx, req.ToEntity(ctx))
}

func (c *categoryService) DeleteCategory(ctx context.Context, req *request.CategoryID) error {
	return c.categoryRepository.DeleteCategory(ctx, req.ToEntity())
}
