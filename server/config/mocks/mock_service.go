// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cpoile/mattermost-plugin-later/server/config (interfaces: Service)

// Package mock_config is a generated GoMock package.
package mock_config

import (
	reflect "reflect"

	config "github.com/cpoile/mattermost-plugin-later/server/config"
	gomock "github.com/golang/mock/gomock"
	model "github.com/mattermost/mattermost-server/v6/model"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetConfiguration mocks base method.
func (m *MockService) GetConfiguration() *config.Configuration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfiguration")
	ret0, _ := ret[0].(*config.Configuration)
	return ret0
}

// GetConfiguration indicates an expected call of GetConfiguration.
func (mr *MockServiceMockRecorder) GetConfiguration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfiguration", reflect.TypeOf((*MockService)(nil).GetConfiguration))
}

// GetManifest mocks base method.
func (m *MockService) GetManifest() *model.Manifest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManifest")
	ret0, _ := ret[0].(*model.Manifest)
	return ret0
}

// GetManifest indicates an expected call of GetManifest.
func (mr *MockServiceMockRecorder) GetManifest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManifest", reflect.TypeOf((*MockService)(nil).GetManifest))
}

// RegisterConfigChangeListener mocks base method.
func (m *MockService) RegisterConfigChangeListener(arg0 func()) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterConfigChangeListener", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// RegisterConfigChangeListener indicates an expected call of RegisterConfigChangeListener.
func (mr *MockServiceMockRecorder) RegisterConfigChangeListener(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterConfigChangeListener", reflect.TypeOf((*MockService)(nil).RegisterConfigChangeListener), arg0)
}

// UnregisterConfigChangeListener mocks base method.
func (m *MockService) UnregisterConfigChangeListener(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnregisterConfigChangeListener", arg0)
}

// UnregisterConfigChangeListener indicates an expected call of UnregisterConfigChangeListener.
func (mr *MockServiceMockRecorder) UnregisterConfigChangeListener(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterConfigChangeListener", reflect.TypeOf((*MockService)(nil).UnregisterConfigChangeListener), arg0)
}

// UpdateConfiguration mocks base method.
func (m *MockService) UpdateConfiguration(arg0 func(*config.Configuration)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConfiguration", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConfiguration indicates an expected call of UpdateConfiguration.
func (mr *MockServiceMockRecorder) UpdateConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfiguration", reflect.TypeOf((*MockService)(nil).UpdateConfiguration), arg0)
}
