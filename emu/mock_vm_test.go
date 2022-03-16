// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/akita/mem/v3/vm (interfaces: PageTable)

package emu

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	vm "gitlab.com/akita/mem/v3/vm"
	ca "gitlab.com/akita/util/v2/ca"
)

// MockPageTable is a mock of PageTable interface.
type MockPageTable struct {
	ctrl     *gomock.Controller
	recorder *MockPageTableMockRecorder
}

// MockPageTableMockRecorder is the mock recorder for MockPageTable.
type MockPageTableMockRecorder struct {
	mock *MockPageTable
}

// NewMockPageTable creates a new mock instance.
func NewMockPageTable(ctrl *gomock.Controller) *MockPageTable {
	mock := &MockPageTable{ctrl: ctrl}
	mock.recorder = &MockPageTableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPageTable) EXPECT() *MockPageTableMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockPageTable) Find(arg0 vm.PID, arg1 uint64) (vm.Page, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(vm.Page)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockPageTableMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockPageTable)(nil).Find), arg0, arg1)
}

// Insert mocks base method.
func (m *MockPageTable) Insert(arg0 vm.Page) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Insert", arg0)
}

// Insert indicates an expected call of Insert.
func (mr *MockPageTableMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPageTable)(nil).Insert), arg0)
}

// Remove mocks base method.
func (m *MockPageTable) Remove(arg0 vm.PID, arg1 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", arg0, arg1)
}

// Remove indicates an expected call of Remove.
func (mr *MockPageTableMockRecorder) Remove(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockPageTable)(nil).Remove), arg0, arg1)
}

// Update mocks base method.
func (m *MockPageTable) Update(arg0 vm.Page) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Update", arg0)
}

// Update indicates an expected call of Update.
func (mr *MockPageTableMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPageTable)(nil).Update), arg0)
}
