package course_repository

import (
	"time"

	course_entity "github.com/TesyarRAz/penggerak/internal/app/course/entity"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/TesyarRAz/penggerak/internal/pkg/repository"
	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type CourseRepository struct {
	Log *logrus.Logger
	DB  *sqlx.DB

	repository.Repository
}

var _ repository.BaseActionRepository[course_entity.Course] = &CourseRepository{}

func NewCourseRepository(log *logrus.Logger, db *sqlx.DB) *CourseRepository {
	return &CourseRepository{
		Log: log,
		DB:  db,
	}
}

func (c *CourseRepository) List(db *sqlx.Tx, entities *[]*course_entity.Course, request *model.PageRequest) (*model.PageMetadata, error) {
	limit := util.Clamp(request.PerPage, 1, 100)

	result, pageInfo, err := repository.Paginate(&repository.PaginationConfig[course_entity.Course]{
		DB:      db,
		Limit:   limit,
		Request: request,
		Table:   "courses",
		FnId: func(course *course_entity.Course) string {
			return course.ID
		},
		FnCreatedAt: func(course *course_entity.Course) time.Time {
			return course.CreatedAt
		},
	})
	if err != nil {
		c.Log.Warnf("Failed to paginate course : %+v", err)
		return nil, err
	}

	(*entities) = result

	return pageInfo, err
}

func (c *CourseRepository) Count(db *sqlx.Tx) (int64, error) {
	panic("unimplemented")
}

func (c *CourseRepository) Create(db *sqlx.Tx, entity *course_entity.Course) error {
	entity.ID = util.GenerateUUID().String()

	_, err := db.NamedExec("INSERT INTO courses (id, name, image) VALUES (:id, :name, :image)", entity)

	return err
}

// Delete implements repository.BaseActionRepository.
func (c *CourseRepository) Delete(db *sqlx.Tx, entity *course_entity.Course) error {
	panic("unimplemented")
}

// FindById implements repository.BaseActionRepository.
func (c *CourseRepository) FindById(db *sqlx.Tx, entity *course_entity.Course, id any) error {
	panic("unimplemented")
}

// Update implements repository.BaseActionRepository.
func (c *CourseRepository) Update(db *sqlx.Tx, entity *course_entity.Course) error {
	panic("unimplemented")
}
