package service

import (
	"context"

	"github.com/TesyarRAz/penggerak/internal/pkg/model"
)

type AuthService interface {
	Verify(context.Context, *model.VerifyUserRequest) (*model.Auth, error)
}
