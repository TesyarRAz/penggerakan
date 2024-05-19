package course_repository

import (
	"time"

	course_entity "github.com/TesyarRAz/penggerak/internal/app/course/entity"
	course_model "github.com/TesyarRAz/penggerak/internal/app/course/model"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/TesyarRAz/penggerak/internal/pkg/repository"
	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ModuleRepository struct {
	Log *logrus.Logger
	DB  *sqlx.DB

	repository.Repository
}

func NewModuleRepository(log *logrus.Logger, db *sqlx.DB) *ModuleRepository {
	return &ModuleRepository{
		Log: log,
		DB:  db,
	}
}

func (m *ModuleRepository) Count(db *sqlx.Tx) (int64, error) {
	var count int64
	err := db.Get(&count, "SELECT COUNT(*) FROM modules")

	return count, err
}

func (m *ModuleRepository) Create(db *sqlx.Tx, entity *course_entity.Module) error {
	entity.ID = util.GenerateUUID().String()
	now := time.Now()
	entity.CreatedAt = &now

	_, err := db.NamedExec("INSERT INTO modules (id, course_id, name, created_at) VALUES (:id, :course_id, :name, :created_at)", entity)

	return err
}

func (m *ModuleRepository) Delete(db *sqlx.Tx, entity *course_entity.Module) error {
	_, err := db.Exec("DELETE FROM modules WHERE id = $1", entity.ID)

	return err
}

func (m *ModuleRepository) FindById(db *sqlx.Tx, entity *course_entity.Module, id any) error {
	err := db.Get(entity, "SELECT * FROM modules WHERE id = $1", id)

	return err
}

func (m *ModuleRepository) List(db *sqlx.Tx, entities *[]*course_entity.Module, request *course_model.ListModuleRequest) (*model.PageMetadata, error) {
	limit := util.Clamp(request.PerPage, 1, 100)

	result, pageInfo, err := repository.Paginate(&repository.PaginationConfig[course_entity.Module]{
		DB:      db,
		Limit:   limit,
		Request: &request.PageRequest,
		Table:   "modules",
		SearchColumn: []string{
			"name",
		},
		FnWhereBuilder: func(namedVar *map[string]interface{}) string {
			(*namedVar)["course_id"] = request.CourseID
			return "course_id = :course_id"
		},
		FnId: func(module *course_entity.Module) string {
			return module.ID
		},
		FnCreatedAt: func(module *course_entity.Module) time.Time {
			return *module.CreatedAt
		},
	})
	if err != nil {
		m.Log.Warnf("Failed to paginate module : %+v", err)
		return nil, err
	}

	(*entities) = result

	return pageInfo, err
}

func (m *ModuleRepository) Update(db *sqlx.Tx, entity *course_entity.Module) error {
	now := time.Now()
	entity.UpdatedAt = &now

	_, err := db.NamedExec("UPDATE modules SET name = :name, updated_at = :updated_at WHERE id = :id", entity)

	return err
}

var _ repository.BaseActionRepository[course_entity.Module, course_model.ListModuleRequest] = &ModuleRepository{}
