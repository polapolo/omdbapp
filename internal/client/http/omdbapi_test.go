package http

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestNewOMDBApiHTTPClient(t *testing.T) {
	type args struct {
		host   string
		apiKey string
	}
	tests := []struct {
		name string
		args args
		want OMDBApiHTTPClient
	}{
		{
			"success",
			args{},
			OMDBApiHTTPClient{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOMDBApiHTTPClient(tt.args.host, tt.args.apiKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOMDBApiHTTPClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOMDBApiHTTPClient_Search(t *testing.T) {
	defer gock.Off()

	mockCtx := context.Background()
	defaultIoutilReadAll := ioutilReadAll
	defaultHTTPGet := httpGet
	defaultURLParse := urlParse

	restore := func() {
		ioutilReadAll = defaultIoutilReadAll
		httpGet = defaultHTTPGet
		urlParse = defaultURLParse
	}

	type args struct {
		ctx     context.Context
		keyword string
		page    int
	}
	tests := []struct {
		name    string
		c       OMDBApiHTTPClient
		args    args
		want    OMDBSearchResponse
		wantErr bool
		mock    func()
	}{
		{
			"success",
			OMDBApiHTTPClient{
				"http://server.local",
				"apikey",
			},
			args{
				mockCtx,
				"keyword",
				1,
			},
			OMDBSearchResponse{},
			false,
			func() {
				gock.New("http://server.local").
					Get("/").
					Reply(200).
					JSON(`{}`)
			},
		},
		{
			"error response",
			OMDBApiHTTPClient{
				"http://server.local",
				"apikey",
			},
			args{
				mockCtx,
				"keyword",
				1,
			},
			OMDBSearchResponse{
				Response: "False",
			},
			true,
			func() {
				gock.New("http://server.local").
					Get("/").
					Reply(200).
					JSON(`{"Response": "False"}`)
			},
		},
		{
			"error unmarshal",
			OMDBApiHTTPClient{
				"http://server.local",
				"apikey",
			},
			args{
				mockCtx,
				"keyword",
				1,
			},
			OMDBSearchResponse{},
			true,
			func() {
				gock.New("http://server.local").
					Get("/").
					Reply(200).
					JSON(`a`)
			},
		},
		{
			"error ioutilReadAll",
			OMDBApiHTTPClient{
				"http://server.local",
				"apikey",
			},
			args{
				mockCtx,
				"keyword",
				1,
			},
			OMDBSearchResponse{},
			true,
			func() {
				gock.New("http://server.local").
					Get("/").
					Reply(200).
					JSON(`a`)

				ioutilReadAll = func(r io.Reader) ([]byte, error) {
					return nil, errors.New("error")
				}
			},
		},
		{
			"error httpGet",
			OMDBApiHTTPClient{
				"http://server.local",
				"apikey",
			},
			args{
				mockCtx,
				"keyword",
				1,
			},
			OMDBSearchResponse{},
			true,
			func() {
				httpGet = func(url string) (resp *http.Response, err error) {
					return nil, errors.New("error")
				}
			},
		},
		{
			"error urlParse",
			OMDBApiHTTPClient{
				"http://server.local",
				"apikey",
			},
			args{
				mockCtx,
				"keyword",
				1,
			},
			OMDBSearchResponse{},
			true,
			func() {
				urlParse = func(rawurl string) (*url.URL, error) {
					return nil, errors.New("error")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.c.Search(tt.args.ctx, tt.args.keyword, tt.args.page)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
			restore()
		})
	}
}

func TestOMDBApiHTTPClient_GetByID(t *testing.T) {
	defer gock.Off()

	mockCtx := context.Background()
	defaultIoutilReadAll := ioutilReadAll
	defaultHTTPGet := httpGet
	defaultURLParse := urlParse

	restore := func() {
		ioutilReadAll = defaultIoutilReadAll
		httpGet = defaultHTTPGet
		urlParse = defaultURLParse
	}

	type args struct {
		ctx    context.Context
		imdbID string
	}
	tests := []struct {
		name    string
		c       OMDBApiHTTPClient
		args    args
		want    OMDBGetByIDResponse
		wantErr bool
		mock    func()
	}{
		{
			"success",
			OMDBApiHTTPClient{
				"http://server.local",
				"apikey",
			},
			args{
				mockCtx,
				"imdbid",
			},
			OMDBGetByIDResponse{},
			false,
			func() {
				gock.New("http://server.local").
					Get("/").
					Reply(200).
					JSON(`{}`)
			},
		},
		{
			"error response",
			OMDBApiHTTPClient{
				"http://server.local",
				"apikey",
			},
			args{
				mockCtx,
				"imdbid",
			},
			OMDBGetByIDResponse{
				Response: "False",
			},
			true,
			func() {
				gock.New("http://server.local").
					Get("/").
					Reply(200).
					JSON(`{"Response": "False"}`)
			},
		},
		{
			"error unmarshal",
			OMDBApiHTTPClient{
				"http://server.local",
				"apikey",
			},
			args{
				mockCtx,
				"imdbid",
			},
			OMDBGetByIDResponse{},
			true,
			func() {
				gock.New("http://server.local").
					Get("/").
					Reply(200).
					JSON(`a`)
			},
		},
		{
			"error ioutilReadAll",
			OMDBApiHTTPClient{
				"http://server.local",
				"apikey",
			},
			args{
				mockCtx,
				"imdbid",
			},
			OMDBGetByIDResponse{},
			true,
			func() {
				gock.New("http://server.local").
					Get("/").
					Reply(200).
					JSON(`a`)

				ioutilReadAll = func(r io.Reader) ([]byte, error) {
					return nil, errors.New("error")
				}
			},
		},
		{
			"error httpGet",
			OMDBApiHTTPClient{
				"http://server.local",
				"apikey",
			},
			args{
				mockCtx,
				"imdbid",
			},
			OMDBGetByIDResponse{},
			true,
			func() {
				httpGet = func(url string) (resp *http.Response, err error) {
					return nil, errors.New("error")
				}
			},
		},
		{
			"error urlParse",
			OMDBApiHTTPClient{
				"http://server.local",
				"apikey",
			},
			args{
				mockCtx,
				"imdbid",
			},
			OMDBGetByIDResponse{},
			true,
			func() {
				urlParse = func(rawurl string) (*url.URL, error) {
					return nil, errors.New("error")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.c.GetByID(tt.args.ctx, tt.args.imdbID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
			restore()
		})
	}
}
