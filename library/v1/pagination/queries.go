package pagination

import (
	"net/http"
	"strconv"
	"strings"
)

type StatusForMe string

const StatusForMeDone StatusForMe = "done"

type Queries struct {
	// OrderType can be asc or desc
	OrderType  string
	OrderBy    string
	Search     string
	Status     string
	CategoryID int64

	// Limit Page can get data 25, 50, 100, all like as laravel filament concept
	Limit int
	Page  int

	// userID and below are legacy for learning purposes
	//`json:"status_for_me" validate:"optional,enum=need_action|ongoing|done"`
	//`json:"pengguna_jasa_id" validate:"optional,uuid"`
	userID          uint
	userInstituteID string
	unitIDs         []string
	statusForMe     StatusForMe
	penggunaJasaID  string
}

func NewQueriesNetHTTP(r *http.Request) *Queries {
	return &Queries{
		OrderBy:    getQueryWithDefault(r, "order_by", "id"),
		OrderType:  getOrderType(r),
		Search:     getQuery(r, "search"),
		Status:     getQuery(r, "status"),
		CategoryID: getInt64Query(r, "category_id"),
	}
}

// getQuery using strings.TrimSpace does not affect keyword searches
// because it only removes prefix and suffix spaces, not spaces in the middle.
//
// getQuery also generic function for all next functions
func getQuery(r *http.Request, key string) string {
	return strings.TrimSpace(r.URL.Query().Get(key))
}

func getInt64Query(r *http.Request, key string) int64 {
	value := getQuery(r, key)
	if value == "" {
		return 0
	}

	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}

	return intValue
}

func getQueryWithDefault(r *http.Request, key, defaultValue string) string {
	val := getQuery(r, key)
	if val == "" {
		return defaultValue
	}
	return val
}

func getOrderType(r *http.Request) string {
	val := strings.ToLower(getQuery(r, "order_type"))
	if val != "asc" && val != "desc" {
		return "desc"
	}
	return val
}
