package user_converter

import (
	user_entity "github.com/TesyarRAz/penggerak/internal/app/user/entity"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"

	lop "github.com/samber/lo/parallel"
)

func UserToResponse(user *user_entity.User, isDetailed bool) *model.UserResponse {
	if !isDetailed {
		return &model.UserResponse{
			Email: user.Email,
			Name:  user.Name,
		}
	}

	return &model.UserResponse{
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		RoleResponses: lop.Map(user.Roles, func(role *user_entity.Role, _ int) *model.RoleResponse {
			return RoleToResponse(role)
		}),
	}
}

func RoleToResponse(role *user_entity.Role) *model.RoleResponse {
	return &model.RoleResponse{
		ID:   role.ID,
		Name: role.Name,
		PermissionResponses: lop.Map(role.Permissions, func(permission *user_entity.Permission, _ int) *model.PermissionResponse {
			return PermissionToResponse(permission)
		}),
	}
}

func PermissionToResponse(permission *user_entity.Permission) *model.PermissionResponse {
	return &model.PermissionResponse{
		ID:   permission.ID,
		Name: permission.Name,
	}
}
