package course_model

import (
	"time"

	"github.com/TesyarRAz/penggerak/internal/pkg/model"
)

type ParamModuleRequest struct {
	CourseID string `params:"course_id" validate:"required"`
}

type ListModuleRequest struct {
	*ParamModuleRequest

	*model.PageRequest
}

type ModuleResponse struct {
	*ParamModuleRequest

	ID        string     `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
}

type CreateModuleRequest struct {
	*ParamModuleRequest

	Name string `json:"name" validate:"required,max=100"`
}

type UpdateModuleRequest struct {
	*ParamModuleRequest

	ID   string `params:"id" validate:"required"`
	Name string `json:"name" validate:"required,max=100"`
}

type DeleteModuleRequest struct {
	*ParamModuleRequest

	ID string `json:"id" validate:"required"`
}

type FindModuleRequest struct {
	ParamModuleRequest

	ID string `json:"id" validate:"required"`
}
