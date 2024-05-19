package monolith_service

import (
	"context"

	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	shared_model "github.com/TesyarRAz/penggerak/internal/pkg/model/shared"
	"github.com/TesyarRAz/penggerak/internal/pkg/service"
)

type AuthService struct {
	AuthProvider service.AuthService
}

var _ service.AuthService = &AuthService{}

// Verify implements service.AuthService.
func (a *AuthService) Verify(context.Context, *shared_model.VerifyUserRequest) (*model.Auth, error) {
	return a.AuthProvider.Verify(context.Background(), &shared_model.VerifyUserRequest{})
}
