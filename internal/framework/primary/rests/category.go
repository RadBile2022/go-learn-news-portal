package rests

import (
	"context"
	"github.com/go-chi/chi/v5"
	"go-learn-news-portal/internal/core/service"
	"go-learn-news-portal/internal/framework/primary/rests/request"
	"go-learn-news-portal/internal/framework/primary/rests/response"
	"go-learn-news-portal/library/v1/convert"
	"go-learn-news-portal/library/v1/router"
	"go-learn-news-portal/library/v1/validator"
	"net/http"
)

var defaultSuccessResponse response.DefaultSuccessResponse

func (h *categoryRest) Start(ctx context.Context) chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/", router.RootHandler(h.FindAllCategories).ToHandlerFunc())
		r.Post("/", router.RootHandler(h.CreateCategory).ToHandlerFunc())
		r.Get("/{id}", router.RootHandler(h.FindCategoryByID).ToHandlerFunc())
		r.Put("/{id}", router.RootHandler(h.UpdateCategory).ToHandlerFunc())
		r.Delete("/{id}", router.RootHandler(h.DeleteCategory).ToHandlerFunc())
	})
	return r
}

type Category interface {
	FindAllCategories(w http.ResponseWriter, r *http.Request) error
	FindCategoryByID(w http.ResponseWriter, r *http.Request) error
	CreateCategory(w http.ResponseWriter, r *http.Request) error
	UpdateCategory(w http.ResponseWriter, r *http.Request) error
	DeleteCategory(w http.ResponseWriter, r *http.Request) error

	FindCategoryFE(w http.ResponseWriter, r *http.Request) error
	Start(ctx context.Context) chi.Router
}

func NewCategory(categoryCore service.Category) Category {
	return &categoryRest{categoryCore: categoryCore}
}

type categoryRest struct {
	categoryCore service.Category
}

func (h *categoryRest) FindAllCategories(w http.ResponseWriter, r *http.Request) error {
	categories, err := h.categoryCore.FindAllCategories(r.Context())
	if err != nil {
		return err
	}

	response.SuccessData(categories, "Category fetched All successfully", w)

	return nil
}

func (h *categoryRest) FindCategoryByID(w http.ResponseWriter, r *http.Request) error {
	req := &request.CategoryID{ID: convert.PathValueIDInt64Chi(r)}

	category, err := h.categoryCore.FindCategoryByID(r.Context(), req)
	if err != nil {
		return err
	}

	response.SuccessData(category, "Category fetched successfully", w)

	return nil
}

func (h *categoryRest) CreateCategory(w http.ResponseWriter, r *http.Request) error {
	req, err := h.handleCreateOrUpdate(w, r)
	if err != nil {
		return err
	}

	if err = h.categoryCore.CreateCategory(r.Context(), req); err != nil {
		return err
	}

	response.Success("Created category successfully", w)

	return nil
}

func (h *categoryRest) UpdateCategory(w http.ResponseWriter, r *http.Request) error {
	categoryID := convert.PathValueIDInt64Chi(r)

	req, err := h.handleCreateOrUpdate(w, r, categoryID)
	if err != nil {
		return err
	}

	if err = h.categoryCore.UpdateCategoryByID(r.Context(), req); err != nil {
		return err
	}

	response.Success("Updated category successfully", w)

	return nil
}

func (h *categoryRest) handleCreateOrUpdate(w http.ResponseWriter, r *http.Request, IDs ...int64) (*request.CategoryForm, error) {
	var req = new(request.CategoryForm)

	if err := validator.Decode(r.Body, req, w); err != nil {
		return nil, err
	}
	if err := validator.Request(*req, w); err != nil {
		return nil, err
	}

	if len(IDs) > 0 {
		req.ID = IDs[0]
	}

	return req, nil
}

func (h *categoryRest) DeleteCategory(w http.ResponseWriter, r *http.Request) error {
	categoryID := convert.PathValueIDInt64Chi(r)
	req := &request.CategoryID{ID: categoryID}
	if err := h.categoryCore.DeleteCategory(r.Context(), req); err != nil {
		return err
	}

	response.Success("Updated category successfully", w)

	return nil
}

func (h *categoryRest) FindCategoryFE(w http.ResponseWriter, r *http.Request) error {
	categories, err := h.categoryCore.FindAllCategories(r.Context())
	if err != nil {
		return err
	}

	response.SuccessData(categories, "Updated category successfully", w)

	return nil
}
