package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/polapolo/omdbapp/internal/mock_deliver_http_provider"
	"github.com/polapolo/omdbapp/internal/usecase"
)

func TestNewOMDBHttpHandler(t *testing.T) {
	type args struct {
		omdbUsecase OMDBUsecaseInterface
	}
	tests := []struct {
		name string
		args args
		want OMDBHttpHandler
	}{
		{
			"success",
			args{},
			OMDBHttpHandler{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOMDBHttpHandler(tt.args.omdbUsecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOMDBHttpHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOMDBHttpHandler_Search(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOMDBUsecase := mock_deliver_http_provider.NewMockOMDBUsecaseInterface(ctrl)
	// mock success
	cSuccess, _ := gin.CreateTestContext(httptest.NewRecorder())
	cSuccess.Request, _ = http.NewRequest("GET", "http://example.com?keyword=batman&page=2", nil)
	// mock error
	cError, _ := gin.CreateTestContext(httptest.NewRecorder())
	cError.Request, _ = http.NewRequest("GET", "http://example.com?keyword=batman&page=a", nil)

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    OMDBHttpHandler
		args args
		mock func()
	}{
		{
			"success",
			OMDBHttpHandler{mockOMDBUsecase},
			args{
				cSuccess,
			},
			func() {
				mockOMDBUsecase.EXPECT().
					Search(context.Background(), "batman", int(2)).
					Return(usecase.OMDBSearchResponse{}, nil)
			},
		},
		{
			"error h.omdbUsecase.Search",
			OMDBHttpHandler{mockOMDBUsecase},
			args{
				cSuccess,
			},
			func() {
				mockOMDBUsecase.EXPECT().
					Search(context.Background(), "batman", int(2)).
					Return(usecase.OMDBSearchResponse{}, errors.New("error"))
			},
		},
		{
			"error strconv.Atoi",
			OMDBHttpHandler{mockOMDBUsecase},
			args{
				cError,
			},
			func() {},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Search(tt.args.c)
		})
	}
}

func TestOMDBHttpHandler_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOMDBUsecase := mock_deliver_http_provider.NewMockOMDBUsecaseInterface(ctrl)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("GET", "http://localhost/detail/imdbid", nil)
	c.Params = gin.Params{
		gin.Param{
			Key:   "id",
			Value: "imdbid",
		},
	}

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    OMDBHttpHandler
		args args
		mock func()
	}{
		{
			"success",
			OMDBHttpHandler{mockOMDBUsecase},
			args{
				c,
			},
			func() {
				mockOMDBUsecase.EXPECT().
					GetByID(context.Background(), "imdbid").
					Return(usecase.OMDBGetByIDResponse{}, nil)
			},
		},
		{
			"error",
			OMDBHttpHandler{mockOMDBUsecase},
			args{
				c,
			},
			func() {
				mockOMDBUsecase.EXPECT().
					GetByID(context.Background(), "imdbid").
					Return(usecase.OMDBGetByIDResponse{}, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			tt.h.GetByID(tt.args.c)
		})
	}
}
