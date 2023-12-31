// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_repo.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	models "sinarmas/models"

	gomock "github.com/golang/mock/gomock"
)

// MockIUserRepo is a mock of IUserRepo interface.
type MockIUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepoMockRecorder
}

// MockIUserRepoMockRecorder is the mock recorder for MockIUserRepo.
type MockIUserRepoMockRecorder struct {
	mock *MockIUserRepo
}

// NewMockIUserRepo creates a new mock instance.
func NewMockIUserRepo(ctrl *gomock.Controller) *MockIUserRepo {
	mock := &MockIUserRepo{ctrl: ctrl}
	mock.recorder = &MockIUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepo) EXPECT() *MockIUserRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIUserRepo) Create(user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIUserRepoMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUserRepo)(nil).Create), user)
}

// GetByUserAndRequestId mocks base method.
func (m *MockIUserRepo) GetByUserAndRequestId(userId, requestId string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserAndRequestId", userId, requestId)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserAndRequestId indicates an expected call of GetByUserAndRequestId.
func (mr *MockIUserRepoMockRecorder) GetByUserAndRequestId(userId, requestId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserAndRequestId", reflect.TypeOf((*MockIUserRepo)(nil).GetByUserAndRequestId), userId, requestId)
}

// Save mocks base method.
func (m *MockIUserRepo) Save(user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIUserRepoMockRecorder) Save(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIUserRepo)(nil).Save), user)
}
