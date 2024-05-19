package user_converter

import (
	user_entity "github.com/TesyarRAz/penggerak/internal/app/user/entity"
	user_model "github.com/TesyarRAz/penggerak/internal/app/user/model"

	lop "github.com/samber/lo/parallel"
)

func UserToResponse(user *user_entity.User, isDetailed bool) *user_model.UserResponse {
	if !isDetailed {
		return &user_model.UserResponse{
			ID:                  user.ID,
			Email:               user.Email,
			Name:                user.Name,
			ProfileImage:        user.ProfileImage.String,
			RoleResponses:       []*user_model.RoleResponse{},
			PermissionResponses: []*user_model.PermissionResponse{},
		}
	}

	return &user_model.UserResponse{
		ID:           user.ID,
		Email:        user.Email,
		Name:         user.Name,
		ProfileImage: user.ProfileImage.String,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		RoleResponses: lop.Map(user.Roles, func(role *user_entity.Role, _ int) *user_model.RoleResponse {
			return RoleToResponse(role)
		}),
		PermissionResponses: lop.Map(user.Permissions, func(permission *user_entity.Permission, _ int) *user_model.PermissionResponse {
			return PermissionToResponse(permission)
		}),
	}
}

func RoleToResponse(role *user_entity.Role) *user_model.RoleResponse {
	return &user_model.RoleResponse{
		ID:   role.ID,
		Name: role.Name,
		PermissionResponses: lop.Map(role.Permissions, func(permission *user_entity.Permission, _ int) *user_model.PermissionResponse {
			return PermissionToResponse(permission)
		}),
	}
}

func PermissionToResponse(permission *user_entity.Permission) *user_model.PermissionResponse {
	return &user_model.PermissionResponse{
		ID:   permission.ID,
		Name: permission.Name,
	}
}
