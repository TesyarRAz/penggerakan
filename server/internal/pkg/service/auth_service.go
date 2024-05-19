package service

import (
	"context"

	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	shared_model "github.com/TesyarRAz/penggerak/internal/pkg/model/shared"
)

const AUTH_SERVICE = "AuthService"

type AuthService interface {
	Verify(context.Context, *shared_model.VerifyUserRequest) (*model.Auth, error)
}
