package user_provider

import (
	"context"

	user_model "github.com/TesyarRAz/penggerak/internal/app/user/model"
	user_usecase "github.com/TesyarRAz/penggerak/internal/app/user/usecase"
	"github.com/TesyarRAz/penggerak/internal/pkg/config"
	shared_model "github.com/TesyarRAz/penggerak/internal/pkg/model/shared"
	"github.com/TesyarRAz/penggerak/internal/pkg/service"

	lop "github.com/samber/lo/parallel"
)

type TeacherProvider struct {
	UseCase *user_usecase.TeacherUseCase
}

func NewTeacherProvider(useCase *user_usecase.TeacherUseCase) *TeacherProvider {
	return &TeacherProvider{
		UseCase: useCase,
	}
}

func (t *TeacherProvider) FindByIds(ctx context.Context, ids ...string) ([]*shared_model.TeacherResponse, error) {
	teachers, err := t.UseCase.FindByIds(ctx, &user_model.FindTeachersRequest{
		IDs: ids,
	})

	if err != nil {
		return nil, err
	}

	return lop.Map(teachers, func(teacher *user_model.TeacherResponse, _ int) *shared_model.TeacherResponse {
		return &shared_model.TeacherResponse{
			ID:           teacher.ID,
			UserID:       teacher.UserID,
			Name:         teacher.Name,
			ProfileImage: teacher.ProfileImage,
		}
	}), nil
}

func (t *TeacherProvider) Boot() {

}

var _ config.ServiceProvider = &TeacherProvider{}
var _ service.TeacherService = &TeacherProvider{}
