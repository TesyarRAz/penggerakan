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

type SubModuleRepository struct {
	Log *logrus.Logger
	DB  *sqlx.DB

	repository.Repository
}

func NewSubModuleRepository(log *logrus.Logger, db *sqlx.DB) *SubModuleRepository {
	return &SubModuleRepository{
		Log: log,
		DB:  db,
	}
}

func (m *SubModuleRepository) Count(db *sqlx.Tx) (int64, error) {
	var count int64
	err := db.Get(&count, "SELECT COUNT(*) FROM submodules")

	return count, err
}

func (m *SubModuleRepository) Create(db *sqlx.Tx, entity *course_entity.SubModule) error {
	entity.ID = util.GenerateUUID().String()
	now := time.Now()
	entity.CreatedAt = &now

	_, err := db.NamedExec("INSERT INTO submodules (id, module_id, name, structure, created_at) VALUES (:id, :module_id, :name, :structure, :created_at)", entity)

	return err
}

func (m *SubModuleRepository) Delete(db *sqlx.Tx, entity *course_entity.SubModule) error {
	_, err := db.Exec("DELETE FROM submodules WHERE id = $1", entity.ID)

	return err
}

func (m *SubModuleRepository) FindById(db *sqlx.Tx, entity *course_entity.SubModule, id any) error {
	err := db.Get(entity, "SELECT * FROM submodules WHERE id = $1", id)

	return err
}

func (m *SubModuleRepository) List(db *sqlx.Tx, entities *[]*course_entity.SubModule, request *course_model.ListSubModuleRequest) (*model.PageMetadata, error) {
	limit := util.Clamp(request.PerPage, 1, 100)

	result, pageInfo, err := repository.Paginate(&repository.PaginationConfig[course_entity.SubModule]{
		DB:      db,
		Limit:   limit,
		Request: &request.PageRequest,
		Table:   "submodules",
		FnWhereBuilder: func(namedVar *map[string]interface{}) string {
			(*namedVar)["module_id"] = request.ModuleID
			return "module_id = :module_id"
		},
		FnId: func(module *course_entity.SubModule) string {
			return module.ID
		},
		FnCreatedAt: func(module *course_entity.SubModule) time.Time {
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

func (m *SubModuleRepository) Update(db *sqlx.Tx, entity *course_entity.SubModule) error {
	now := time.Now()
	entity.UpdatedAt = &now

	_, err := db.NamedExec("UPDATE submodules SET name = :name, structure = :structure, updated_at = :updated_at WHERE id = :id", entity)

	return err
}
