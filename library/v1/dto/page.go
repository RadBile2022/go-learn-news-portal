package dto

type Page struct {
	CurrentPage int   `json:"current_page" json:"current_page,omitempty"`
	TotalData   int64 `json:"total_data" json:"total_data,omitempty"`
	Limit       int   `json:"limit" json:"limit,omitempty"`
	TotalPage   int   `json:"total_page" json:"total_page,omitempty"`
}
