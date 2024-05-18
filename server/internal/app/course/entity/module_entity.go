package course_entity

import "time"

type Module struct {
	ID        int        `json:"id"`
	CourseID  int        `json:"course_id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	Course *Course `json:"course,omitempty"`
}
