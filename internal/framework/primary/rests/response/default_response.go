package response

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type ErrorResponseDefault struct {
	Meta
}

type Meta struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type DefaultSuccessResponse struct {
	Meta       Meta                `json:"meta"`
	Data       any                 `json:"data,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

type PaginationResponse struct {
	TotalRecords int `json:"total_records"`
	Page         int `json:"page"`
	PerPage      int `json:"per_page"`
	TotalPages   int `json:"total_pages"`
}

func SuccessPaginate(data any, msg string, pg *PaginationResponse, w http.ResponseWriter) {
	response := DefaultSuccessResponse{
		Data: data,
		Meta: Meta{
			Message: msg,
			Status:  true,
		},
		Pagination: pg,
	}
	JSON(response, http.StatusOK, w)
}

func SuccessData(data any, msg string, w http.ResponseWriter) {
	response := DefaultSuccessResponse{
		Data: data,
		Meta: Meta{
			Message: msg,
			Status:  true,
		},
	}
	JSON(response, http.StatusOK, w)
}

func Success(msg string, w http.ResponseWriter) {
	response := DefaultSuccessResponse{
		Data: nil,
		Meta: Meta{
			Message: msg,
			Status:  true,
		},
	}
	JSON(response, http.StatusOK, w)
}

func JSON(response any, statusCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("error", slog.Any("error", err))
		return
	}
}
