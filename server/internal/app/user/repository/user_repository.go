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

func NewUserRepository(log *logrus.Logger, db *sqlx.DB) *UserRepository {
	return &UserRepository{
		Log: log,
		DB:  db,
	}
}

func (r *UserRepository) List(db *sqlx.Tx, entities *[]*user_entity.User, request *model.PageRequest) (*model.PageMetadata, error) {
	limit := util.Clamp(request.PerPage, 1, 100)

	result, pageInfo, err := repository.Paginate(&repository.PaginationConfig[user_entity.User]{
		DB:      db,
		Limit:   limit,
		Request: request,
		Table:   "users",
		SearchColumn: []string{
			"name", "email",
		},
		FnId: func(user *user_entity.User) string {
			return user.ID
		},
		FnCreatedAt: func(user *user_entity.User) time.Time {
			return *user.CreatedAt
		},
	})
	if err != nil {
		r.Log.Warnf("Failed to paginate user : %+v", err)
		return nil, err
	}

	(*entities) = result

	return pageInfo, err
}

func (r *UserRepository) Create(db *sqlx.Tx, entity *user_entity.User) (err error) {
	entity.ID = util.GenerateUUID().String()
	now := time.Now()
	entity.CreatedAt = &now

	_, err = db.NamedExec("INSERT INTO users (id, name, email, password, profile_image, created_at) VALUES (:id, :name, :email, :password, :profile_image, :created_at)", entity)

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

	_, err = db.NamedExec("UPDATE users SET name = :name, email = :email, password = :email, profile_image = :profile_image, updated_at = :updated_at WHERE id = :id", entity)

	return
}

func (r *UserRepository) Count(db *sqlx.Tx) (int64, error) {
	var count int64
	err := db.Get(&count, "SELECT COUNT(*) FROM users")

	return count, err
}

var _ repository.BaseActionRepository[user_entity.User, model.PageRequest] = &UserRepository{}
