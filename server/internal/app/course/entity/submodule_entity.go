package course_entity

import (
	"time"

	"github.com/jmoiron/sqlx/types"
)

type SubModule struct {
	ID        string         `db:"id"`
	ModuleID  string         `db:"module_id"`
	Name      string         `db:"name"`
	Structure types.JSONText `db:"structure"`
	CreatedAt *time.Time     `db:"created_at"`
	UpdatedAt *time.Time     `db:"updated_at"`
}
