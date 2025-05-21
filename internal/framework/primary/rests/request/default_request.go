package request

import (
	"errors"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/convert"
	"net/http"
)

type ObjectID struct {
	ID int64
}

func HelperToObjectID(r *http.Request) (*ObjectID, error) {
	id := convert.PathValueIDInt64Chi(r)
	if id == 0 {
		return nil, errors.New("id is required")
	}
	return &ObjectID{
		ID: id,
	}, nil
}
