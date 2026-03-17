package web

type PageRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type PageResponse struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}
