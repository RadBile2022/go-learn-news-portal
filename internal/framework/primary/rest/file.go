package rest

import (
	"context"
	"github.com/RadBile2022/go-learn-news-portal/internal/core/entity"
	"github.com/RadBile2022/go-learn-news-portal/internal/core/service"
	"github.com/RadBile2022/go-learn-news-portal/library/dto/response"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/constant"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/handling"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/router"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/validator"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type File interface {
	//BaseRest[File]

	BulkUpload(w http.ResponseWriter, r *http.Request) error

	Start(ctx context.Context) chi.Router
}

func NewFile(api service.File) File {
	return &fileRest{
		service: api,
	}
}

type fileRest struct {
	service service.File
}

// GetRoutePath implements File.
//func (c *fileRest) GetRoutePath() string {
//	return "/files"
//}

func (c *fileRest) BulkUpload(w http.ResponseWriter, r *http.Request) error {
	//path := r.PathValue("path")
	//if !slices.ContainsFunc(entity.ValidFilesMapping, func(f entity.FileMap) bool {
	//	return f.Path == path
	//}) {
	//	return handling.NewHttpError(nil, http.StatusBadRequest, fmt.Sprintf("path: %s is invalid", path), constant.ERR_VALIDATION_ERROR)
	//}
	fieldStr := r.FormValue("fields")
	fields := strings.Split(fieldStr, ",")

	if len(fields) == 0 {
		return handling.NewHttpError(nil, http.StatusBadRequest, "fields is required", constant.ERR_VALIDATION_ERROR)
	}

	if err := validator.DecodeFormData(w, r, fields); err != nil {
		return err
	}

	var files []*entity.FileUpload
	for _, field := range fields {
		if fhs := r.MultipartForm.File[field]; len(fhs) > 0 {
			files = append(files, &entity.FileUpload{
				Field:       field,
				Path:        "radstore",
				MFileHeader: fhs[0],
			})
		}
	}

	if err := c.service.UploadFiles(r.Context(), files); err != nil {
		return err
	}

	result := map[string]any{}
	for idx, field := range fields {
		file := files[idx]
		result[field] = map[string]any{
			"filename": file.ObjectName,
			"filesize": file.Size,
		}
	}

	response.SuccessData("files uploaded succesfully", "uploaded_files", result, w)
	return nil
}

func (c *fileRest) Start(ctx context.Context) chi.Router {
	r := chi.NewRouter()

	r.Method(http.MethodPost, "/bulk_upload", router.RootHandler(c.BulkUpload))

	return r
}
