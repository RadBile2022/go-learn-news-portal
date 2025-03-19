package router

import (
	"go-learn-news-portal/library/v1/response_library"
	"net/http"
)

type RootHandler func(http.ResponseWriter, *http.Request) error

func (h RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if err == nil {
		return
	}

	response_library.Error(err, w)
}

func (h RootHandler) ToHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			response_library.Error(err, w)
		}
	}
}
