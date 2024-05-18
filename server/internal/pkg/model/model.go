package model

type WebResponse[T any] struct {
	Data   T             `json:"data"`
	Paging *PageMetadata `json:"paging,omitempty"`
	Errors string        `json:"errors,omitempty"`
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
	Data         []T          `json:"data,omitempty"`
	PageMetadata PageMetadata `json:"paging,omitempty"`
}
