// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entity "github.com/Rasikrr/learning_platform/internal/domain/entity"
	gomock "github.com/golang/mock/gomock"
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

// GetByID mocks base method.
func (m *MockRepository) GetByID(ctx context.Context, id string) (*entity.PracticalTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*entity.PracticalTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockRepositoryMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockRepository)(nil).GetByID), ctx, id)
}

// GetByTopicIDAndOrderNum mocks base method.
func (m *MockRepository) GetByTopicIDAndOrderNum(ctx context.Context, id string, order int) (*entity.PracticalTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTopicIDAndOrderNum", ctx, id, order)
	ret0, _ := ret[0].(*entity.PracticalTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTopicIDAndOrderNum indicates an expected call of GetByTopicIDAndOrderNum.
func (mr *MockRepositoryMockRecorder) GetByTopicIDAndOrderNum(ctx, id, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTopicIDAndOrderNum", reflect.TypeOf((*MockRepository)(nil).GetByTopicIDAndOrderNum), ctx, id, order)
}

// GetByTopicIDs mocks base method.
func (m *MockRepository) GetByTopicIDs(ctx context.Context, ids []string) ([]*entity.PracticalTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTopicIDs", ctx, ids)
	ret0, _ := ret[0].([]*entity.PracticalTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTopicIDs indicates an expected call of GetByTopicIDs.
func (mr *MockRepositoryMockRecorder) GetByTopicIDs(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTopicIDs", reflect.TypeOf((*MockRepository)(nil).GetByTopicIDs), ctx, ids)
}
