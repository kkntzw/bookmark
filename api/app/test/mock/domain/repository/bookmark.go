// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/domain/repository/bookmark.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	bookmark "github.com/kkntzw/bookmark/internal/domain/bookmark"
)

// MockBookmark is a mock of Bookmark interface.
type MockBookmark struct {
	ctrl     *gomock.Controller
	recorder *MockBookmarkMockRecorder
}

// MockBookmarkMockRecorder is the mock recorder for MockBookmark.
type MockBookmarkMockRecorder struct {
	mock *MockBookmark
}

// NewMockBookmark creates a new mock instance.
func NewMockBookmark(ctrl *gomock.Controller) *MockBookmark {
	mock := &MockBookmark{ctrl: ctrl}
	mock.recorder = &MockBookmarkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookmark) EXPECT() *MockBookmarkMockRecorder {
	return m.recorder
}

// FindByID mocks base method.
func (m *MockBookmark) FindByID(id *bookmark.ID) (*bookmark.Bookmark, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*bookmark.Bookmark)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockBookmarkMockRecorder) FindByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockBookmark)(nil).FindByID), id)
}

// NextID mocks base method.
func (m *MockBookmark) NextID() *bookmark.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextID")
	ret0, _ := ret[0].(*bookmark.ID)
	return ret0
}

// NextID indicates an expected call of NextID.
func (mr *MockBookmarkMockRecorder) NextID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextID", reflect.TypeOf((*MockBookmark)(nil).NextID))
}

// Save mocks base method.
func (m *MockBookmark) Save(bookmark *bookmark.Bookmark) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", bookmark)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockBookmarkMockRecorder) Save(bookmark interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockBookmark)(nil).Save), bookmark)
}