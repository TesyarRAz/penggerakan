package user_entity

import (
	"database/sql"
	"time"
)

type Teacher struct {
	ID           string         `db:"id"`
	UserID       string         `db:"user_id"`
	Name         string         `db:"name"`
	ProfileImage sql.NullString `db:"profile_image"`
	CreatedAt    *time.Time     `db:"created_at"`
	UpdatedAt    *time.Time     `db:"updated_at"`

	User *User
}
