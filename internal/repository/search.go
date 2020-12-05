package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/polapolo/omdbapp/internal/entity"
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
	_, err := r.db.ExecContext(ctx, `INSERT INTO search(keyword, page) VALUES(?, ?);`, row.Keyword, row.Page)
	if err != nil {
		return err
	}
	return nil
}
