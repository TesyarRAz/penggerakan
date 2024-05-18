package user_provider

import (
	"context"

	user_usecase "github.com/TesyarRAz/penggerak/internal/app/user/usecase"
	"github.com/TesyarRAz/penggerak/internal/pkg/config"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/TesyarRAz/penggerak/internal/pkg/service"
)

type AuthProvider struct {
	UseCase *user_usecase.UserUseCase
}

var _ service.AuthService = &AuthProvider{}
var _ config.ServiceProvider = &AuthProvider{}

func NewAuthProvider(useCase *user_usecase.UserUseCase) *AuthProvider {
	return &AuthProvider{
		UseCase: useCase,
	}
}

func (a *AuthProvider) Boot() {

}

func (a *AuthProvider) Verify(ctx context.Context, request *model.VerifyUserRequest) (*model.Auth, error) {
	return a.UseCase.Verify(ctx, request)
}
