package service

import (
	"context"
	"fmt"
	radstore "github.com/RadBile2022/go-library-radstore"
	"go-learn-news-portal/internal/framework/primary/rests/request"
	"go-learn-news-portal/internal/framework/primary/rests/response"
	"go-learn-news-portal/internal/framework/secondary/repository"
	"go-learn-news-portal/library/v1/handling"
	"go-learn-news-portal/library/v1/middleware"
	"go-learn-news-portal/library/v1/pagination"
	"mime/multipart"
	"time"
)

type Content interface {
	FindAllContents(ctx context.Context, q *pagination.Queries, pg *pagination.Pagination) ([]*response.Content, error)
	FindContentByID(ctx context.Context, req *request.ContentID) (*response.Content, error)
	CreateContent(ctx context.Context, req *request.ContentForm) error
	UpdateContent(ctx context.Context, req *request.ContentForm) error
	DeleteContent(ctx context.Context, req *request.ContentID) error
	UploadImageR2(ctx context.Context, reqs []*multipart.FileHeader) (string, error)
}

func NewContent(repo repository.Content, storage radstore.Storage) Content {
	return &contentService{
		contentRepo: repo,
		storage:     storage,
	}
}

type contentService struct {
	contentRepo repository.Content
	storage     radstore.Storage
}

func (c *contentService) FindAllContents(ctx context.Context, q *pagination.Queries, pg *pagination.Pagination) ([]*response.Content, error) {
	contents, err := c.contentRepo.FindAllContents(ctx, q, pg)
	if err != nil {
		return nil, err
	}
	return response.FromEntitiesContent(contents), nil
}

func (c *contentService) FindContentByID(ctx context.Context, req *request.ContentID) (*response.Content, error) {
	e := req.ToEntity()

	if err := c.contentRepo.FindContentByID(ctx, e); err != nil {
		return nil, err
	}

	return response.FromEntityContent(e), nil
}

func (c *contentService) CreateContent(ctx context.Context, req *request.ContentForm) error {
	return c.contentRepo.CreateContent(ctx, req.ToEntity(ctx))
}

func (c *contentService) UpdateContent(ctx context.Context, req *request.ContentForm) error {
	return c.contentRepo.UpdateContent(ctx, req.ToEntity(ctx))
}

func (c *contentService) DeleteContent(ctx context.Context, req *request.ContentID) error {
	return c.contentRepo.DeleteContent(ctx, req.ToEntity())
}

func (c *contentService) UploadImageR2(ctx context.Context, fhs []*multipart.FileHeader) (string, error) {
	userId := middleware.GetUserIDFromContext(ctx)
	path := "https://storage.radarcoding.my.id/%s"
	Filename := fmt.Sprintf("%d-%d", userId, time.Now().UnixNano())

	fha := &radstore.FileHeader{
		Filename:    Filename,
		Size:        fhs[0].Size,
		FileHandle:  radstore.FileFromMultipartHeader(fhs[0]),
		ContentType: fhs[0].Header.Get("Content-Type"),
	}

	err := c.storage.UploadFiles(ctx, []*radstore.FileHeader{fha})
	if err != nil {
		return "", handling.NewHttpError500(err)
	}

	return fmt.Sprintf(path, Filename), nil
}
