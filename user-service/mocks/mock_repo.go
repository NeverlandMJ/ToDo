// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NeverlandMJ/ToDo/user-service/server (interfaces: Repository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entity "github.com/NeverlandMJ/ToDo/user-service/pkg/entity"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
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

// ChangePassword mocks base method.
func (m *MockRepository) ChangePassword(arg0 context.Context, arg1 uuid.UUID, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockRepositoryMockRecorder) ChangePassword(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockRepository)(nil).ChangePassword), arg0, arg1, arg2, arg3)
}

// ChangeUserName mocks base method.
func (m *MockRepository) ChangeUserName(arg0 context.Context, arg1 uuid.UUID, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserName", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeUserName indicates an expected call of ChangeUserName.
func (mr *MockRepositoryMockRecorder) ChangeUserName(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserName", reflect.TypeOf((*MockRepository)(nil).ChangeUserName), arg0, arg1, arg2)
}

// CreateUser mocks base method.
func (m *MockRepository) CreateUser(arg0 context.Context, arg1 entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockRepositoryMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockRepository)(nil).CreateUser), arg0, arg1)
}

// DeleteAccount mocks base method.
func (m *MockRepository) DeleteAccount(arg0 context.Context, arg1 uuid.UUID, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockRepositoryMockRecorder) DeleteAccount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockRepository)(nil).DeleteAccount), arg0, arg1, arg2, arg3)
}

// GetUser mocks base method.
func (m *MockRepository) GetUser(arg0 context.Context, arg1, arg2 string) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockRepositoryMockRecorder) GetUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockRepository)(nil).GetUser), arg0, arg1, arg2)
}
