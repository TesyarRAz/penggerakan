package course_entity

import "time"

type UserCourse struct {
	UserID      string     `db:"user_id"`
	CourseID    string     `db:"course_id"`
	IsCompleted bool       `db:"is_completed"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}
