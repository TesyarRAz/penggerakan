package course_entity

import "time"

type ContentPage struct {
	ID        string     `db:"id"`
	CourseID  string     `db:"course_id"`
	Title     string     `db:"title"`
	Content   string     `db:"content"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}
