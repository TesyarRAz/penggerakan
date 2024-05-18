package model

import "time"

type CourseResponse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Image     string     `json:"image"`
	CreatedAt *time.Time `json:"created_at"`
}

type CreateCourseRequest struct {
	Name  string `json:"name" validate:"required,max=100"`
	Image string `json:"image" validate:"required,max=100"`
}

type UpdateCourseRequest struct {
	ID    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required,max=100"`
	Image string `json:"image" validate:"required,max=100"`
}

type DeleteCourseRequest struct {
	ID string `json:"id" validate:"required"`
}

type FindCourseRequest struct {
	ID string `json:"id" validate:"required"`
}
