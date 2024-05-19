package course_converter

import (
	course_entity "github.com/TesyarRAz/penggerak/internal/app/course/entity"
	course_model "github.com/TesyarRAz/penggerak/internal/app/course/model"
)

func ModuleToResponse(module *course_entity.Module) *course_model.ModuleResponse {
	return &course_model.ModuleResponse{
		ID:        module.ID,
		Name:      module.Name,
		CreatedAt: module.CreatedAt,
	}
}
