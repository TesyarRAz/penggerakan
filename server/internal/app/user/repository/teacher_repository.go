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

type TeacherRepository struct {
	Log *logrus.Logger
	DB  *sqlx.DB

	repository.Repository
}

func NewTeacherRepository(log *logrus.Logger, db *sqlx.DB) *TeacherRepository {
	return &TeacherRepository{
		Log: log,
		DB:  db,
	}
}

func (t *TeacherRepository) Count(db *sqlx.Tx) (int64, error) {
	var count int64
	err := db.Get(&count, "SELECT COUNT(*) FROM teachers")

	return count, err
}

func (t *TeacherRepository) Create(db *sqlx.Tx, entity *user_entity.Teacher) error {
	entity.ID = util.GenerateUUID().String()
	now := time.Now()
	entity.CreatedAt = &now

	_, err := db.NamedExec("INSERT INTO teachers (id, user_id, name, profile_image, created_at) VALUES (:id, :user_id, :name, :profile_image, :created_at)", entity)

	return err
}

func (t *TeacherRepository) Delete(db *sqlx.Tx, entity *user_entity.Teacher) error {
	_, err := db.Exec("DELETE FROM teachers WHERE id = $1", entity.ID)

	return err
}

func (t *TeacherRepository) FindById(db *sqlx.Tx, entity *user_entity.Teacher, id any) error {
	err := db.Get(entity, "SELECT * FROM teachers WHERE id = $1", id)

	return err
}

func (t *TeacherRepository) FindByUserId(db *sqlx.Tx, entity *user_entity.Teacher, userId any) error {
	err := db.Get(entity, "SELECT * FROM teachers WHERE user_id = $1", userId)

	return err
}

func (t *TeacherRepository) List(db *sqlx.Tx, entities *[]*user_entity.Teacher, request *model.PageRequest) (*model.PageMetadata, error) {
	limit := util.Clamp(request.PerPage, 1, 100)

	result, pageInfo, err := repository.Paginate(&repository.PaginationConfig[user_entity.Teacher]{
		DB:      db,
		Limit:   limit,
		Request: request,
		Table:   "teachers",
		SearchColumn: []string{
			"name",
		},
		FnId: func(teacher *user_entity.Teacher) string {
			return teacher.ID
		},
		FnCreatedAt: func(teacher *user_entity.Teacher) time.Time {
			return *teacher.CreatedAt
		},
	})
	if err != nil {
		t.Log.Warnf("Failed to paginate teacher : %+v", err)
		return nil, err
	}

	(*entities) = result

	return pageInfo, err
}

func (t *TeacherRepository) Update(db *sqlx.Tx, entity *user_entity.Teacher) error {
	now := time.Now()
	entity.UpdatedAt = &now

	_, err := db.NamedExec("UPDATE teachers SET name = :name, profile_image = :profile_image, updated_at = :updated_at WHERE id = :id", entity)

	return err
}

var _ repository.BaseActionRepository[user_entity.Teacher, model.PageRequest] = &TeacherRepository{}
