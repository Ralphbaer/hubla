// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Ralphbaer/hubla/backend/transaction/repository (interfaces: TransactionRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entity "github.com/Ralphbaer/hubla/backend/transaction/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockTransactionRepository is a mock of TransactionRepository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// ListTransactionsByFileID mocks base method.
func (m *MockTransactionRepository) ListTransactionsByFileID(arg0 context.Context, arg1 string) ([]*entity.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTransactionsByFileID", arg0, arg1)
	ret0, _ := ret[0].([]*entity.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTransactionsByFileID indicates an expected call of ListTransactionsByFileID.
func (mr *MockTransactionRepositoryMockRecorder) ListTransactionsByFileID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTransactionsByFileID", reflect.TypeOf((*MockTransactionRepository)(nil).ListTransactionsByFileID), arg0, arg1)
}

// Save mocks base method.
func (m *MockTransactionRepository) Save(arg0 context.Context, arg1 *entity.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockTransactionRepositoryMockRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockTransactionRepository)(nil).Save), arg0, arg1)
}
