package user_model

import (
	"time"
)

type TeacherResponse struct {
	ID           string     `json:"id"`
	UserID       string     `json:"user_id"`
	Name         string     `json:"name"`
	ProfileImage string     `json:"profile_image"`
	CreatedAt    *time.Time `json:"created_at"`
}

type ParamTeacherRequest struct {
	ID string `params:"id" validate:"required" name:"id"`
}

type CreateTeacherRequest struct {
	UserID       string `json:"user_id" validate:"required" name:"user_id"`
	Name         string `json:"name" validate:"required" name:"name"`
	ProfileImage string `json:"profile_image" name:"profile_image"`
}

type UpdateTeacherRequest struct {
	*ParamTeacherRequest

	Name         string `json:"name" validate:"required" name:"name"`
	ProfileImage string `json:"profile_image" name:"profile_image"`
}

type DeleteTeacherRequest struct {
	*ParamTeacherRequest
}

type FindTeacherRequest struct {
	*ParamTeacherRequest
}

type FindTeacherByUserIdRequest struct {
	UserID string `params:"user_id" validate:"required" name:"user_id"`
}
