package user_entity

type Role struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Permission struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type RolePermission struct {
	RoleID       string `json:"role_id" db:"role_id"`
	PermissionID string `json:"permission_id" db:"permission_id"`
}

type UserRole struct {
	RoleID string `json:"role_id" db:"role_id"`
	UserID string `json:"user_id" db:"user_id"`
}

type UserPermission struct {
	UserID       string `json:"user_id" db:"user_id"`
	PermissionID string `json:"permission_id" db:"permission_id"`
}
