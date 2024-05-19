package user_converter

import (
	user_entity "github.com/TesyarRAz/penggerak/internal/app/user/entity"
	user_model "github.com/TesyarRAz/penggerak/internal/app/user/model"
)

func TeacherToResponse(teacher *user_entity.Teacher) *user_model.TeacherResponse {
	return &user_model.TeacherResponse{
		ID:           teacher.ID,
		UserID:       teacher.UserID,
		Name:         teacher.Name,
		ProfileImage: teacher.ProfileImage.String,
		CreatedAt:    teacher.CreatedAt,
	}
}
