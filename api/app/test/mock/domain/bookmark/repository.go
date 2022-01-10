// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/domain/bookmark/repository.go

// Package mock_bookmark is a generated GoMock package.
package mock_bookmark

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	bookmark "github.com/kkntzw/bookmark/internal/domain/bookmark"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// FindByID mocks base method.
func (m *MockRepository) FindByID(id *bookmark.ID) (*bookmark.Bookmark, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*bookmark.Bookmark)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockRepositoryMockRecorder) FindByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockRepository)(nil).FindByID), id)
}

// NextID mocks base method.
func (m *MockRepository) NextID() *bookmark.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextID")
	ret0, _ := ret[0].(*bookmark.ID)
	return ret0
}

// NextID indicates an expected call of NextID.
func (mr *MockRepositoryMockRecorder) NextID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextID", reflect.TypeOf((*MockRepository)(nil).NextID))
}

// Save mocks base method.
func (m *MockRepository) Save(bookmark *bookmark.Bookmark) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", bookmark)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockRepositoryMockRecorder) Save(bookmark interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRepository)(nil).Save), bookmark)
}
