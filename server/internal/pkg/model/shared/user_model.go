package shared_model

import (
	"strings"

	"github.com/TesyarRAz/penggerak/internal/pkg/errors"
)

type VerifyUserRequest struct {
	Token string `validate:"required"`
}

func NewVerifyUserRequestFromAuthorizationHeader(header string) (*VerifyUserRequest, error) {
	if header == "" {
		return nil, errors.Unauthorized{}
	}

	splitToken := strings.Split(header, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return nil, errors.Unauthorized{}
	}

	parsedToken := splitToken[1]
	return &VerifyUserRequest{
		Token: parsedToken,
	}, nil
}
