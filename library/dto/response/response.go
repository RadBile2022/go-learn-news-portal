package response

import (
	"context"
	"errors"
	"go-learn-news-portal/library/v1/constant"
	"go-learn-news-portal/library/v1/handling"
	"log/slog"
	"net/http"
	"os"
)

//var reportedCodes = types.NewSet(
//	http.StatusInternalServerError,
//	http.StatusBadGateway,
//	http.StatusRequestTimeout,
//	http.StatusServiceUnavailable,
//	http.StatusGatewayTimeout,
//	http.StatusTooManyRequests,
//)

func Error(ctx context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		return
	}

	var res Response
	clientErr, isClientErr := err.(handling.ClientError)

	timeoutRes := Response{
		Meta: ResponseMeta{
			Message: "Request timeout. Please try again.",
			Code:    http.StatusRequestTimeout,
			//ErrorCode: constraint.ERR_REQUEST_TIMEOUT,
		},
	}

	if errors.Is(err, context.DeadlineExceeded) {
		res = timeoutRes
	} else if os.IsTimeout(err) {
		res = timeoutRes
	} else if isClientErr {
		res = Response{
			Meta: ResponseMeta{
				Message:   clientErr.GetMessage(),
				Code:      clientErr.GetStatus(),
				ErrorCode: clientErr.GetErrorCode(),
			},
		}
	} else {
		slog.Error("http error occurred", slog.String("error", err.Error()))
		res = Response{
			Meta: ResponseMeta{
				Message:   "Internal server error",
				Code:      http.StatusInternalServerError,
				ErrorCode: constant.ERR_INTERNAL_SERVER_ERROR,
			},
		}
	}

	//if reportedCodes.Contains(res.Meta.Code) {
	//	sentryHub := sentry.GetHubFromContext(ctx)
	//	if sentryHub != nil {
	//		sentryHub.CaptureException(err)
	//	}
	//}

	Json(res, res.Meta.Code, w)
}

func Success(msg string, w http.ResponseWriter) {
	response := Response{
		Data: nil,
		Meta: ResponseMeta{
			Message: msg,
			Code:    http.StatusOK,
		},
	}
	Json(response, http.StatusOK, w)
}

func SuccessCreate(msg, dataName string, data any, w http.ResponseWriter) {
	response := Response{
		Data: map[string]any{
			dataName: data,
		},
		Meta: ResponseMeta{
			Message: msg,
			Code:    http.StatusCreated,
		},
	}
	Json(response, http.StatusCreated, w)
}

func SuccessDataPaginate(msg, dataName string, data any, page *Page, w http.ResponseWriter) {
	response := PaginationResponse{
		Data: map[string]any{
			dataName: data,
		},
		Meta: ResponseMeta{
			Message: msg,
			Code:    http.StatusOK,
		},
		Page: Page{
			Limit:       page.Limit,
			TotalPage:   page.TotalPage,
			TotalData:   page.TotalData,
			CurrentPage: page.CurrentPage,
		},
	}
	Json(response, http.StatusOK, w)
}

func SuccessData(msg, dataName string, data any, w http.ResponseWriter) {
	response := Response{
		Data: map[string]any{
			dataName: data,
		},
		Meta: ResponseMeta{
			Message: msg,
			Code:    http.StatusOK,
		},
	}

	Json(response, http.StatusOK, w)
}
