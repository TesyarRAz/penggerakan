package course_converter

import (
	"encoding/json"

	course_entity "github.com/TesyarRAz/penggerak/internal/app/course/entity"
	course_model "github.com/TesyarRAz/penggerak/internal/app/course/model"
)

func SubModuleToResponse(subModule *course_entity.SubModule) *course_model.SubModuleResponse {
	return &course_model.SubModuleResponse{
		ID:        subModule.ID,
		ModuleId:  subModule.ModuleID,
		Name:      subModule.Name,
		Structure: json.RawMessage(subModule.Structure),
		CreatedAt: subModule.CreatedAt,
	}
}
