// Code generated by MockGen. DO NOT EDIT.
// Source: omdbapi.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	http "github.com/polapolo/omdbapp/internal/client/http"
	reflect "reflect"
)

// MockOMDBApiHTTPClientInterface is a mock of OMDBApiHTTPClientInterface interface
type MockOMDBApiHTTPClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockOMDBApiHTTPClientInterfaceMockRecorder
}

// MockOMDBApiHTTPClientInterfaceMockRecorder is the mock recorder for MockOMDBApiHTTPClientInterface
type MockOMDBApiHTTPClientInterfaceMockRecorder struct {
	mock *MockOMDBApiHTTPClientInterface
}

// NewMockOMDBApiHTTPClientInterface creates a new mock instance
func NewMockOMDBApiHTTPClientInterface(ctrl *gomock.Controller) *MockOMDBApiHTTPClientInterface {
	mock := &MockOMDBApiHTTPClientInterface{ctrl: ctrl}
	mock.recorder = &MockOMDBApiHTTPClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOMDBApiHTTPClientInterface) EXPECT() *MockOMDBApiHTTPClientInterfaceMockRecorder {
	return m.recorder
}

// Search mocks base method
func (m *MockOMDBApiHTTPClientInterface) Search(ctx context.Context, keyword string, page int) (http.OMDBSearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, keyword, page)
	ret0, _ := ret[0].(http.OMDBSearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockOMDBApiHTTPClientInterfaceMockRecorder) Search(ctx, keyword, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockOMDBApiHTTPClientInterface)(nil).Search), ctx, keyword, page)
}

// GetByID mocks base method
func (m *MockOMDBApiHTTPClientInterface) GetByID(ctx context.Context, imdbID string) (http.OMDBGetByIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, imdbID)
	ret0, _ := ret[0].(http.OMDBGetByIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockOMDBApiHTTPClientInterfaceMockRecorder) GetByID(ctx, imdbID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockOMDBApiHTTPClientInterface)(nil).GetByID), ctx, imdbID)
}