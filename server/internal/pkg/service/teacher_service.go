package service

import (
	"context"

	shared_model "github.com/TesyarRAz/penggerak/internal/pkg/model/shared"
)

const TEACHER_SERVICE = "TeacherService"

type TeacherService interface {
	FindByIds(context.Context, ...string) ([]*shared_model.TeacherResponse, error)
}
