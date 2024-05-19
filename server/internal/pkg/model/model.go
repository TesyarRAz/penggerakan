package model

type WebResponse[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message"`
	Errors  string `json:"errors"`
}

type SearchRequest struct {
	Search string `query:"search"`
}
type PageRequest struct {
	PerPage int    `query:"per_page" validate:"min=1,max=100"`
	Order   string `query:"order" validate:"oneof=asc desc"`
	Cursor  string `query:"cursor"`

	*SearchRequest
}

type PageMetadata struct {
	NextCursor string `json:"next_cursor"`
	PrevCursor string `json:"prev_cursor"`
}

type PageResponse[T any] struct {
	Data         []T          `json:"data"`
	PageMetadata PageMetadata `json:"paging"`
}
