package user_model

import (
	"time"
)

type UserResponse struct {
	ID           string     `json:"id"`
	Email        string     `json:"email"`
	Name         string     `json:"name"`
	ProfileImage string     `json:"profile_image"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`

	RoleResponses       []*RoleResponse       `json:"roles"`
	PermissionResponses []*PermissionResponse `json:"permissions"`
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

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,max=100" name:"email"`
	Password string `json:"password" validate:"required,max=100" name:"password"`
}

type LoginUserResponse struct {
	ID        string     `json:"id,omitempty"`
	Email     string     `json:"email,omitempty"`
	Name      string     `json:"name,omitempty"`
	Token     string     `json:"token,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type LogoutUserRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}

type FindUserRequest struct {
	ID         string `json:"id" validate:"required,max=100"`
	IsDetailed bool   `json:"-"`
}

type CreateUserRequest struct {
	Email        string `json:"email" validate:"required,max=100,email" name:"email"`
	Password     string `json:"password" validate:"required,max=100" name:"password"`
	Name         string `json:"name" validate:"required,max=100" name:"name"`
	ProfileImage string `json:"profile_image" validate:"omitempty,http_url" name:"profile_image"`
}

type ParamUserRequest struct {
	ID string `params:"id" validate:"required" name:"id"`
}

type DeleteUserRequest struct {
	ParamUserRequest
}

type UpdateUserRequest struct {
	ParamUserRequest

	Email        string `json:"email,omitempty" validate:"required,max=100,email" name:"email"`
	Password     string `json:"password,omitempty" validate:"max=100" name:"password"`
	Name         string `json:"name,omitempty" validate:"required,max=100" name:"name"`
	ProfileImage string `json:"profile_image" validate:"http_url" name:"profile_image"`
}

type AttachRoleToUserRequest struct {
	ParamUserRequest

	Role string `json:"role" validate:"required,max=100" name:"role"`
}

type DetachRoleFromUserRequest struct {
	ParamUserRequest

	Role string `json:"role" validate:"required,max=100" name:"role"`
}
