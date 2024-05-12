package user_converter

import (
	user_entity "github.com/TesyarRAz/penggerak/internal/app/user/entity"
	user_model "github.com/TesyarRAz/penggerak/internal/app/user/model"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
)

func UserToResponse(user *user_entity.User, isMe bool) *model.UserResponse {
	if !isMe {
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
	}
}

func UserToEvent(user *user_entity.User) *user_model.UserEvent {
	return &user_model.UserEvent{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
