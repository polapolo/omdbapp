package repository

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/polapolo/omdbapp/internal/entity"
)

func TestNewSearchRepository(t *testing.T) {
	type args struct {
		db *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want SearchRepository
	}{
		{
			"success",
			args{},
			SearchRepository{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSearchRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSearchRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchRepository_InsertSearchHistory(t *testing.T) {
	mockDB, mockQuery, _ := sqlmock.New()
	defer mockDB.Close()
	mockSqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	type args struct {
		ctx context.Context
		row entity.Search
	}
	tests := []struct {
		name    string
		r       SearchRepository
		args    args
		wantErr bool
		mock    func()
	}{
		{
			"success",
			SearchRepository{
				db: mockSqlxDB,
			},
			args{
				context.Background(),
				entity.Search{
					Keyword: "keyword",
					Page:    1,
				},
			},
			false,
			func() {
				mockQuery.ExpectExec(`INSERT INTO search`).WithArgs("keyword", int(1)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			"error",
			SearchRepository{
				db: mockSqlxDB,
			},
			args{
				context.Background(),
				entity.Search{
					Keyword: "keyword",
					Page:    1,
				},
			},
			true,
			func() {
				mockQuery.ExpectExec(`INSERT INTO search`).WithArgs("keyword", int(1)).WillReturnError(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			if err := tt.r.InsertSearchHistory(tt.args.ctx, tt.args.row); (err != nil) != tt.wantErr {
				t.Errorf("SearchRepository.InsertSearchHistory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
