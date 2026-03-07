package web

type Paging struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalPage int `json:"total_page"`
	Total     int `json:"total"`
}
