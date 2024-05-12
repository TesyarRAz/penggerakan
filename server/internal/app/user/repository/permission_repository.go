package user_repository

import (
	user_entity "github.com/TesyarRAz/penggerak/internal/app/user/entity"
	"github.com/TesyarRAz/penggerak/internal/pkg/repository"
	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

func (*PermissionRepository) AttachPermissionToRole(db *sqlx.Tx, roleID, permissionID string) (err error) {
	_, err = db.Exec("INSERT INTO role_permission (role_id, permission_id) VALUES ($1, $2)", roleID, permissionID)

	return
}

func (*PermissionRepository) DetachPermissionFromRole(db *sqlx.Tx, roleID, permissionID string) (err error) {
	_, err = db.Exec("DELETE FROM role_permission WHERE role_id = $1 AND permission_id = $2", roleID, permissionID)

	return
}

func (*PermissionRepository) AttachRoleToUser(db *sqlx.Tx, roleID, userID string) (err error) {
	_, err = db.Exec("INSERT INTO user_role (role_id, user_id) VALUES ($1, $2)", roleID, userID)

	return
}

func (*PermissionRepository) DetachRoleFromUser(db *sqlx.Tx, roleID, userID string) (err error) {
	_, err = db.Exec("DELETE FROM user_role WHERE role_id = $1 AND user_id = $2", roleID, userID)

	return
}

func (*PermissionRepository) AttachPermissionToUser(db *sqlx.Tx, permissionID, userID string) (err error) {
	_, err = db.Exec("INSERT INTO user_permission (permission_id, user_id) VALUES ($1, $2)", permissionID, userID)

	return
}

func (*PermissionRepository) DetachPermissionFromUser(db *sqlx.Tx, permissionID, userID string) (err error) {
	_, err = db.Exec("DELETE FROM user_permission WHERE permission_id = $1 AND user_id = $2", permissionID, userID)

	return
}

func (*PermissionRepository) PermissionsByRole(db *sqlx.Tx, roleID string) (permissions []user_entity.Permission, err error) {
	err = db.Select(&permissions, "SELECT p.* FROM permissions p JOIN role_permission rp ON p.id = rp.permission_id WHERE rp.role_id = $1", roleID)

	return
}

func (*PermissionRepository) PermissionsByUser(db *sqlx.Tx, userID string) (permissions []user_entity.Permission, err error) {
	err = db.Select(&permissions, "SELECT p.* FROM permissions p JOIN user_permission up ON p.id = up.permission_id WHERE up.user_id = $1", userID)

	return
}

func (*PermissionRepository) RolesByUser(db *sqlx.Tx, userID string) (roles []user_entity.Role, err error) {
	err = db.Select(&roles, "SELECT r.* FROM roles r JOIN user_role ur ON r.id = ur.role_id WHERE ur.user_id = $1", userID)

	return
}

func (*PermissionRepository) UsersByRole(db *sqlx.Tx, roleID string) (users []user_entity.User, err error) {
	err = db.Select(&users, "SELECT u.* FROM users u JOIN user_role ur ON u.id = ur.user_id WHERE ur.role_id = $1", roleID)

	return
}

func (*PermissionRepository) RolesByPermission(db *sqlx.Tx, permissionID string) (roles []user_entity.Role, err error) {
	err = db.Select(&roles, "SELECT r.* FROM roles r JOIN role_permission rp ON r.id = rp.role_id WHERE rp.permission_id = $1", permissionID)

	return
}
