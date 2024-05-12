package user_migration

import (
	"context"

	user_entity "github.com/TesyarRAz/penggerak/internal/app/user/entity"
	user_repository "github.com/TesyarRAz/penggerak/internal/app/user/repository"
	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type MigrationConfig struct {
	Dsn       string
	SourceURL string
	Logger    *logrus.Logger
	DB        *sqlx.DB
}

type UserMigration struct {
	config  *MigrationConfig
	migrate *migrate.Migrate
}

func New(config *MigrationConfig) (*UserMigration, error) {
	m, err := migrate.New(config.SourceURL, config.Dsn)
	if err != nil {
		return nil, err
	}

	return &UserMigration{
		config:  config,
		migrate: m,
	}, nil
}

func (u *UserMigration) Up(ctx context.Context, withSeed bool) error {
	if err := u.migrate.Up(); err != nil {
		return err
	}

	if withSeed {
		u.seed(ctx)
	}

	return nil
}

func (u *UserMigration) Down(force int) error {
	return u.migrate.Down()
}

func (u *UserMigration) Drop() error {
	return u.migrate.Drop()
}

func (u *UserMigration) seed(ctx context.Context) {
	userRepository := user_repository.NewUserRepository(u.config.Logger, u.config.DB)
	roleRepository := user_repository.NewPermissionRepository(u.config.Logger, u.config.DB)

	tx, err := u.config.DB.BeginTxx(ctx, nil)
	if err != nil {
		u.config.Logger.Errorf("Error when starting transaction: %v", err)
		return
	}
	defer tx.Rollback()

	admin := &user_entity.User{
		Name:     "Admin",
		Email:    "admin@example.com",
		Password: util.HashPassword("password"),
	}
	if err = userRepository.Create(tx, admin); err != nil {
		u.config.Logger.Errorf("Error when creating user: %v", err)
		return
	}

	roleAdmin := user_entity.Role{
		Name: "admin",
	}
	if err = roleRepository.CreateRole(tx, &roleAdmin); err != nil {
		u.config.Logger.Errorf("Error when creating role: %v", err)
		return
	}

	permissionAdminFullAccess := user_entity.Permission{
		Name: "Admin_FullAccess",
	}
	if err = roleRepository.CreatePermission(tx, &permissionAdminFullAccess); err != nil {
		u.config.Logger.Errorf("Error when creating permission: %v", err)
		return
	}

	if err = roleRepository.AttachPermissionToRole(tx, roleAdmin.ID, permissionAdminFullAccess.ID); err != nil {
		u.config.Logger.Errorf("Error when attaching permission to role: %v", err)
		return
	}

	if err = roleRepository.AttachRoleToUser(tx, roleAdmin.ID, admin.ID); err != nil {
		u.config.Logger.Errorf("Error when attaching role to user: %v", err)
		return
	}

	if err = tx.Commit(); err != nil {
		u.config.Logger.Errorf("Error when committing transaction: %v", err)
		return
	}

	u.config.Logger.Info("Seeding data success")
}
