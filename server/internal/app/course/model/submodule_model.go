package course_model

import (
	"encoding/json"
	"time"

	"github.com/TesyarRAz/penggerak/internal/pkg/model"
)

type ParamSubModuleRequest struct {
	ModuleID string `params:"module_id" validate:"required"`
}

type ListSubModuleRequest struct {
	*ParamSubModuleRequest

	model.PageRequest
}

type SubModuleResponse struct {
	*ParamSubModuleRequest

	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Structure json.RawMessage `json:"structure"`
	CreatedAt *time.Time      `json:"created_at"`
}

type CreateSubModuleRequest struct {
	*ParamSubModuleRequest

	Name      string          `json:"name" validate:"required,max=100"`
	Structure json.RawMessage `json:"structure" validate:"required,json"`
}

type UpdateSubModuleRequest struct {
	*ParamSubModuleRequest

	ID        string          `json:"id" validate:"required"`
	Name      string          `json:"name" validate:"required,max=100"`
	Structure json.RawMessage `json:"structure" validate:"required,json"`
}

type DeleteSubModuleRequest struct {
	*ParamSubModuleRequest

	ID string `json:"id" validate:"required"`
}

type FindSubModuleRequest struct {
	ParamSubModuleRequest

	ID string `json:"id" validate:"required"`
}
