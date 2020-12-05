package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/polapolo/omdbapp/internal/mock_service_provider"
	"github.com/polapolo/omdbapp/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestNewOMDBService(t *testing.T) {
	type args struct {
		omdbAPIRepository OMDBApiRepositoryInterface
	}
	tests := []struct {
		name string
		args args
		want OMDBService
	}{
		{
			"success",
			args{},
			OMDBService{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOMDBService(tt.args.omdbAPIRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOMDBService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOMDBService_Search(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOMDBApiRepository := mock_service_provider.NewMockOMDBApiRepositoryInterface(ctrl)
	mockCtx := context.Background()

	type args struct {
		ctx     context.Context
		keyword string
		page    int
	}
	tests := []struct {
		name    string
		s       OMDBService
		args    args
		want    OMDBSearchResponse
		wantErr bool
		mock    func()
	}{
		{
			"success",
			OMDBService{mockOMDBApiRepository},
			args{
				mockCtx,
				"keyword",
				1,
			},
			OMDBSearchResponse{
				Search: []OMDBSearchResponseSearch{
					{},
				},
			},
			false,
			func() {
				mockOMDBApiRepository.EXPECT().
					Search(mockCtx, "keyword", int(1)).
					Return(repository.OMDBSearchResponse{
						Search: []repository.OMDBSearchResponseSearch{
							{},
						},
					}, nil)
			},
		},
		{
			"error",
			OMDBService{mockOMDBApiRepository},
			args{
				mockCtx,
				"keyword",
				1,
			},
			OMDBSearchResponse{},
			true,
			func() {
				mockOMDBApiRepository.EXPECT().
					Search(mockCtx, "keyword", int(1)).
					Return(repository.OMDBSearchResponse{}, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.Search(tt.args.ctx, tt.args.keyword, tt.args.page)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestOMDBService_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOMDBApiRepository := mock_service_provider.NewMockOMDBApiRepositoryInterface(ctrl)
	mockCtx := context.Background()

	type args struct {
		ctx    context.Context
		imdbID string
	}
	tests := []struct {
		name    string
		s       OMDBService
		args    args
		want    OMDBGetByIDResponse
		wantErr bool
		mock    func()
	}{
		{
			"success",
			OMDBService{mockOMDBApiRepository},
			args{
				mockCtx,
				"imdbid",
			},
			OMDBGetByIDResponse{
				Ratings: []OMDBGetByIDResponseRating{
					{},
				},
			},
			false,
			func() {
				mockOMDBApiRepository.EXPECT().
					GetByID(mockCtx, "imdbid").
					Return(repository.OMDBGetByIDResponse{
						Ratings: []repository.OMDBGetByIDResponseRating{
							{},
						},
					}, nil)
			},
		},
		{
			"error",
			OMDBService{mockOMDBApiRepository},
			args{
				mockCtx,
				"imdbid",
			},
			OMDBGetByIDResponse{},
			true,
			func() {
				mockOMDBApiRepository.EXPECT().
					GetByID(mockCtx, "imdbid").
					Return(repository.OMDBGetByIDResponse{}, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.GetByID(tt.args.ctx, tt.args.imdbID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
