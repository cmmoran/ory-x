// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/x/hasherx (interfaces: PBKDF2Configurator)
//
// Generated by this command:
//
//	mockgen -package hasherx_test -destination hasherx/mocks_pkdbf2_test.go github.com/ory/x/hasherx PBKDF2Configurator
//

// Package hasherx_test is a generated GoMock package.
package hasherx_test

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	hasherx "github.com/ory/x/hasherx"
)

// MockPBKDF2Configurator is a mock of PBKDF2Configurator interface.
type MockPBKDF2Configurator struct {
	ctrl     *gomock.Controller
	recorder *MockPBKDF2ConfiguratorMockRecorder
	isgomock struct{}
}

// MockPBKDF2ConfiguratorMockRecorder is the mock recorder for MockPBKDF2Configurator.
type MockPBKDF2ConfiguratorMockRecorder struct {
	mock *MockPBKDF2Configurator
}

// NewMockPBKDF2Configurator creates a new mock instance.
func NewMockPBKDF2Configurator(ctrl *gomock.Controller) *MockPBKDF2Configurator {
	mock := &MockPBKDF2Configurator{ctrl: ctrl}
	mock.recorder = &MockPBKDF2ConfiguratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPBKDF2Configurator) EXPECT() *MockPBKDF2ConfiguratorMockRecorder {
	return m.recorder
}

// HasherPBKDF2Config mocks base method.
func (m *MockPBKDF2Configurator) HasherPBKDF2Config(ctx context.Context) *hasherx.PBKDF2Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasherPBKDF2Config", ctx)
	ret0, _ := ret[0].(*hasherx.PBKDF2Config)
	return ret0
}

// HasherPBKDF2Config indicates an expected call of HasherPBKDF2Config.
func (mr *MockPBKDF2ConfiguratorMockRecorder) HasherPBKDF2Config(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasherPBKDF2Config", reflect.TypeOf((*MockPBKDF2Configurator)(nil).HasherPBKDF2Config), ctx)
}
