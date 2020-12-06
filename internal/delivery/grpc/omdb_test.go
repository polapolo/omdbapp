package grpc

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/polapolo/omdbapp/internal/delivery/grpc/proto"
	"github.com/polapolo/omdbapp/internal/mock_deliver_grpc_provider"
	"github.com/polapolo/omdbapp/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestNewOMDBGRPCService(t *testing.T) {
	type args struct {
		omdbUsecase OMDBUsecaseInterface
	}
	tests := []struct {
		name string
		args args
		want OMDBGRPCService
	}{
		{
			"success",
			args{},
			OMDBGRPCService{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOMDBGRPCService(tt.args.omdbUsecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOMDBGRPCService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOMDBGRPCService_MovieSearch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOMDBUsecase := mock_deliver_grpc_provider.NewMockOMDBUsecaseInterface(ctrl)
	mockCtx := context.Background()

	type args struct {
		ctx     context.Context
		request *proto.MovieSearchRequest
	}
	tests := []struct {
		name    string
		h       OMDBGRPCService
		args    args
		want    *proto.MovieSearchResponse
		wantErr bool
		mock    func()
	}{
		{
			"success",
			OMDBGRPCService{mockOMDBUsecase},
			args{
				mockCtx,
				&proto.MovieSearchRequest{
					Keyword: "batman",
					Page:    1,
				},
			},
			&proto.MovieSearchResponse{
				Data: &proto.MovieSearchResponseData{
					Search: []*proto.MovieSearchResponseDataSearch{
						{},
					},
				},
			},
			false,
			func() {
				mockOMDBUsecase.EXPECT().
					Search(context.Background(), "batman", int(1)).
					Return(usecase.OMDBSearchResponse{
						Search: []usecase.OMDBSearchResponseSearch{
							{},
						},
					}, nil)
			},
		},
		{
			"error",
			OMDBGRPCService{mockOMDBUsecase},
			args{
				mockCtx,
				&proto.MovieSearchRequest{
					Keyword: "batman",
					Page:    1,
				},
			},
			nil,
			true,
			func() {
				mockOMDBUsecase.EXPECT().
					Search(context.Background(), "batman", int(1)).
					Return(usecase.OMDBSearchResponse{}, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.h.MovieSearch(tt.args.ctx, tt.args.request)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestOMDBGRPCService_MovieDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOMDBUsecase := mock_deliver_grpc_provider.NewMockOMDBUsecaseInterface(ctrl)
	mockCtx := context.Background()

	type args struct {
		ctx     context.Context
		request *proto.MovieDetailRequest
	}
	tests := []struct {
		name    string
		h       OMDBGRPCService
		args    args
		want    *proto.MovieDetailResponse
		wantErr bool
		mock    func()
	}{
		{
			"success",
			OMDBGRPCService{mockOMDBUsecase},
			args{
				mockCtx,
				&proto.MovieDetailRequest{
					Id: "imdbid",
				},
			},
			&proto.MovieDetailResponse{
				Data: &proto.MovieDetailResponseData{
					Ratings: []*proto.MovieDetailResponseDataRating{
						{},
					},
				},
			},
			false,
			func() {
				mockOMDBUsecase.EXPECT().
					GetByID(context.Background(), "imdbid").
					Return(usecase.OMDBGetByIDResponse{
						Ratings: []usecase.OMDBGetByIDResponseRating{
							{},
						},
					}, nil)
			},
		},
		{
			"error",
			OMDBGRPCService{mockOMDBUsecase},
			args{
				mockCtx,
				&proto.MovieDetailRequest{
					Id: "imdbid",
				},
			},
			nil,
			true,
			func() {
				mockOMDBUsecase.EXPECT().
					GetByID(context.Background(), "imdbid").
					Return(usecase.OMDBGetByIDResponse{}, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.h.MovieDetail(tt.args.ctx, tt.args.request)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
