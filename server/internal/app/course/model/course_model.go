package course_model

import "time"

type CourseResponse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Image     string     `json:"image"`
	CreatedAt *time.Time `json:"created_at"`
}

type ParamCourseRequest struct {
	ID string `params:"id" validate:"required"`
}

type CreateCourseRequest struct {
	Name  string `json:"name" validate:"required,max=100"`
	Image string `json:"image" validate:"required,max=100"`
}

type UpdateCourseRequest struct {
	*ParamCourseRequest
	Name  string `json:"name" validate:"required,max=100"`
	Image string `json:"image" validate:"required,max=100"`
}

type DeleteCourseRequest struct {
	*ParamCourseRequest
}

type FindCourseRequest struct {
	*ParamCourseRequest
}
