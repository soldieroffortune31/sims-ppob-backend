package web

type WebResponses struct {
	WebResponse
	Paging interface{} `json:"paging"`
}
