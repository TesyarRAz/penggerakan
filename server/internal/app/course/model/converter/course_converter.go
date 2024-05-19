package course_converter

import (
	course_entity "github.com/TesyarRAz/penggerak/internal/app/course/entity"
	course_model "github.com/TesyarRAz/penggerak/internal/app/course/model"
)

func CourseToResponse(course *course_entity.Course) *course_model.CourseResponse {
	return &course_model.CourseResponse{
		ID:        course.ID,
		TeacherID: course.TeacherID,
		Name:      course.Name,
		Image:     course.Image,
		CreatedAt: course.CreatedAt,
	}
}
