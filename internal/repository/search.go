package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/polapolo/omdbapp/internal/constant"
	"github.com/polapolo/omdbapp/internal/entity"
	"github.com/polapolo/omdbapp/internal/lib/distributedtracing"
)

// SearchRepository -> SearchRepository object
type SearchRepository struct {
	db *sqlx.DB
}

// NewSearchRepository -> Create new SearchRepository object
func NewSearchRepository(db *sqlx.DB) SearchRepository {
	return SearchRepository{
		db: db,
	}
}

// InsertSearchHistory -> insert 1 row of search history
func (r SearchRepository) InsertSearchHistory(ctx context.Context, row entity.Search) error {
	query := `INSERT INTO search(keyword, page) VALUES(?, ?);`

	ctxSegment, tracerSpan := distributedtracing.StartPostgresSegment(ctx, constant.SegmentRepository+"InsertSearchHistory", query)
	defer tracerSpan.End()

	_, err := r.db.ExecContext(ctxSegment, query, row.Keyword, row.Page)
	if err != nil {
		return err
	}
	return nil
}
