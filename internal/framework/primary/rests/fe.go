package rests

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/RadBile2022/go-learn-news-portal/internal/core/service"
	"github.com/RadBile2022/go-learn-news-portal/internal/framework/primary/rests/request"
	"github.com/RadBile2022/go-learn-news-portal/internal/framework/primary/rests/response"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/convert"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/pagination"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/response_library"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/router"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/validator"
	"net/http"
)

func (h *frontEndRest) Start(ctx context.Context) chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {

		r.Get("/categories", router.RootHandler(h.FindAllCategories).ToHandlerFunc())
		r.Get("/contents", router.RootHandler(h.FindAllContents).ToHandlerFunc())
		r.Get("/contents/{id}", router.RootHandler(h.FindContentByID).ToHandlerFunc())
	})

	return r
}

type FrontEnd interface {
	FindAllCategories(w http.ResponseWriter, r *http.Request) error
	FindAllContents(w http.ResponseWriter, r *http.Request) error
	FindContentByID(w http.ResponseWriter, r *http.Request) error
	Start(ctx context.Context) chi.Router
}

func NewFrontEnd(category service.Category, content service.Content) FrontEnd {
	return &frontEndRest{categoryService: category, contentService: content}
}

type frontEndRest struct {
	categoryService service.Category
	contentService  service.Content
}

func (h *frontEndRest) FindAllCategories(w http.ResponseWriter, r *http.Request) error {
	categories, err := h.categoryService.FindAllCategories(r.Context())
	if err != nil {
		return err
	}

	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Pagination = nil
	defaultSuccessResponse.Meta.Message = "Categories fetched successfully"
	defaultSuccessResponse.Data = categories

	response_library.JSON(defaultSuccessResponse, http.StatusOK, w)
	return nil
}

func (h *frontEndRest) FindAllContents(w http.ResponseWriter, r *http.Request) error {
	var err error
	queries := pagination.NewQueriesNetHTTP(r)

	if err = validator.Request(queries, w); err != nil {
		return err
	}

	pg := pagination.NewPagination(r)

	contents, err := h.contentService.FindAllContents(r.Context(), queries, pg)
	if err != nil {
		return err
	}

	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Pagination = &response.PaginationResponse{
		TotalRecords: int(pg.TotalRows),
		TotalPages:   pg.TotalPages,
		Page:         pg.Page,
		PerPage:      pg.GetLimit(),
	}
	defaultSuccessResponse.Meta.Message = "Category fetched successfully"
	defaultSuccessResponse.Data = contents

	response_library.JSON(defaultSuccessResponse, http.StatusOK, w)

	return nil
}

func (h *frontEndRest) FindContentByID(w http.ResponseWriter, r *http.Request) error {
	req := &request.ContentID{ID: convert.PathValueIDInt64Chi(r)}

	category, err := h.contentService.FindContentByID(r.Context(), req)
	if err != nil {
		return err
	}

	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Pagination = nil
	defaultSuccessResponse.Meta.Message = "Category fetched successfully"
	defaultSuccessResponse.Data = category

	response_library.JSON(defaultSuccessResponse, http.StatusOK, w)

	return nil
}
