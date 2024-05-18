package model

import (
	"strings"
	"time"

	"github.com/TesyarRAz/penggerak/internal/pkg/errors"
)

type UserResponse struct {
	Email     string     `json:"email"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	RoleResponses       []*RoleResponse       `json:"roles,omitempty"`
	PermissionResponses []*PermissionResponse `json:"permissions,omitempty"`
}

type RoleResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	PermissionResponses []*PermissionResponse `json:"permissions,omitempty"`
}

type PermissionResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

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

type UpdateUserRequest struct {
	ID       string `json:"-" validate:"required,max=100"`
	Password string `json:"password,omitempty" validate:"max=100"`
	Name     string `json:"name,omitempty" validate:"max=100"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type LoginUserResponse struct {
	ID        string     `json:"id,omitempty"`
	Email     string     `json:"email,omitempty"`
	Name      string     `json:"name,omitempty"`
	Token     string     `json:"token,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type LogoutUserRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}

type GetUserRequest struct {
	ID         string `json:"id" validate:"required,max=100"`
	IsDetailed bool   `json:"-"`
}
