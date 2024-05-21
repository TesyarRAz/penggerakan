package shared_model

import (
	"strings"

	"github.com/TesyarRAz/penggerak/internal/pkg/errors"
)

type VerifyUserRequest struct {
	AccessToken string `validate:"required"`
}

func NewVerifyUserRequestFromAuthorizationHeader(header string) (*VerifyUserRequest, error) {
	if header == "" {
		return nil, errors.NewUnauthorized()
	}

	splitToken := strings.Split(header, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return nil, errors.NewUnauthorized()
	}

	parsedToken := splitToken[1]
	return &VerifyUserRequest{
		AccessToken: parsedToken,
	}, nil
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required" name:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken    string `json:"access_token"`
	AccessTokenExp int64  `json:"access_token_exp"`
	RefreshToken   string `json:"refresh_token"`
}
