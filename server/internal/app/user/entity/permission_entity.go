package user_entity

type Role struct {
	ID   string `db:"id"`
	Name string `db:"name"`

	Permissions []*Permission
}

type Permission struct {
	ID   string `db:"id"`
	Name string `db:"name"`

	Roles []*Role
}

type RolePermission struct {
	RoleID       string `db:"role_id"`
	PermissionID string `db:"permission_id"`

	RoleName       string `db:"role_name"`
	PermissionName string `db:"permission_name"`
}

type UserRole struct {
	RoleID string `db:"role_id"`
	UserID string `db:"user_id"`

	RoleName string `db:"role_name"`
}

type UserPermission struct {
	UserID       string `db:"user_id"`
	PermissionID string `db:"permission_id"`

	PermissionName string `db:"permission_name"`
}
