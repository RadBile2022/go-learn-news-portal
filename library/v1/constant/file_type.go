package constant

const (
	FileTypePDF       = "application/pdf"
	FileTypeImagePng  = "image/png"
	FileTypeImageJpeg = "image/jpeg"
	FileTypeImageJpg  = "image/jpg"
)

var ValidFileUploadTypes = []string{FileTypePDF, FileTypeImageJpeg, FileTypeImagePng, FileTypeImageJpg}
