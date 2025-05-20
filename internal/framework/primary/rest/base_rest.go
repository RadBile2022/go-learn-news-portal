package rest

import (
	"context"

	"github.com/go-chi/chi/v5"
)

type BaseRest[R any] interface {
	GetRoutePath() string

	Start(ctx context.Context) chi.Router
}
