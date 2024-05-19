package user_repository

import (
	user_entity "github.com/TesyarRAz/penggerak/internal/app/user/entity"
	"github.com/TesyarRAz/penggerak/internal/pkg/repository"
	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	lo "github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

type PermissionRepository struct {
	Log *logrus.Logger
	DB  *sqlx.DB

	repository.Repository
}

func NewPermissionRepository(log *logrus.Logger, db *sqlx.DB) *PermissionRepository {
	return &PermissionRepository{
		Log: log,
		DB:  db,
	}
}

func (*PermissionRepository) CreateRole(db *sqlx.Tx, entity *user_entity.Role) (err error) {
	entity.ID = util.GenerateUUID().String()

	_, err = db.NamedExec("INSERT INTO roles (id, name) VALUES (:id, :name)", entity)

	return
}

func (*PermissionRepository) FindRoleByID(db *sqlx.Tx, entity *user_entity.Role, id string) (err error) {
	err = db.Get(entity, "SELECT * FROM roles WHERE id = $1", id)

	return
}

func (*PermissionRepository) FindRoleByName(db *sqlx.Tx, entity *user_entity.Role, name string) (err error) {
	err = db.Get(entity, "SELECT * FROM roles WHERE name = $1", name)

	return
}

func (*PermissionRepository) UserHasRoles(db *sqlx.Tx, userID string, roles ...string) (bool, error) {
	query, args, err := sqlx.In("SELECT COUNT(*) FROM user_role WHERE user_id = ? AND role_id IN (?)", userID, roles)

	if err != nil {
		return false, err
	}

	query = db.Rebind(query)

	var count int
	if err := db.Get(&count, query, args...); err != nil {
		return false, err
	}

	return count > 0, nil
}

func (*PermissionRepository) DeleteRole(db *sqlx.Tx, entity *user_entity.Role) (err error) {
	err = db.Select(entity, "DELETE FROM roles WHERE id = $1", entity.ID)

	return
}

func (*PermissionRepository) CreatePermission(db *sqlx.Tx, entity *user_entity.Permission) (err error) {
	entity.ID = util.GenerateUUID().String()

	_, err = db.NamedExec("INSERT INTO permissions (id, name) VALUES (:id, :name)", entity)

	return
}

func (*PermissionRepository) DeletePermission(db *sqlx.Tx, entity *user_entity.Permission) (err error) {
	err = db.Select(entity, "DELETE FROM permissions WHERE id = $1", entity.ID)

	return
}

func (*PermissionRepository) AttachPermissionToRole(db *sqlx.Tx, roleID string, permissionID string) (err error) {
	_, err = db.Exec("INSERT INTO role_permission (role_id, permission_id) VALUES ($1, $2)", roleID, permissionID)

	return
}

func (*PermissionRepository) DetachPermissionFromRole(db *sqlx.Tx, roleID string, permissionID string) (err error) {
	_, err = db.Exec("DELETE FROM role_permission WHERE role_id = $1 AND permission_id = $2", roleID, permissionID)

	return
}

func (*PermissionRepository) AttachRoleToUser(db *sqlx.Tx, roleID string, userID string) (err error) {
	_, err = db.Exec("INSERT INTO user_role (role_id, user_id) VALUES ($1, $2)", roleID, userID)

	return
}

func (*PermissionRepository) DetachRoleFromUser(db *sqlx.Tx, roleID string, userID string) (err error) {
	_, err = db.Exec("DELETE FROM user_role WHERE role_id = $1 AND user_id = $2", roleID, userID)

	return
}

func (*PermissionRepository) AttachPermissionToUser(db *sqlx.Tx, permissionID string, userID string) (err error) {
	_, err = db.Exec("INSERT INTO user_permission (permission_id, user_id) VALUES ($1, $2)", permissionID, userID)

	return
}

func (*PermissionRepository) DetachPermissionFromUser(db *sqlx.Tx, permissionID string, userID string) (err error) {
	_, err = db.Exec("DELETE FROM user_permission WHERE permission_id = $1 AND user_id = $2", permissionID, userID)

	return
}

func (*PermissionRepository) PermissionsByRoles(db *sqlx.Tx, roles ...*user_entity.Role) error {
	if len(roles) == 0 {
		return nil
	}
	roleIds := lop.Map(roles, func(role *user_entity.Role, _ int) string {
		return role.ID
	})
	roleIds = lo.Uniq(roleIds)
	query, args, err := sqlx.In("SELECT rp.*, p.name as permission_name, r.name as role_name FROM role_permission rp JOIN permissions p ON p.id = rp.permission_id JOIN roles r ON r.id = rp.role_id WHERE rp.role_id IN (?)", roleIds)

	if err != nil {
		return err
	}

	query = db.Rebind(query)

	var rps []*user_entity.RolePermission
	if err := db.Select(&rps, query, args...); err != nil {
		return err
	}

	for _, role := range roles {
		frps := lo.Filter(rps, func(rp *user_entity.RolePermission, _ int) bool {
			return rp.RoleID == role.ID
		})

		role.Permissions = lo.Map(frps, func(rp *user_entity.RolePermission, _ int) *user_entity.Permission {
			return &user_entity.Permission{
				ID:   rp.PermissionID,
				Name: rp.PermissionName,
			}
		})
	}

	return nil
}

func (*PermissionRepository) PermissionsByUsers(db *sqlx.Tx, users ...*user_entity.User) error {
	if len(users) == 0 {
		return nil
	}
	userIds := lop.Map(users, func(user *user_entity.User, _ int) string {
		return user.ID
	})

	userIds = lo.Uniq(userIds)

	query, args, err := sqlx.In("SELECT up.*, p.name as permission_name FROM user_permission up JOIN permissions p ON p.id = up.permission_id WHERE up.user_id IN (?)", userIds)

	if err != nil {
		return err
	}

	query = db.Rebind(query)

	var ups []*user_entity.UserPermission
	if err := db.Select(&ups, query, args...); err != nil {
		return err
	}

	for _, user := range users {
		fups := lo.Filter(ups, func(up *user_entity.UserPermission, _ int) bool {
			return up.UserID == user.ID
		})

		user.Permissions = lo.Map(fups, func(up *user_entity.UserPermission, _ int) *user_entity.Permission {
			return &user_entity.Permission{
				ID:   up.PermissionID,
				Name: up.PermissionName,
			}
		})
	}

	return nil
}

func (*PermissionRepository) RolesByUser(db *sqlx.Tx, users ...*user_entity.User) error {
	if len(users) == 0 {
		return nil
	}

	userIds := lop.Map(users, func(user *user_entity.User, _ int) string {
		return user.ID
	})
	userIds = lo.Uniq(userIds)

	query, args, err := sqlx.In("SELECT ur.*, p.name as role_name FROM user_role ur JOIN roles p ON p.id = ur.role_id WHERE ur.user_id IN (?)", userIds)

	if err != nil {
		return err
	}

	query = db.Rebind(query)

	var urs []*user_entity.UserRole
	if err := db.Select(&urs, query, args...); err != nil {
		return err
	}

	for _, user := range users {
		furs := lo.Filter(urs, func(ur *user_entity.UserRole, _ int) bool {
			return ur.UserID == user.ID
		})

		user.Roles = lo.Map(furs, func(ur *user_entity.UserRole, _ int) *user_entity.Role {
			return &user_entity.Role{
				ID:   ur.RoleID,
				Name: ur.RoleName,
			}
		})
	}

	return nil
}

// func (*PermissionRepository) UsersByRole(db *sqlx.Tx, users *[]*user_entity.User, roleID string) error {
// 	return db.Select(users, "SELECT u.* FROM users u JOIN user_role ur ON u.id = ur.user_id WHERE ur.role_id = $1", roleID)
// }

// func (*PermissionRepository) RolesByPermission(db *sqlx.Tx, roles *[]*user_entity.Role, permissionID string) error {
// 	return db.Select(roles, "SELECT r.* FROM roles r JOIN role_permission rp ON r.id = rp.role_id WHERE rp.permission_id = $1", permissionID)
// }
