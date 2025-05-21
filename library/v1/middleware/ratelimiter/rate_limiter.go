package ratelimiter

import (
	"github.com/go-chi/httprate"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/constant"
	response "github.com/RadBile2022/go-learn-news-portal/library/v1/response_library"
	"net/http"
	"time"
)

func Middleware(rate int, period time.Duration) func(http.Handler) http.Handler {
	limitHandler := func(w http.ResponseWriter, r *http.Request) {
		res := response.Response{
			Meta: response.ResponseMeta{
				Message:   "too many requests",
				Code:      http.StatusTooManyRequests,
				ErrorCode: constant.ERR_TOO_MANY_REQUESTS,
			},
		}

		response.JSON(res, http.StatusTooManyRequests, w)
	}

	keyFuncs := []httprate.KeyFunc{
		httprate.KeyByIP,
		httprate.KeyByEndpoint,
		httprate.KeyByRealIP,
	}

	return httprate.Limit(
		rate,
		period,
		httprate.WithKeyFuncs(keyFuncs...),
		httprate.WithLimitHandler(limitHandler),
	)
}
