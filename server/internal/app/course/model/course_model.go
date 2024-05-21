package course_model

import (
	"time"

	shared_model "github.com/TesyarRAz/penggerak/internal/pkg/model/shared"
)

type CourseResponse struct {
	ID        string     `json:"id"`
	TeacherID string     `json:"teacher_id"`
	Name      string     `json:"name"`
	Image     string     `json:"image"`
	CreatedAt *time.Time `json:"created_at"`

	Teacher *shared_model.TeacherResponse `json:"teacher"`
}

type ParamCourseRequest struct {
	ID string `params:"id" validate:"required" name:"id"`
}

type CreateCourseRequest struct {
	TeacherID string `json:"teacher_id" validate:"required" name:"teacher_id"`
	Name      string `json:"name" validate:"required,max=100" name:"name"`
	Image     string `json:"image" validate:"required,max=100" name:"image"`
}

type UpdateCourseRequest struct {
	ParamCourseRequest
	Name  string `json:"name" validate:"required,max=100" name:"name"`
	Image string `json:"image" validate:"required,max=100" name:"image"`
}

type DeleteCourseRequest struct {
	ParamCourseRequest
}

type FindCourseRequest struct {
	ParamCourseRequest
}
