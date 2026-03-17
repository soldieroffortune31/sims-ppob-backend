package web

type PagedResponse[T any] struct {
	Data []T          `json:"data"`
	Meta PageResponse `json:"meta"`
}
