package course_entity

import "time"

type Course struct {
	ID        string     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Image     string     `json:"image" db:"image"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
