package course_entity

import "time"

type Course struct {
	ID        string     `db:"id"`
	TeacherID string     `db:"teacher_id"`
	Name      string     `db:"name"`
	Image     string     `db:"image"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}
