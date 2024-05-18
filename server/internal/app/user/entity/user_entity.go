package user_entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID           string         `db:"id"`
	Name         string         `db:"name"`
	Email        string         `db:"email"`
	Password     string         `db:"password"`
	ProfileImage sql.NullString `db:"profile_image"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    *time.Time     `db:"updated_at"`

	Roles       []*Role
	Permissions []*Permission
}

func (u *User) HasRoles(roleNames ...string) bool {
	for _, r := range u.Roles {
		for _, roleName := range roleNames {
			if r.Name == roleName {
				return true
			}
		}
	}

	return false
}

func (u *User) HasPermission(permissionNames ...string) bool {
	permissions := u.AllPermission()
	for _, p := range permissions {
		for _, permissionName := range permissionNames {
			if p.Name == permissionName {
				return true
			}
		}
	}

	return false
}

func (u *User) AllPermission() []*Permission {
	permissions := u.Permissions

	for _, r := range u.Roles {
		permissions = append(permissions, r.Permissions...)
	}

	return permissions
}
