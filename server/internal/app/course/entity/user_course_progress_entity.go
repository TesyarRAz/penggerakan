package course_entity

import (
	"time"

	"github.com/jmoiron/sqlx/types"
)

type UserCourseProgress struct {
	UserCourseID string          `db:"user_course_id"`
	ResourceType string          `db:"resource_type"`
	ResourceID   string          `db:"resource_id"`
	IsCompleted  bool            `db:"is_completed"`
	Metadata     *types.JSONText `db:"metadata"`
	CreatedAt    *time.Time      `db:"created_at"`
	UpdatedAt    *time.Time      `db:"updated_at"`
}
