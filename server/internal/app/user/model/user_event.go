package user_model

import (
	"time"

	"github.com/TesyarRAz/penggerak/internal/pkg/model"
)

type UserEvent struct {
	ID        string     `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

var _ model.Event = &UserEvent{}

// GetId implements Event.
func (u *UserEvent) GetId() string {
	return u.ID
}
