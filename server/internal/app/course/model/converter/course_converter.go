package course_converter

import (
	course_entity "github.com/TesyarRAz/penggerak/internal/app/course/entity"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
)

func CourseToResponse(course *course_entity.Course) *model.CourseResponse {
	return &model.CourseResponse{
		ID:        course.ID,
		Name:      course.Name,
		Image:     course.Image,
		CreatedAt: course.CreatedAt,
	}
}
