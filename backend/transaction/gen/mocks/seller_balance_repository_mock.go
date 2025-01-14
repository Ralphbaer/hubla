// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Ralphbaer/hubla/backend/transaction/repository (interfaces: SellerBalanceRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entity "github.com/Ralphbaer/hubla/backend/transaction/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockSellerBalanceRepository is a mock of SellerBalanceRepository interface.
type MockSellerBalanceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSellerBalanceRepositoryMockRecorder
}

// MockSellerBalanceRepositoryMockRecorder is the mock recorder for MockSellerBalanceRepository.
type MockSellerBalanceRepositoryMockRecorder struct {
	mock *MockSellerBalanceRepository
}

// NewMockSellerBalanceRepository creates a new mock instance.
func NewMockSellerBalanceRepository(ctrl *gomock.Controller) *MockSellerBalanceRepository {
	mock := &MockSellerBalanceRepository{ctrl: ctrl}
	mock.recorder = &MockSellerBalanceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSellerBalanceRepository) EXPECT() *MockSellerBalanceRepositoryMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockSellerBalanceRepository) Find(arg0 context.Context, arg1 string) (*entity.SellerBalanceView, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*entity.SellerBalanceView)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockSellerBalanceRepositoryMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockSellerBalanceRepository)(nil).Find), arg0, arg1)
}

// Upsert mocks base method.
func (m *MockSellerBalanceRepository) Upsert(arg0 context.Context, arg1 *entity.SellerBalance) (*float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(*float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert.
func (mr *MockSellerBalanceRepositoryMockRecorder) Upsert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockSellerBalanceRepository)(nil).Upsert), arg0, arg1)
}
