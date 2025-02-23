// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	enum "github.com/Rasikrr/learning_platform/internal/domain/enum"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ExecuteCode mocks base method.
func (m *MockClient) ExecuteCode(arg0 context.Context, code string, lang enum.ProgrammingLanguage) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteCode", arg0, code, lang)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteCode indicates an expected call of ExecuteCode.
func (mr *MockClientMockRecorder) ExecuteCode(arg0, code, lang interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteCode", reflect.TypeOf((*MockClient)(nil).ExecuteCode), arg0, code, lang)
}

// ExecuteTestCase mocks base method.
func (m *MockClient) ExecuteTestCase(arg0 context.Context, code, stdin string, lang enum.ProgrammingLanguage) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteTestCase", arg0, code, stdin, lang)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteTestCase indicates an expected call of ExecuteTestCase.
func (mr *MockClientMockRecorder) ExecuteTestCase(arg0, code, stdin, lang interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteTestCase", reflect.TypeOf((*MockClient)(nil).ExecuteTestCase), arg0, code, stdin, lang)
}
