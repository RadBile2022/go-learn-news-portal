package handling

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/constant"
	"net/http"
	"runtime/debug"
)

type ClientError interface {
	Error() string
	GetErrorCode() string
	GetMessage() string
	GetStatus() int
}

func NewHttpError500(err error) error {
	return NewHttpError(err, http.StatusInternalServerError, "Internal server error", constant.ERR_INTERNAL_SERVER_ERROR)
}

func NewHttpError400(err error) error {
	return NewHttpError(err, http.StatusBadRequest, "Bad Request", constant.ERR_BAD_REQUEST)
}

func NewHttpError401(err error) error {
	return NewHttpError(err, http.StatusUnauthorized, "User unauthorized", constant.ERR_UNAUTHORIZED)
}

func NewHttpError403(err error) error {
	return NewHttpError(err, http.StatusForbidden, "Forbidden", constant.ERR_FORBIDDEN)
}

func NewHttpError404(err error, message string) error {
	return NewHttpError(err, http.StatusNotFound, message, constant.ERR_RECORD_NOT_FOUND)
}

func NewHttpError(err error, status int, message string, errorCode string) error {
	if err != nil && status == http.StatusInternalServerError {
		fmt.Printf("error occured: %v\n", err)
		debug.PrintStack()
		sentry.CaptureException(err)
	}

	return &HttpError{
		Cause:     err,
		Status:    status,
		ErrorCode: errorCode,
		Message:   message,
	}
}

type HttpError struct {
	Cause     error
	Message   string
	ErrorCode string
	Status    int
}

func (e *HttpError) Error() string {
	if e.Cause == nil {
		return e.Message
	}
	return e.Message + " : " + e.Cause.Error()
}

func (e *HttpError) GetMessage() string {
	if e.Message == "" {
		return e.Error()
	}
	return e.Message
}

func (e *HttpError) GetStatus() int {
	return e.Status
}

func (e *HttpError) GetErrorCode() string {
	return e.ErrorCode
}
