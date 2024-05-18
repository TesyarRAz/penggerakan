package course_entity

import "time"

type SubModule struct {
	ID        int        `json:"id"`
	ModuleID  int        `json:"module_id"`
	Name      string     `json:"name"`
	Structure string     `json:"structure"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
