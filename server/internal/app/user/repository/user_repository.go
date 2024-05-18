package user_repository

import (
	"time"

	user_entity "github.com/TesyarRAz/penggerak/internal/app/user/entity"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/TesyarRAz/penggerak/internal/pkg/repository"
	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	Log *logrus.Logger
	DB  *sqlx.DB

	repository.Repository
}

var _ repository.BaseActionRepository[user_entity.User] = &UserRepository{}

func NewUserRepository(log *logrus.Logger, db *sqlx.DB) *UserRepository {
	return &UserRepository{
		Log: log,
		DB:  db,
	}
}

func (r *UserRepository) List(db *sqlx.Tx, entities *[]*user_entity.User, filter *model.PageRequest) (*model.PageMetadata, error) {
	return nil, nil
}

func (r *UserRepository) Create(db *sqlx.Tx, entity *user_entity.User) (err error) {
	entity.ID = util.GenerateUUID().String()
	now := time.Now()
	entity.CreatedAt = &now

	_, err = db.NamedExec("INSERT INTO users (id, name, email, password, created_at) VALUES (:id, :name, :email, :password, :created_at)", entity)

	return
}

func (r *UserRepository) Delete(db *sqlx.Tx, entity *user_entity.User) (err error) {
	_, err = db.Exec("DELETE FROM users WHERE id = $1", entity.ID)

	return
}

func (r *UserRepository) FindByToken(user *user_entity.User, token any) error {
	return nil
}

func (r *UserRepository) FindById(db *sqlx.Tx, entity *user_entity.User, id any) (err error) {
	err = db.Get(entity, "SELECT * FROM users WHERE id = $1", id)

	return
}

func (r *UserRepository) FindByEmail(db *sqlx.Tx, entity *user_entity.User, email any) (err error) {
	err = db.Get(entity, "SELECT * FROM users WHERE email = $1", email)

	return
}

func (r *UserRepository) Update(db *sqlx.Tx, entity *user_entity.User) (err error) {
	now := time.Now()
	entity.UpdatedAt = &now

	_, err = db.NamedExec("UPDATE users SET id = :id, name = :name, email = :email, password = :email, updated_at = :updated_at", entity)

	return
}

func (r *UserRepository) Count(db *sqlx.Tx) (int64, error) {
	var count int64
	err := db.Get(&count, "SELECT COUNT(*) FROM users")

	return count, err
}
