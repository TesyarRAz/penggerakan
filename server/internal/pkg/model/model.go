package model

type WebResponse[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message"`
	Errors  string `json:"errors"`
}

type SearchRequest struct {
	Search string `json:"search"`
}
type PageRequest struct {
	PerPage int    `json:"per_page" validate:"min=1,max=100"`
	Order   string `json:"order" validate:"oneof=asc desc"`
	Cursor  string `json:"cursor"`

	SearchRequest
}

type PageMetadata struct {
	NextCursor string `json:"next_cursor"`
	PrevCursor string `json:"prev_cursor"`
}

type PageResponse[T any] struct {
	Data         []T          `json:"data"`
	PageMetadata PageMetadata `json:"paging"`
}
