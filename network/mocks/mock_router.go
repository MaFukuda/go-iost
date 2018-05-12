// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iost-official/prototype/network (interfaces: Router)

// Package protocol_mock is a generated GoMock package.
package protocol_mock

import (
	gomock "github.com/golang/mock/gomock"
	message "github.com/iost-official/prototype/core/message"
	network "github.com/iost-official/prototype/network"
	reflect "reflect"
)

// MockRouter is a mock of Router interface
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMockRecorder
}

// MockRouterMockRecorder is the mock recorder for MockRouter
type MockRouterMockRecorder struct {
	mock *MockRouter
}

// NewMockRouter creates a new mock instance
func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &MockRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouter) EXPECT() *MockRouterMockRecorder {
	return m.recorder
}

// Broadcast mocks base method
func (m *MockRouter) Broadcast(arg0 message.Message) {
	m.ctrl.Call(m, "Broadcast", arg0)
}

// Broadcast indicates an expected call of Broadcast
func (mr *MockRouterMockRecorder) Broadcast(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockRouter)(nil).Broadcast), arg0)
}

// CancelDownload mocks base method
func (m *MockRouter) CancelDownload(arg0, arg1 uint64) error {
	ret := m.ctrl.Call(m, "CancelDownload", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelDownload indicates an expected call of CancelDownload
func (mr *MockRouterMockRecorder) CancelDownload(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelDownload", reflect.TypeOf((*MockRouter)(nil).CancelDownload), arg0, arg1)
}

// Download mocks base method
func (m *MockRouter) Download(arg0, arg1 uint64) error {
	ret := m.ctrl.Call(m, "Download", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Download indicates an expected call of Download
func (mr *MockRouterMockRecorder) Download(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockRouter)(nil).Download), arg0, arg1)
}

// FilteredChan mocks base method
func (m *MockRouter) FilteredChan(arg0 network.Filter) (chan message.Message, error) {
	ret := m.ctrl.Call(m, "FilteredChan", arg0)
	ret0, _ := ret[0].(chan message.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilteredChan indicates an expected call of FilteredChan
func (mr *MockRouterMockRecorder) FilteredChan(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilteredChan", reflect.TypeOf((*MockRouter)(nil).FilteredChan), arg0)
}

// Init mocks base method
func (m *MockRouter) Init(arg0 network.Network, arg1 uint16) error {
	ret := m.ctrl.Call(m, "Init", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockRouterMockRecorder) Init(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockRouter)(nil).Init), arg0, arg1)
}

// Run mocks base method
func (m *MockRouter) Run() {
	m.ctrl.Call(m, "Run")
}

// Run indicates an expected call of Run
func (mr *MockRouterMockRecorder) Run() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockRouter)(nil).Run))
}

// Send mocks base method
func (m *MockRouter) Send(arg0 message.Message) {
	m.ctrl.Call(m, "Send", arg0)
}

// Send indicates an expected call of Send
func (mr *MockRouterMockRecorder) Send(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockRouter)(nil).Send), arg0)
}

// Stop mocks base method
func (m *MockRouter) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockRouterMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockRouter)(nil).Stop))
}