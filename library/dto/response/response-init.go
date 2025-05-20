package response

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type ResponseMeta struct {
	Message   string `json:"message"`
	Code      int    `json:"code"`
	ErrorCode string `json:"error_code,omitempty"`
}

type Response struct {
	Data any          `json:"data,omitempty"`
	Meta ResponseMeta `json:"meta,omitempty"`
}

type PaginationResponse struct {
	Data map[string]any `json:"data"`
	Page Page           `json:"page"`
	Meta ResponseMeta   `json:"meta,omitempty"`
}

type Page struct {
	CurrentPage int   `json:"current_page,omitempty"`
	TotalData   int64 `json:"total_data,omitempty"`
	Limit       int   `json:"limit,omitempty"`
	TotalPage   int   `json:"total_page,omitempty"`
}

func Json(response any, statusCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("error", slog.Any("error", err))
		return
	}
}
