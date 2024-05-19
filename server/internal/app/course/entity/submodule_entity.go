package course_entity

import (
	"time"

	"github.com/jmoiron/sqlx/types"
)

type SubModule struct {
	ID        string         `json:"id" db:"id"`
	ModuleID  string         `json:"module_id" db:"module_id"`
	Name      string         `json:"name" db:"name"`
	Structure types.JSONText `json:"structure" db:"structure"`
	CreatedAt *time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at" db:"updated_at"`
}
