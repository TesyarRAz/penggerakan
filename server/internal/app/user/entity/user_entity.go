package user_entity

import "database/sql"

type User struct {
	ID           string         `db:"id"`
	Name         string         `db:"name"`
	Email        string         `db:"email"`
	Password     string         `db:"password"`
	ProfileImage sql.NullString `db:"profile_image"`
	CreatedAt    int64          `db:"created_at"`
	UpdatedAt    int64          `db:"updated_at"`
}
