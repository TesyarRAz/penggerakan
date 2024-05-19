package course_entity

import "time"

type ContentVideo struct {
	ID          string     `db:"id"`
	CourseID    string     `db:"course_id"`
	VideoURL    string     `db:"video_url"`
	Title       string     `db:"title"`
	Description string     `db:"description"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}
