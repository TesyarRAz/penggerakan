package user_entity

type Role struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`

	Permissions []*Permission `json:"permissions"`
}

type Permission struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`

	Roles []*Role `json:"roles"`
}

type RolePermission struct {
	RoleID       string `json:"role_id" db:"role_id"`
	PermissionID string `json:"permission_id" db:"permission_id"`

	RoleName       string `db:"role_name"`
	PermissionName string `db:"permission_name"`
}

type UserRole struct {
	RoleID string `json:"role_id" db:"role_id"`
	UserID string `json:"user_id" db:"user_id"`

	RoleName string `db:"role_name"`
}

type UserPermission struct {
	UserID       string `json:"user_id" db:"user_id"`
	PermissionID string `json:"permission_id" db:"permission_id"`

	PermissionName string `db:"permission_name"`
}
