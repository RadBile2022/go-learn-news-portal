package rests

import (
	"context"
	"github.com/go-chi/chi/v5"
	"go-learn-news-portal/internal/core/service"
	"go-learn-news-portal/internal/framework/primary/rests/request"
	"go-learn-news-portal/library/v1/response_library"
	"go-learn-news-portal/library/v1/router"
	"go-learn-news-portal/library/v1/validator"
	"net/http"
)

func (h *userRest) Start(ctx context.Context) chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Put("/update-password", router.RootHandler(h.UpdatePassword).ToHandlerFunc())
		r.Get("/profile", router.RootHandler(h.FindUserByID).ToHandlerFunc())
	})
	return r
}

type User interface {
	FindUserByID(w http.ResponseWriter, r *http.Request) error
	UpdatePassword(w http.ResponseWriter, r *http.Request) error
	Start(ctx context.Context) chi.Router
}

func NewUser(core service.User) User {
	return &userRest{coreUser: core}
}

type userRest struct {
	coreUser service.User
}

// FindUserByID implements UserHandler.
func (h *userRest) FindUserByID(w http.ResponseWriter, r *http.Request) error {

	user, err := h.coreUser.FindUserByID(r.Context())
	if err != nil {
		return err
	}

	response_library.SuccessData("Find User Successfully", user, w)
	return nil
}

func (h *userRest) UpdatePassword(w http.ResponseWriter, r *http.Request) error {

	var (
		req request.UpdatePassword
		err error
	)

	err = validator.Decode(r.Body, &req, w)
	if err != nil {
		return err
	}

	err = validator.Request(req, w)
	if err != nil {
		return err
	}

	err = h.coreUser.UpdatePassword(r.Context(), &req)
	if err != nil {
		return err
	}

	response_library.Success("Update Password Successfully", w)

	return nil
}
