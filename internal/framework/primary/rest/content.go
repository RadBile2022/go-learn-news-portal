package rest

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"go-learn-news-portal/internal/core/service"
	"go-learn-news-portal/internal/framework/primary/rest/request"
	"go-learn-news-portal/internal/framework/primary/rest/response"
	"go-learn-news-portal/library/v1/convert"
	"go-learn-news-portal/library/v1/pagination"
	"go-learn-news-portal/library/v1/router"
	"go-learn-news-portal/library/v1/validator"
	"net/http"
)

func (h *contentRest) Start(ctx context.Context) chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/upload-image", router.RootHandler(h.UploadImageR2).ToHandlerFunc())
		r.Post("/", router.RootHandler(h.CreateContent).ToHandlerFunc())
		r.Get("/", router.RootHandler(h.FindAllContents).ToHandlerFunc())
		r.Get("/{id}", router.RootHandler(h.FindContentByID).ToHandlerFunc())
		r.Put("/{id}", router.RootHandler(h.UpdateContent).ToHandlerFunc())
		r.Delete("/{id}", router.RootHandler(h.DeleteContent).ToHandlerFunc())
	})
	return r
}

type Content interface {
	Start(ctx context.Context) chi.Router
	FindAllContents(w http.ResponseWriter, r *http.Request) error
	FindContentByID(w http.ResponseWriter, r *http.Request) error
	CreateContent(w http.ResponseWriter, r *http.Request) error
	UpdateContent(w http.ResponseWriter, r *http.Request) error
	DeleteContent(w http.ResponseWriter, r *http.Request) error
	UploadImageR2(w http.ResponseWriter, r *http.Request) error
}

func NewContent(contentCore service.Content) Content {
	return &contentRest{contentService: contentCore}
}

type contentRest struct {
	contentService service.Content
}

func (h *contentRest) FindAllContents(w http.ResponseWriter, r *http.Request) error {
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

	response.SuccessPaginate(contents, "Fetch All Content is successfully", &response.PaginationResponse{
		TotalRecords: int(pg.TotalRows),
		TotalPages:   pg.TotalPages,
		Page:         pg.Page,
		PerPage:      pg.GetLimit(),
	}, w)

	return nil
}

func (h *contentRest) FindContentByID(w http.ResponseWriter, r *http.Request) error {
	req := &request.ContentID{ID: convert.PathValueIDInt64Chi(r)}

	category, err := h.contentService.FindContentByID(r.Context(), req)
	if err != nil {
		return err
	}

	response.SuccessData(category, "Fetch Content by ID is successfully", w)

	return nil
}

func (h *contentRest) CreateContent(w http.ResponseWriter, r *http.Request) error {
	req, err := h.handleCreateOrUpdate(w, r)
	if err != nil {
		return err
	}

	if err = h.contentService.CreateContent(r.Context(), req); err != nil {
		return err
	}

	response.Success("Create Content is successfully", w)

	return nil
}

func (h *contentRest) UpdateContent(w http.ResponseWriter, r *http.Request) error {
	categoryID := convert.PathValueIDInt64Chi(r)
	req, err := h.handleCreateOrUpdate(w, r, categoryID)
	if err != nil {
		return err
	}

	if err = h.contentService.UpdateContent(r.Context(), req); err != nil {
		return err
	}
	if err != nil {
		return err
	}

	response.Success("Update Content is successfully", w)

	return nil
}

func (h *contentRest) handleCreateOrUpdate(w http.ResponseWriter, r *http.Request, IDs ...int64) (*request.ContentForm, error) {
	var req = new(request.ContentForm)

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

func (h *contentRest) DeleteContent(w http.ResponseWriter, r *http.Request) error {
	categoryID := convert.PathValueIDInt64Chi(r)
	req := &request.ContentID{ID: categoryID}

	err := h.contentService.DeleteContent(r.Context(), req)
	if err != nil {
		return err
	}

	response.Success("Delete Content is successfully", w)
	return nil
}

func (h *contentRest) UploadImageR2(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		return err
	}

	fh, boolean := r.MultipartForm.File["image"]

	if !boolean {
		return errors.New("failed to read file")
	}

	fileName, err := h.contentService.UploadImageR2(r.Context(), fh)
	if err != nil {
		return err
	}

	urlImageResp := map[string]interface{}{
		"urlImage": fileName,
	}

	response.SuccessData(urlImageResp, "Upload Image is successfully", w)

	return nil
}
