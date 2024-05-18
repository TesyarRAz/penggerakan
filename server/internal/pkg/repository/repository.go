package repository

import (
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	DB *sqlx.DB
}

type BaseActionRepository[T any] interface {
	List(db *sqlx.Tx, entities *[]*T, request *model.PageRequest) (*model.PageMetadata, error)
	Create(db *sqlx.Tx, entity *T) error
	Update(db *sqlx.Tx, entity *T) error
	Delete(db *sqlx.Tx, entity *T) error
	FindById(db *sqlx.Tx, entity *T, id any) error
	Count(db *sqlx.Tx) (int64, error)
}
