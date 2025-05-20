package service

import (
	"context"
	"fmt"
	radstore "github.com/RadBile2022/go-library-radstore"
	"go-learn-news-portal/internal/core/entity"
	"go-learn-news-portal/internal/framework/secondary/repository"
	"go-learn-news-portal/library/v1/constant"
	"go-learn-news-portal/library/v1/convert"
	"go-learn-news-portal/library/v1/handling"
	"go-learn-news-portal/library/v1/middleware"
	"net/http"
	"slices"
)

type File interface {
	SumSizeByUploaderID(ctx context.Context, uID uint) (int64, error)

	FindByID(ctx context.Context, e *entity.File) error
	FindByName(ctx context.Context, e *entity.File) error
	FindByNames(ctx context.Context, names []string) ([]*entity.File, error)

	Create(ctx context.Context, e *entity.File) error
	//Delete(ctx context.Context, e *entity.File) error
	//DeleteByNames(ctx context.Context, filenames []string) error

	//GetFileObject(ctx context.Context, e *entity.File) ([]byte, error)
	//GetFileObjectReader(ctx context.Context, e *entity.File) (io.Reader, error)
	UploadFiles(ctx context.Context, fs []*entity.FileUpload) error
}

type fileOption func(fac *file)

func NewFile(opts ...fileOption) File {
	fac := &file{}

	for _, opt := range opts {
		opt(fac)
	}

	return fac
}

func FileWithFileRepo(repo repository.File) fileOption {
	return func(fac *file) {
		fac.repo = repo
	}
}

func FileWithStorage(storage radstore.Storage) fileOption {
	return func(fac *file) {
		fac.storage = storage
	}
}

type file struct {
	repo    repository.File
	storage radstore.Storage
}

func (c *file) SumSizeByUploaderID(ctx context.Context, uID uint) (int64, error) {
	return c.repo.SumSizeByUploaderID(ctx, uID)
}

func (c *file) FindByID(ctx context.Context, e *entity.File) error {
	return c.repo.FindByID(ctx, e)
}

func (c *file) FindByName(ctx context.Context, e *entity.File) error {
	return c.repo.FindByName(ctx, e)
}

func (c *file) FindByNames(ctx context.Context, names []string) ([]*entity.File, error) {
	return c.repo.FindByNames(ctx, names)
}

func (c *file) Create(ctx context.Context, e *entity.File) error {
	return c.repo.Create(ctx, e)
}

//func (c *file) Delete(ctx context.Context, e *entity.File) error {
//	var err error
//
//	if e.ID != uuid.Nil {
//		err = c.repo.FindByID(ctx, e)
//	} else if e.Name != "" {
//		err = c.repo.FindByName(ctx, e)
//	} else {
//		return handling.NewHttpError404(nil, "file not found")
//	}
//
//	if err != nil {
//		return err
//	}
//
//	if err = c.repo.Delete(ctx, e); err != nil {
//		return err
//	}
//
//	_ = c.storage.DeleteFileByKey(context.Background(), fmt.Sprintf("%s/%s", e.Path, e.Name))
//	return nil
//}

func (c *file) UploadFiles(ctx context.Context, fs []*entity.FileUpload) error {
	userId := middleware.GetUserIDFromContext(ctx)

	var fileMap *entity.FileMap
	var accFileSize int64
	var accFileSizeLimit int64
	var files []*entity.File
	var fileObjects []*radstore.FileHeader
	var fileNames []string

	accFileSize = 0
	accFileSizeLimit = 0

	for _, f := range fs {
		fileMap = getFileValidationFromPath(f.Path)

		if accFileSizeLimit == 0 {
			accFileSizeLimit = fileMap.LimitInMB * (1 << 20)
		}

		fileSize := f.MFileHeader.Size
		contentType := f.MFileHeader.Header.Get("Content-Type")

		fileExt := getFileExtFromMime(contentType)

		if !entity.FileTypeContains(fileMap.FileType, fileExt) {
			return handling.NewHttpError(nil, http.StatusBadRequest, fmt.Sprintf("file: %s format is not supported", f.Field), constant.ERR_INVALID_FILE_FORMAT)
		}

		accFileSize += fileSize

		if !fileMap.IsAccumulated && fileSize > accFileSizeLimit {
			return handling.NewHttpError(nil, http.StatusBadRequest, fmt.Sprintf("file: %s size is too big", f.Field), constant.ERR_FILE_SIZE_TOO_BIG)
		}

		name := convert.GenerateFileNameWithExt(userId, fileExt)
		f.Path = fmt.Sprintf("%d", userId)

		file := &entity.File{
			Name: name,
			Path: f.Path,
			Mime: contentType,
			Size: fileSize,
			Tags: []string{
				f.Path,
			},
			UploaderID: userId,
		}
		files = append(files, file)

		fha := &radstore.FileHeader{
			Filename:    name,
			Size:        f.Size,
			FileHandle:  radstore.FileFromMultipartHeader(f.MFileHeader),
			ContentType: f.MFileHeader.Header.Get("Content-Type"),
			Path:        f.Path,
		}

		fileObjects = append(fileObjects, fha)
		fileNames = append(fileNames, name)

		//return fmt.Sprintf(path, Filename), nil
		//
		//fileObject := &storage.FileObject{
		//	Name: fmt.Sprintf("%s/%s", file.Path, file.Name),
		//	Size: fileSize,
		//	File: storage.FileFromMultipartHeader(f.MFileHeader),
		//}
		//fileObjects = append(fileObjects, fileObject)
		path := "https://storage.radarcoding.my.id"

		f.ObjectName = fmt.Sprintf("%s/%s", path, file.Name)
		f.Size = fileSize
	}

	//if fileMap.IsAccumulated && accFileSize != 0 && accFileSize > accFileSizeLimit {
	//	return handling.NewHttpError(nil, http.StatusBadRequest, "file size is too big", constant.ERR_FILE_SIZE_TOO_BIG)
	//}
	//
	//todayUsage, err := c.SumSizeByUploaderID(ctx, user.ID)
	//if err != nil {
	//	return err
	//}

	//if (todayUsage + accFileSize) > user.DiskQuota {
	//	return handling.NewHttpError(nil, http.StatusUnprocessableEntity, "user disk's quota is exceeded", handling.ERR_USER_QUOTA_EXCEEDED)
	//}

	uploadCtx, cancel := context.WithCancel(context.Background())

	err := c.storage.UploadFiles(uploadCtx, fileObjects)
	if err != nil {
		cancel()
		return handling.NewHttpError500(err)
	}

	if err := c.repo.Create(ctx, files...); err != nil {
		cancel()
		c.storage.DeleteFiles(ctx, fileNames)
		return err
	}

	return nil
}

func getFileValidationFromPath(path string) *entity.FileMap {
	idx := slices.IndexFunc(entity.ValidFilesMapping, func(m entity.FileMap) bool {
		return m.Path == path
	})
	if idx < 0 {
		return nil
	}
	return &entity.ValidFilesMapping[idx]
}

func getFileExtFromMime(mime string) string {
	switch mime {
	case constant.FileTypePDF:
		return "pdf"
	case constant.FileTypeImageJpeg:
		return "jpeg"

	case constant.FileTypeImagePng:
		return "png"
	case constant.FileTypeImageJpg:
		return "jpg"
	}
	return ""
}
