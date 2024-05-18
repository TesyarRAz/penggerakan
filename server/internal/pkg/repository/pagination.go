package repository

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/jmoiron/sqlx"
)

type Cursor map[string]interface{}

type PaginationConfig[T any] struct {
	DB      *sqlx.Tx
	Table   string
	Limit   int
	Request *model.PageRequest

	FnId        func(*T) string
	FnCreatedAt func(*T) time.Time
}

func Paginate[T any](config *PaginationConfig[T]) ([]*T, *model.PageMetadata, error) {
	namedVar := map[string]interface{}{}
	pointsNext := false
	sortOrder := config.Request.Order

	isFirstPage := config.Request.Cursor == ""

	query := fmt.Sprintf("SELECT * FROM %s ", config.Table)

	if config.Request.Cursor != "" {
		decodedCursor, err := decodeCursor(config.Request.Cursor)
		if err != nil {
			return nil, nil, err
		}
		pointsNext = decodedCursor["points_next"] == true

		operator, order := getPaginationOperator(pointsNext, config.Request.Order)
		query = query + fmt.Sprintf("WHERE (created_at %s :created_at OR (created_at = :created_at AND id %s :id)) ", operator, operator)
		namedVar["created_at"] = decodedCursor["created_at"]
		namedVar["id"] = decodedCursor["id"]
		if order != "" {
			sortOrder = order
		}
	}

	// Ini bisa aja di SQL Injection
	query = query + fmt.Sprintf("ORDER BY created_at %s LIMIT %v ", sortOrder, config.Limit+1)
	query, args, err := config.DB.BindNamed(query, namedVar)
	if err != nil {
		return nil, nil, err
	}

	query = config.DB.Rebind(query)

	var entities []*T
	if err := config.DB.Select(&entities, query, args...); err != nil {
		return nil, nil, err
	}

	hasPagination := len(entities) > config.Limit
	if hasPagination {
		entities = entities[:config.Limit]
	}

	if !isFirstPage && !pointsNext {
		entities = util.Reverse(entities)
	}

	pageInfo := model.PageMetadata{}
	nextCur := Cursor{}
	prevCur := Cursor{}
	if hasPagination {
		if isFirstPage {
			nextCur := createCursor(config.FnId(entities[config.Limit-1]), config.FnCreatedAt(entities[config.Limit-1]), true)
			pageInfo = model.PageMetadata{
				NextCursor: encodeCursor(nextCur),
			}
		} else {
			if pointsNext {
				if hasPagination {
					nextCur = createCursor(config.FnId(entities[config.Limit-1]), config.FnCreatedAt(entities[config.Limit-1]), true)
				}
				prevCur = createCursor(config.FnId(entities[0]), config.FnCreatedAt(entities[0]), false)
			} else {
				nextCur = createCursor(config.FnId(entities[config.Limit-1]), config.FnCreatedAt(entities[config.Limit-1]), true)
				if hasPagination {
					prevCur = createCursor(config.FnId(entities[0]), config.FnCreatedAt(entities[0]), false)
				}
			}
			pageInfo = model.PageMetadata{
				NextCursor: encodeCursor(nextCur),
				PrevCursor: encodeCursor(prevCur),
			}
		}
	}

	return entities, &pageInfo, nil
}

func createCursor(id string, createdAt time.Time, pointsNext bool) Cursor {
	return Cursor{
		"id":          id,
		"created_at":  createdAt,
		"points_next": pointsNext,
	}
}

func encodeCursor(cursor Cursor) string {
	if len(cursor) == 0 {
		return ""
	}
	serializedCursor, err := json.Marshal(cursor)
	if err != nil {
		return ""
	}
	encodedCursor := base64.StdEncoding.EncodeToString(serializedCursor)
	return encodedCursor
}

func decodeCursor(cursor string) (Cursor, error) {
	decodedCursor, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return nil, err
	}

	var cur Cursor
	if err := json.Unmarshal(decodedCursor, &cur); err != nil {
		return nil, err
	}
	return cur, nil
}

func getPaginationOperator(pointsNext bool, sortOrder string) (string, string) {
	if pointsNext && sortOrder == "asc" {
		return ">", ""
	}
	if pointsNext && sortOrder == "desc" {
		return "<", ""
	}
	if !pointsNext && sortOrder == "asc" {
		return "<", "desc"
	}
	if !pointsNext && sortOrder == "desc" {
		return ">", "asc"
	}

	return "", ""
}
