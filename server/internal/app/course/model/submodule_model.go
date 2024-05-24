package course_model

import (
	"encoding/json"
	"time"

	"github.com/TesyarRAz/penggerak/internal/pkg/model"
)

type ParamSubModuleRequest struct {
	ModuleID string `params:"module_id" validate:"required,uuid" name:"module_id"`
}

type ListSubModuleRequest struct {
	ParamSubModuleRequest

	model.PageRequest
}

type SubModuleResponse struct {
	ID        string          `json:"id"`
	ModuleId  string          `json:"module_id"`
	Name      string          `json:"name"`
	Structure json.RawMessage `json:"structure"`
	CreatedAt *time.Time      `json:"created_at"`
}

type CreateSubModuleRequest struct {
	ParamSubModuleRequest

	Name      string          `json:"name" validate:"required,max=100" name:"name"`
	Structure json.RawMessage `json:"structure" validate:"required,json" name:"structure"`
}

type UpdateSubModuleRequest struct {
	ParamSubModuleRequest

	ID        string          `json:"id" validate:"required,uuid" name:"id"`
	Name      string          `json:"name" validate:"required,max=100" name:"name"`
	Structure json.RawMessage `json:"structure" validate:"required,json" name:"structure"`
}

type DeleteSubModuleRequest struct {
	ParamSubModuleRequest

	ID string `json:"id" validate:"required,uuid" name:"id"`
}

type FindSubModuleRequest struct {
	ParamSubModuleRequest

	ID string `json:"id" validate:"required,uuid" name:"id"`
}
