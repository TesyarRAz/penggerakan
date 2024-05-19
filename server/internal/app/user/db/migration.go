package user_migration

import (
	"context"

	user_entity "github.com/TesyarRAz/penggerak/internal/app/user/entity"
	user_repository "github.com/TesyarRAz/penggerak/internal/app/user/repository"
	migration "github.com/TesyarRAz/penggerak/internal/pkg/db"
	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
)

func New(config *migration.MigrationConfig) (*migration.Migration, error) {
	return migration.New(config, &pgx.Config{
		MigrationsTable: "user_schema_migrations",
	}, seed)
}

func seed(ctx context.Context, config *migration.MigrationConfig) error {
	userRepository := user_repository.NewUserRepository(config.Logger, config.DB)
	roleRepository := user_repository.NewPermissionRepository(config.Logger, config.DB)

	tx, err := config.DB.BeginTxx(ctx, nil)
	if err != nil {
		config.Logger.Errorf("Error when starting transaction: %v", err)
		return err
	}
	defer tx.Rollback()

	admin := &user_entity.User{
		Name:     "Admin",
		Email:    "admin@example.com",
		Password: util.HashPassword("password"),
	}
	if err = userRepository.Create(tx, admin); err != nil {
		config.Logger.Errorf("Error when creating user: %v", err)
		return err
	}

	roleAdmin := user_entity.Role{
		Name: "admin",
	}
	if err = roleRepository.CreateRole(tx, &roleAdmin); err != nil {
		config.Logger.Errorf("Error when creating role: %v", err)
		return err
	}

	permissionAdminFullAccess := user_entity.Permission{
		Name: "Admin_FullAccess",
	}
	if err = roleRepository.CreatePermission(tx, &permissionAdminFullAccess); err != nil {
		config.Logger.Errorf("Error when creating permission: %v", err)
		return err
	}

	if err = roleRepository.AttachPermissionToRole(tx, roleAdmin.ID, permissionAdminFullAccess.ID); err != nil {
		config.Logger.Errorf("Error when attaching permission to role: %v", err)
		return err
	}

	if err = roleRepository.AttachRoleToUser(tx, roleAdmin.ID, admin.ID); err != nil {
		config.Logger.Errorf("Error when attaching role to user: %v", err)
		return err
	}

	teacher := &user_entity.User{
		Name:     "Teacher",
		Email:    "teacher@example.com",
		Password: util.HashPassword("password"),
	}
	if err = userRepository.Create(tx, teacher); err != nil {
		config.Logger.Errorf("Error when creating user: %v", err)
		return err
	}

	roleTeacher := user_entity.Role{
		Name: "teacher",
	}
	if err = roleRepository.CreateRole(tx, &roleTeacher); err != nil {
		config.Logger.Errorf("Error when creating role: %v", err)
		return err
	}

	permissionTeacherFullAccess := user_entity.Permission{
		Name: "Teacher_FullAccess",
	}
	if err = roleRepository.CreatePermission(tx, &permissionTeacherFullAccess); err != nil {
		config.Logger.Errorf("Error when creating permission: %v", err)
		return err
	}

	if err = roleRepository.AttachPermissionToRole(tx, roleTeacher.ID, permissionTeacherFullAccess.ID); err != nil {
		config.Logger.Errorf("Error when attaching permission to role: %v", err)
		return err
	}

	if err = roleRepository.AttachRoleToUser(tx, roleTeacher.ID, teacher.ID); err != nil {
		config.Logger.Errorf("Error when attaching role to user: %v", err)
		return err
	}

	student := &user_entity.User{
		Name:     "Student",
		Email:    "student@example.com",
		Password: util.HashPassword("password"),
	}

	if err = userRepository.Create(tx, student); err != nil {
		config.Logger.Errorf("Error when creating user: %v", err)
		return err
	}

	roleStudent := user_entity.Role{
		Name: "student",
	}

	if err = roleRepository.CreateRole(tx, &roleStudent); err != nil {
		config.Logger.Errorf("Error when creating role: %v", err)
		return err
	}

	permissionStudentFullAccess := user_entity.Permission{
		Name: "Student_FullAccess",
	}

	if err = roleRepository.CreatePermission(tx, &permissionStudentFullAccess); err != nil {
		config.Logger.Errorf("Error when creating permission: %v", err)
		return err
	}

	if err = roleRepository.AttachPermissionToRole(tx, roleStudent.ID, permissionStudentFullAccess.ID); err != nil {
		config.Logger.Errorf("Error when attaching permission to role: %v", err)
		return err
	}

	if err = roleRepository.AttachRoleToUser(tx, roleStudent.ID, student.ID); err != nil {
		config.Logger.Errorf("Error when attaching role to user: %v", err)
		return err
	}

	if err = tx.Commit(); err != nil {
		config.Logger.Errorf("Error when committing transaction: %v", err)
		return err
	}

	config.Logger.Info("Seeding data success")

	return nil
}
