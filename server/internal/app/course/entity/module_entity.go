package course_entity

import "time"

type Module struct {
	ID        string     `json:"id" db:"id"`
	CourseID  string     `json:"course_id" db:"course_id"`
	Name      string     `json:"name" db:"name"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`

	Course *Course `json:"course,omitempty"`
}
