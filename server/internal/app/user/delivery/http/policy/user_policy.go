package user_policy

import "github.com/TesyarRAz/penggerak/internal/pkg/model"

func AllowDetailedUser(auth *model.Auth, userID string) bool {
	if auth == nil {
		return false
	}

	if auth.ID == userID {
		return true
	}

	if auth.HasPermission("Admin_FullAccess") {
		return true
	}

	return false
}

func AllowDeleteUser(auth *model.Auth, userID string) bool {
	if auth == nil {
		return false
	}

	if auth.ID == userID {
		return false
	}

	if auth.HasPermission("Admin_FullAccess") {
		return true
	}

	return false
}

func AllowUpdateUser(auth *model.Auth, userID string) bool {
	if auth == nil {
		return false
	}

	if auth.ID == userID {
		return true
	}

	if auth.HasPermission("Admin_FullAccess") {
		return true
	}

	return false
}
