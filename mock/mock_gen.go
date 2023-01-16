// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lotusirous/codescan/core (interfaces: RepositoryStore,ScanScheduler,ScanResultStore,ScanStore)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	core "github.com/lotusirous/codescan/core"
)

// MockRepositoryStore is a mock of RepositoryStore interface.
type MockRepositoryStore struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryStoreMockRecorder
}

// MockRepositoryStoreMockRecorder is the mock recorder for MockRepositoryStore.
type MockRepositoryStoreMockRecorder struct {
	mock *MockRepositoryStore
}

// NewMockRepositoryStore creates a new mock instance.
func NewMockRepositoryStore(ctrl *gomock.Controller) *MockRepositoryStore {
	mock := &MockRepositoryStore{ctrl: ctrl}
	mock.recorder = &MockRepositoryStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryStore) EXPECT() *MockRepositoryStoreMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockRepositoryStore) Count(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockRepositoryStoreMockRecorder) Count(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockRepositoryStore)(nil).Count), arg0)
}

// Create mocks base method.
func (m *MockRepositoryStore) Create(arg0 context.Context, arg1 *core.Repository) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepositoryStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockRepositoryStore) Delete(arg0 context.Context, arg1 *core.Repository) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepositoryStore)(nil).Delete), arg0, arg1)
}

// Find mocks base method.
func (m *MockRepositoryStore) Find(arg0 context.Context, arg1 int64) (*core.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*core.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockRepositoryStoreMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRepositoryStore)(nil).Find), arg0, arg1)
}

// List mocks base method.
func (m *MockRepositoryStore) List(arg0 context.Context) ([]*core.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]*core.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRepositoryStoreMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepositoryStore)(nil).List), arg0)
}

// Update mocks base method.
func (m *MockRepositoryStore) Update(arg0 context.Context, arg1 *core.Repository) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepositoryStore)(nil).Update), arg0, arg1)
}

// MockScanScheduler is a mock of ScanScheduler interface.
type MockScanScheduler struct {
	ctrl     *gomock.Controller
	recorder *MockScanSchedulerMockRecorder
}

// MockScanSchedulerMockRecorder is the mock recorder for MockScanScheduler.
type MockScanSchedulerMockRecorder struct {
	mock *MockScanScheduler
}

// NewMockScanScheduler creates a new mock instance.
func NewMockScanScheduler(ctrl *gomock.Controller) *MockScanScheduler {
	mock := &MockScanScheduler{ctrl: ctrl}
	mock.recorder = &MockScanSchedulerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScanScheduler) EXPECT() *MockScanSchedulerMockRecorder {
	return m.recorder
}

// RestoreLastScan mocks base method.
func (m *MockScanScheduler) RestoreLastScan(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreLastScan", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestoreLastScan indicates an expected call of RestoreLastScan.
func (mr *MockScanSchedulerMockRecorder) RestoreLastScan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreLastScan", reflect.TypeOf((*MockScanScheduler)(nil).RestoreLastScan), arg0)
}

// ScanRepo mocks base method.
func (m *MockScanScheduler) ScanRepo(arg0 context.Context, arg1 *core.Repository) (*core.Scan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanRepo", arg0, arg1)
	ret0, _ := ret[0].(*core.Scan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScanRepo indicates an expected call of ScanRepo.
func (mr *MockScanSchedulerMockRecorder) ScanRepo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanRepo", reflect.TypeOf((*MockScanScheduler)(nil).ScanRepo), arg0, arg1)
}

// Start mocks base method.
func (m *MockScanScheduler) Start(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockScanSchedulerMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockScanScheduler)(nil).Start), arg0)
}

// MockScanResultStore is a mock of ScanResultStore interface.
type MockScanResultStore struct {
	ctrl     *gomock.Controller
	recorder *MockScanResultStoreMockRecorder
}

// MockScanResultStoreMockRecorder is the mock recorder for MockScanResultStore.
type MockScanResultStoreMockRecorder struct {
	mock *MockScanResultStore
}

// NewMockScanResultStore creates a new mock instance.
func NewMockScanResultStore(ctrl *gomock.Controller) *MockScanResultStore {
	mock := &MockScanResultStore{ctrl: ctrl}
	mock.recorder = &MockScanResultStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScanResultStore) EXPECT() *MockScanResultStoreMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockScanResultStore) Count(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockScanResultStoreMockRecorder) Count(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockScanResultStore)(nil).Count), arg0)
}

// Create mocks base method.
func (m *MockScanResultStore) Create(arg0 context.Context, arg1 *core.ScanResult) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockScanResultStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockScanResultStore)(nil).Create), arg0, arg1)
}

// Find mocks base method.
func (m *MockScanResultStore) Find(arg0 context.Context, arg1 int64) (*core.ScanResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*core.ScanResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockScanResultStoreMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockScanResultStore)(nil).Find), arg0, arg1)
}

// List mocks base method.
func (m *MockScanResultStore) List(arg0 context.Context) ([]*core.ScanResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]*core.ScanResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockScanResultStoreMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockScanResultStore)(nil).List), arg0)
}

// MockScanStore is a mock of ScanStore interface.
type MockScanStore struct {
	ctrl     *gomock.Controller
	recorder *MockScanStoreMockRecorder
}

// MockScanStoreMockRecorder is the mock recorder for MockScanStore.
type MockScanStoreMockRecorder struct {
	mock *MockScanStore
}

// NewMockScanStore creates a new mock instance.
func NewMockScanStore(ctrl *gomock.Controller) *MockScanStore {
	mock := &MockScanStore{ctrl: ctrl}
	mock.recorder = &MockScanStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScanStore) EXPECT() *MockScanStoreMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockScanStore) Count(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockScanStoreMockRecorder) Count(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockScanStore)(nil).Count), arg0)
}

// Create mocks base method.
func (m *MockScanStore) Create(arg0 context.Context, arg1 *core.Scan) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockScanStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockScanStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockScanStore) Delete(arg0 context.Context, arg1 *core.Scan) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockScanStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockScanStore)(nil).Delete), arg0, arg1)
}

// Find mocks base method.
func (m *MockScanStore) Find(arg0 context.Context, arg1 int64) (*core.Scan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*core.Scan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockScanStoreMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockScanStore)(nil).Find), arg0, arg1)
}

// FindEnqueued mocks base method.
func (m *MockScanStore) FindEnqueued(arg0 context.Context) ([]*core.Scan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEnqueued", arg0)
	ret0, _ := ret[0].([]*core.Scan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEnqueued indicates an expected call of FindEnqueued.
func (mr *MockScanStoreMockRecorder) FindEnqueued(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEnqueued", reflect.TypeOf((*MockScanStore)(nil).FindEnqueued), arg0)
}

// List mocks base method.
func (m *MockScanStore) List(arg0 context.Context) ([]*core.Scan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]*core.Scan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockScanStoreMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockScanStore)(nil).List), arg0)
}

// Update mocks base method.
func (m *MockScanStore) Update(arg0 context.Context, arg1 *core.Scan) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockScanStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockScanStore)(nil).Update), arg0, arg1)
}
