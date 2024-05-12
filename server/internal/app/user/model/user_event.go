package user_model

import "github.com/TesyarRAz/penggerak/internal/pkg/model"

type UserEvent struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

var _ model.Event = &UserEvent{}

// GetId implements Event.
func (u *UserEvent) GetId() string {
	return u.ID
}
