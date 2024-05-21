package model

type WebResponse struct {
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

	SearchRequest
}

func (p *PageRequest) GenerateDefault() {
	if p.Order == "" {
		p.Order = "desc"
	}
	if p.PerPage == 0 {
		p.PerPage = 10
	}
}

type PageMetadata struct {
	NextCursor string `json:"next_cursor"`
	PrevCursor string `json:"prev_cursor"`
}

type PageResponse[T any] struct {
	Data         []T          `json:"data"`
	PageMetadata PageMetadata `json:"paging"`
}
