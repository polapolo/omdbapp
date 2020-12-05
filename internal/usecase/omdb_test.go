package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/polapolo/omdbapp/internal/mock_usecase_provider"
	"github.com/polapolo/omdbapp/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestNewOMDBUsecase(t *testing.T) {
	type args struct {
		omdbService OMDBServiceInterface
	}
	tests := []struct {
		name string
		args args
		want OMDBUsecase
	}{
		{
			"success",
			args{},
			OMDBUsecase{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOMDBUsecase(tt.args.omdbService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOMDBUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOMDBUsecase_Search(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOMDBService := mock_usecase_provider.NewMockOMDBServiceInterface(ctrl)
	mockCtx := context.Background()

	type args struct {
		ctx     context.Context
		keyword string
		page    int
	}
	tests := []struct {
		name    string
		u       OMDBUsecase
		args    args
		want    OMDBSearchResponse
		wantErr bool
		mock    func()
	}{
		{
			"success",
			OMDBUsecase{mockOMDBService},
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
				mockOMDBService.EXPECT().
					Search(mockCtx, "keyword", int(1)).
					Return(service.OMDBSearchResponse{
						Search: []service.OMDBSearchResponseSearch{
							{},
						},
					}, nil)
			},
		},
		{
			"error",
			OMDBUsecase{mockOMDBService},
			args{
				mockCtx,
				"keyword",
				1,
			},
			OMDBSearchResponse{},
			true,
			func() {
				mockOMDBService.EXPECT().
					Search(mockCtx, "keyword", int(1)).
					Return(service.OMDBSearchResponse{}, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.u.Search(tt.args.ctx, tt.args.keyword, tt.args.page)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestOMDBUsecase_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOMDBService := mock_usecase_provider.NewMockOMDBServiceInterface(ctrl)
	mockCtx := context.Background()

	type args struct {
		ctx    context.Context
		imdbID string
	}
	tests := []struct {
		name    string
		u       OMDBUsecase
		args    args
		want    OMDBGetByIDResponse
		wantErr bool
		mock    func()
	}{
		{
			"success",
			OMDBUsecase{mockOMDBService},
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
				mockOMDBService.EXPECT().
					GetByID(mockCtx, "imdbid").
					Return(service.OMDBGetByIDResponse{
						Ratings: []service.OMDBGetByIDResponseRating{
							{},
						},
					}, nil)
			},
		},
		{
			"error",
			OMDBUsecase{mockOMDBService},
			args{
				mockCtx,
				"imdbid",
			},
			OMDBGetByIDResponse{},
			true,
			func() {
				mockOMDBService.EXPECT().
					GetByID(mockCtx, "imdbid").
					Return(service.OMDBGetByIDResponse{}, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.u.GetByID(tt.args.ctx, tt.args.imdbID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
