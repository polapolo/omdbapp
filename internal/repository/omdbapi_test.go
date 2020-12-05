package repository

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	httpClient "github.com/polapolo/omdbapp/internal/client/http"
	"github.com/polapolo/omdbapp/internal/mock_repository_provider"
	"github.com/stretchr/testify/assert"
)

func TestNewOMDBApiRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOMDBHTTPClient := mock_repository_provider.NewMockOMDBApiHTTPClientInterface(ctrl)

	type args struct {
		omdbAPIHTTPClient OMDBApiHTTPClientInterface
	}
	tests := []struct {
		name string
		args args
		want OMDBApiRepository
	}{
		{
			"success",
			args{mockOMDBHTTPClient},
			OMDBApiRepository{
				mockOMDBHTTPClient,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOMDBApiRepository(tt.args.omdbAPIHTTPClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOMDBApiRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOMDBApiRepository_Search(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOMDBHTTPClient := mock_repository_provider.NewMockOMDBApiHTTPClientInterface(ctrl)
	mockCtx := context.Background()

	type args struct {
		ctx     context.Context
		keyword string
		page    int
	}
	tests := []struct {
		name    string
		r       OMDBApiRepository
		args    args
		want    OMDBSearchResponse
		wantErr bool
		mock    func()
	}{
		{
			"success",
			OMDBApiRepository{
				mockOMDBHTTPClient,
			},
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
				mockOMDBHTTPClient.EXPECT().
					Search(mockCtx, "keyword", int(1)).
					Return(httpClient.OMDBSearchResponse{
						Search: []httpClient.OMDBSearchResponseSearch{
							{},
						},
					}, nil)
			},
		},
		{
			"error r.omdbAPIHTTPClient.Search",
			OMDBApiRepository{
				mockOMDBHTTPClient,
			},
			args{
				mockCtx,
				"keyword",
				1,
			},
			OMDBSearchResponse{},
			true,
			func() {
				mockOMDBHTTPClient.EXPECT().
					Search(mockCtx, "keyword", int(1)).
					Return(httpClient.OMDBSearchResponse{}, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.r.Search(tt.args.ctx, tt.args.keyword, tt.args.page)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestOMDBApiRepository_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOMDBHTTPClient := mock_repository_provider.NewMockOMDBApiHTTPClientInterface(ctrl)
	mockCtx := context.Background()

	type args struct {
		ctx    context.Context
		imdbID string
	}
	tests := []struct {
		name    string
		r       OMDBApiRepository
		args    args
		want    OMDBGetByIDResponse
		wantErr bool
		mock    func()
	}{
		{
			"success",
			OMDBApiRepository{
				mockOMDBHTTPClient,
			},
			args{
				mockCtx,
				"tt4853102",
			},
			OMDBGetByIDResponse{
				Ratings: []OMDBGetByIDResponseRating{
					{},
				},
			},
			false,
			func() {
				mockOMDBHTTPClient.EXPECT().
					GetByID(mockCtx, "tt4853102").
					Return(httpClient.OMDBGetByIDResponse{
						Ratings: []httpClient.OMDBGetByIDResponseRating{
							{},
						},
					}, nil)
			},
		},
		{
			"error r.omdbAPIHTTPClient.GetByID",
			OMDBApiRepository{
				mockOMDBHTTPClient,
			},
			args{
				mockCtx,
				"tt4853102",
			},
			OMDBGetByIDResponse{},
			true,
			func() {
				mockOMDBHTTPClient.EXPECT().
					GetByID(mockCtx, "tt4853102").
					Return(httpClient.OMDBGetByIDResponse{}, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.r.GetByID(tt.args.ctx, tt.args.imdbID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
