package validator

import (
	"encoding/json"
	"fmt"
	"go-learn-news-portal/library/v1/constant"
	"go-learn-news-portal/library/v1/handling"
	response "go-learn-news-portal/library/v1/response_library"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	"golang.org/x/exp/slices"
)

var invalidRequestRes = response.Response2{
	Message: "Request is not valid",
	Data:    nil,
}

func DecodeFormData(w http.ResponseWriter, r *http.Request, fileNames []string) error {
	for _, field := range fileNames {
		uploadedFile, fileHeader, err := r.FormFile(field)
		if err != nil {
			return handling.NewHttpError(err, http.StatusBadRequest, fmt.Sprintf("field %s is required or it's not a file", field), constant.ERR_INVALID_FILE_FORMAT)
		}
		if !slices.Contains(constant.ValidFileUploadTypes, fileHeader.Header.Get("Content-Type")) {
			return handling.NewHttpError(nil, http.StatusBadRequest, fmt.Sprintf("file: %s format is not supported", field), constant.ERR_INVALID_FILE_FORMAT)
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				response.Error(err, w)
			}
		}(uploadedFile)
	}

	return nil
}

func Param1Int(key string, w http.ResponseWriter, r *http.Request) (int, error) {
	param1Str := r.URL.Query().Get(key)
	if param1Str == "" {
		response.JSON(invalidRequestRes, http.StatusBadRequest, w)
		return 0, handling.InternalServerError()
	}

	param1, err := strconv.Atoi(param1Str)
	if err != nil {
		response.JSON(invalidRequestRes, http.StatusBadRequest, w)
		return 0, handling.InternalServerError()
	}
	return param1, nil
}

func Param1Uint(key string, w http.ResponseWriter, r *http.Request) (uint, error) {
	param1Str := r.URL.Query().Get(key)
	if param1Str == "" {
		response.JSON(invalidRequestRes, http.StatusBadRequest, w)
		return 0, handling.InternalServerError()
	}

	param1, err := strconv.ParseUint(param1Str, 10, 32)
	if err != nil {
		response.JSON(invalidRequestRes, http.StatusBadRequest, w)
		return 0, handling.InternalServerError()
	}
	return uint(param1), nil
}

func Param1String(key string, w http.ResponseWriter, r *http.Request) (string, error) {
	param1Str := r.URL.Query().Get(key)
	if param1Str == "" {
		response.JSON(invalidRequestRes, http.StatusBadRequest, w)
		return "", handling.InternalServerError()
	}

	return param1Str, nil
}

func Decode(b io.Reader, v any, w http.ResponseWriter) error {
	err := json.NewDecoder(b).Decode(v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return handling.NewHttpError(err, http.StatusBadRequest, "request body is invalid", constant.ERR_JSON_BODY)
	}
	return nil
}

func Request(s any, w http.ResponseWriter) error {
	if err := validator(s); err != nil {
		return err
	}
	return nil
}
