package course_entity

import "time"

type Module struct {
	ID        string     `db:"id"`
	CourseID  string     `db:"course_id"`
	Name      string     `db:"name"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`

	Course *Course
}
