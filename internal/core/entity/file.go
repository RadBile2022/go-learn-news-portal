package entity

import (
	"golang.org/x/exp/slices"
	"mime/multipart"
)

type File struct {
	BaseModel

	UploaderID int64 `json:"uploader_id" gorm:"not_null"`

	Name    string `json:"name" gorm:"not null;type:varchar(100);unique;"`
	Path    string `json:"path" gorm:"not null;type:varchar(50);"`
	SubPath string `json:"sub_path" gorm:"not null;type:varchar(50);"`

	Size int64  `json:"size" gorm:"not null;default:0;"`
	Mime string `json:"mime" gorm:"not null;type:varchar(100);"`

	Tags []string `json:"tags" gorm:"type:jsonb;serializer:json;"`
}

func (File) TableName() string {
	return "app_news_files"
}

type FileUpload struct {
	ObjectName  string
	Size        int64
	Path        string
	Field       string
	MFileHeader *multipart.FileHeader
}

type FileMap struct {
	FileType      []string
	Path          string
	LimitInMB     int64
	IsAccumulated bool
}

var ValidFilesMapping = []FileMap{
	{FileType: []string{"pdf", "png", "jpeg"}, Path: "radstore", LimitInMB: 5, IsAccumulated: false},
}

func FileTypeContains(arr []string, str string) bool {
	return slices.Contains(arr, str)
}
