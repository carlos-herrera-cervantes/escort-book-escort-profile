// Code generated by MockGen. DO NOT EDIT.
// Source: escort-book-escort-profile/repositories (interfaces: IProfileStatusRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	models "escort-book-escort-profile/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIProfileStatusRepository is a mock of IProfileStatusRepository interface.
type MockIProfileStatusRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIProfileStatusRepositoryMockRecorder
}

// MockIProfileStatusRepositoryMockRecorder is the mock recorder for MockIProfileStatusRepository.
type MockIProfileStatusRepositoryMockRecorder struct {
	mock *MockIProfileStatusRepository
}

// NewMockIProfileStatusRepository creates a new mock instance.
func NewMockIProfileStatusRepository(ctrl *gomock.Controller) *MockIProfileStatusRepository {
	mock := &MockIProfileStatusRepository{ctrl: ctrl}
	mock.recorder = &MockIProfileStatusRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProfileStatusRepository) EXPECT() *MockIProfileStatusRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIProfileStatusRepository) Create(arg0 context.Context, arg1 *models.ProfileStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIProfileStatusRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIProfileStatusRepository)(nil).Create), arg0, arg1)
}

// GetOne mocks base method.
func (m *MockIProfileStatusRepository) GetOne(arg0 context.Context, arg1 string) (models.ProfileStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOne", arg0, arg1)
	ret0, _ := ret[0].(models.ProfileStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOne indicates an expected call of GetOne.
func (mr *MockIProfileStatusRepositoryMockRecorder) GetOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockIProfileStatusRepository)(nil).GetOne), arg0, arg1)
}

// UpdateOne mocks base method.
func (m *MockIProfileStatusRepository) UpdateOne(arg0 context.Context, arg1 string, arg2 *models.ProfileStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOne", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOne indicates an expected call of UpdateOne.
func (mr *MockIProfileStatusRepositoryMockRecorder) UpdateOne(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOne", reflect.TypeOf((*MockIProfileStatusRepository)(nil).UpdateOne), arg0, arg1, arg2)
}
