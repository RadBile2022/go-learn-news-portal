package rests

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/RadBile2022/go-learn-news-portal/internal/core/service"
	"github.com/RadBile2022/go-learn-news-portal/internal/framework/primary/rests/request"
	"github.com/RadBile2022/go-learn-news-portal/internal/framework/primary/rests/response"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/middleware/ratelimiter"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/response_library"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/router"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/validator"
	"net/http"
	"time"
)

func (h *authRest) Start(ctx context.Context) chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(ratelimiter.Middleware(10, time.Minute))

		r.Post("/login", router.RootHandler(h.Login).ToHandlerFunc())
	})

	return r
}

type Auth interface {
	Login(w http.ResponseWriter, r *http.Request) error
	Start(ctx context.Context) chi.Router
}

func NewAuth(core service.Auth) Auth {
	return &authRest{authCore: core}
}

type authRest struct {
	authCore service.Auth
}

func (h *authRest) Login(w http.ResponseWriter, r *http.Request) error {
	var req request.Login
	if err := validator.Decode(r.Body, &req, w); err != nil {
		return err
	}

	if err := validator.Request(req, w); err != nil {
		return err
	}

	auth, err := h.authCore.Login(r.Context(), &req)
	if err != nil {
		return err
	}

	resp := response.SuccessAuthResponse{}

	resp.Meta.Status = true
	resp.Meta.Message = "Login successful"
	resp.AccessToken = auth.AccessToken
	resp.ExpiresAt = 345435

	response_library.JSON(resp, http.StatusOK, w)
	//response_library.SuccessLogin("Login Successfully", auth.AccessToken, auth.RefreshToken, w)
	return nil
}

// UpdatePassword implements UserHandler.
