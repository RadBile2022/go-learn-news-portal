package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
)

type UploadResponse struct {
	Data struct {
		UploadedFile map[string]UploadedFile `json:"uploaded_files"`
	} `json:"data"`
	Meta struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"meta"`
}

type UploadedFile struct {
	Filename string `json:"filename"`
	Filesize int    `json:"filesize"`
}

type UploadSource struct {
	LocalFiles    map[string]string                  // file lokal
	UploadedFiles map[string][]*multipart.FileHeader // dari frontend/postman
}

func Upload(rootUrl, token string, files []multipart.FileHeader) (map[string]UploadedFile, error) {
	// Auto-map jadi file1, file2, ...
	mapped := make(map[string][]*multipart.FileHeader)
	fields := make([]string, 0, len(files))

	for i := range files {
		field := fmt.Sprintf("file%d", i+1)
		mapped[field] = []*multipart.FileHeader{&files[i]}
		fields = append(fields, field)
	}

	uploadSource := UploadSource{
		UploadedFiles: mapped,
	}

	// Langsung pakai existing PrepareMultipartUnified
	body, contentType, err := PrepareMultipartUnified(fields, uploadSource)
	if err != nil {
		return nil, err
	}

	fileNames, err := EndpointUpload(rootUrl, token, body, contentType)
	if err != nil {
		fmt.Println("Upload Error:", err)
		return nil, err
	}
	return fileNames, nil
}

func PrepareMultipartUnified(fields []string, source UploadSource) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Kirim "fields"
	writer.WriteField("fields", strings.Join(fields, ","))

	// Dari file lokal
	if len(source.LocalFiles) > 0 {
		for field, path := range source.LocalFiles {
			data, err := os.ReadFile(path)
			if err != nil {
				return nil, "", fmt.Errorf("gagal membaca file %s: %w", path, err)
			}

			mimeType := getMimeType(path)
			partHeader := textproto.MIMEHeader{}
			partHeader.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, field, filepath.Base(path)))
			partHeader.Set("Content-Type", mimeType)

			part, err := writer.CreatePart(partHeader)
			if err != nil {
				return nil, "", fmt.Errorf("gagal membuat part: %w", err)
			}
			if _, err := part.Write(data); err != nil {
				return nil, "", fmt.Errorf("gagal menulis data: %w", err)
			}
		}
	}

	// Dari multipart upload (misalnya dari handler Gin/Fiber)
	// Dari multipart upload (misalnya dari handler Gin/Fiber)
	if len(source.UploadedFiles) > 0 {
		for field, headers := range source.UploadedFiles {
			for _, fh := range headers {
				file, err := fh.Open()
				if err != nil {
					return nil, "", fmt.Errorf("tidak bisa membuka file: %w", err)
				}
				defer file.Close()

				mimeType := getMimeType(fh.Filename)
				partHeader := textproto.MIMEHeader{}
				partHeader.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, field, fh.Filename))
				partHeader.Set("Content-Type", mimeType)

				part, err := writer.CreatePart(partHeader)
				if err != nil {
					return nil, "", fmt.Errorf("tidak bisa membuat part: %w", err)
				}
				if _, err := io.Copy(part, file); err != nil {
					return nil, "", fmt.Errorf("gagal menyalin isi file: %w", err)
				}
			}
		}
	}

	if err := writer.Close(); err != nil {
		return nil, "", err
	}

	return body, writer.FormDataContentType(), nil
}

func getMimeType(filename string) string {
	ext := filepath.Ext(filename)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		switch ext {
		case ".docx":
			mimeType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
		case ".pdf":
			mimeType = "application/pdf"
		// tambahkan lainnya sesuai kebutuhan
		default:
			mimeType = "application/octet-stream"
		}
	}

	fmt.Println("MimeType:", mimeType)
	return mimeType
}

func EndpointUpload(rootUrl, token string, body *bytes.Buffer, contentType string) (map[string]UploadedFile, error) {
	url := fmt.Sprintf("%s/api/admin/files/bulk_upload", rootUrl)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", contentType)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("Status:", resp.Status)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("gagal upload: %s", string(respBody))
	}

	var uploadResp UploadResponse
	if err := json.Unmarshal(respBody, &uploadResp); err != nil {
		return nil, fmt.Errorf("gagal parse response upload: %w", err)
	}

	fmt.Println("Pesan:", uploadResp.Meta.Message)
	for field, file := range uploadResp.Data.UploadedFile {
		fmt.Printf("Field: %s\n  - URL: %s\n  - Size: %d bytes\n", field, file.Filename, file.Filesize)
	}

	return uploadResp.Data.UploadedFile, nil
}
