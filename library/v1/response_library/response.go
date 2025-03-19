package response_library

import (
	"fmt"
	"go-learn-news-portal/library/v1/constant"
	"go-learn-news-portal/library/v1/handling"
	"net/http"
)

func Error(err error, w http.ResponseWriter) {
	if err == nil {
		return
	}

	clientErr, ok := err.(handling.ClientError)
	if !ok {
		fmt.Printf("an error occured: %v\n", err)
		res := Response{
			Meta: ResponseMeta{
				Message:   "Internal server error",
				Status:    false,
				Code:      http.StatusInternalServerError,
				ErrorCode: constant.ERR_INTERNAL_SERVER_ERROR,
			},
		}
		JSON(res, http.StatusInternalServerError, w)
		return
	}

	res := Response{
		Meta: ResponseMeta{
			Message:   clientErr.GetMessage(),
			Code:      clientErr.GetStatus(),
			ErrorCode: clientErr.GetErrorCode(),
		},
	}
	JSON(res, clientErr.GetStatus(), w)
}

func Success(msg string, w http.ResponseWriter) {
	response := Response{
		Data: nil,
		Meta: ResponseMeta{
			Message: msg,
			Status:  true,
			//Code:    http.StatusOK,
		},
	}
	JSON(response, http.StatusOK, w)
}

func SuccessCreate(msg, dataName string, data interface{}, w http.ResponseWriter) {
	response := Response{
		Data: map[string]interface{}{
			dataName: data,
		},
		Meta: ResponseMeta{
			Message: msg,
			Code:    http.StatusCreated,
		},
	}
	JSON(response, http.StatusCreated, w)
}

func SuccessDataPaginate(msg, dataName string, data interface{}, page *Page, w http.ResponseWriter) {
	response := PaginationResponse{
		Data: map[string]interface{}{
			dataName: data,
		},
		Meta: ResponseMeta{
			Message: msg,
			Status:  true,
			//Code:    http.StatusOK,
		},
		Page: Page{
			Limit:       page.Limit,
			TotalPage:   page.TotalPage,
			TotalData:   page.TotalData,
			CurrentPage: page.CurrentPage,
		},
	}
	JSON(response, http.StatusOK, w)
}

func SuccessDataIdentifier(msg, dataName string, data interface{}, w http.ResponseWriter) {
	response := Response{
		Data: map[string]interface{}{
			dataName: data,
		},
		Meta: ResponseMeta{
			Message: msg,
			Status:  true,
			//Code:    http.StatusOK,
		},
	}

	JSON(response, http.StatusOK, w)
}

func SuccessData(msg string, data interface{}, w http.ResponseWriter) {
	response := Response{
		Data: data,
		Meta: ResponseMeta{
			Message: msg,
			Status:  true,
			//Code:    http.StatusOK,
		},
	}

	JSON(response, http.StatusOK, w)
}

func SuccessLogin(msg, accessToken, refreshToken string, w http.ResponseWriter) {
	response := Response{
		Data: map[string]interface{}{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
		Meta: ResponseMeta{
			Message: msg,
			//Code:    http.StatusOK,
		},
	}

	JSON(response, http.StatusOK, w)
}
