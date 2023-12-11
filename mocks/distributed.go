// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-co-op/gocron/v2 (interfaces: Elector,Locker,Lock)
//
// Generated by this command:
//
//	mockgen -destination=mocks/distributed.go -package=gocronmocks . Elector,Locker,Lock
//
// Package gocronmocks is a generated GoMock package.
package gocronmocks

import (
	context "context"
	reflect "reflect"

	gocron "github.com/go-co-op/gocron/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockElector is a mock of Elector interface.
type MockElector struct {
	ctrl     *gomock.Controller
	recorder *MockElectorMockRecorder
}

// MockElectorMockRecorder is the mock recorder for MockElector.
type MockElectorMockRecorder struct {
	mock *MockElector
}

// NewMockElector creates a new mock instance.
func NewMockElector(ctrl *gomock.Controller) *MockElector {
	mock := &MockElector{ctrl: ctrl}
	mock.recorder = &MockElectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockElector) EXPECT() *MockElectorMockRecorder {
	return m.recorder
}

// IsLeader mocks base method.
func (m *MockElector) IsLeader(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsLeader indicates an expected call of IsLeader.
func (mr *MockElectorMockRecorder) IsLeader(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLeader", reflect.TypeOf((*MockElector)(nil).IsLeader), arg0)
}

// MockLocker is a mock of Locker interface.
type MockLocker struct {
	ctrl     *gomock.Controller
	recorder *MockLockerMockRecorder
}

// MockLockerMockRecorder is the mock recorder for MockLocker.
type MockLockerMockRecorder struct {
	mock *MockLocker
}

// NewMockLocker creates a new mock instance.
func NewMockLocker(ctrl *gomock.Controller) *MockLocker {
	mock := &MockLocker{ctrl: ctrl}
	mock.recorder = &MockLockerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocker) EXPECT() *MockLockerMockRecorder {
	return m.recorder
}

// Lock mocks base method.
func (m *MockLocker) Lock(arg0 context.Context, arg1 string) (gocron.Lock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock", arg0, arg1)
	ret0, _ := ret[0].(gocron.Lock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lock indicates an expected call of Lock.
func (mr *MockLockerMockRecorder) Lock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockLocker)(nil).Lock), arg0, arg1)
}

// MockLock is a mock of Lock interface.
type MockLock struct {
	ctrl     *gomock.Controller
	recorder *MockLockMockRecorder
}

// MockLockMockRecorder is the mock recorder for MockLock.
type MockLockMockRecorder struct {
	mock *MockLock
}

// NewMockLock creates a new mock instance.
func NewMockLock(ctrl *gomock.Controller) *MockLock {
	mock := &MockLock{ctrl: ctrl}
	mock.recorder = &MockLockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLock) EXPECT() *MockLockMockRecorder {
	return m.recorder
}

// Unlock mocks base method.
func (m *MockLock) Unlock(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unlock indicates an expected call of Unlock.
func (mr *MockLockMockRecorder) Unlock(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockLock)(nil).Unlock), arg0)
}
