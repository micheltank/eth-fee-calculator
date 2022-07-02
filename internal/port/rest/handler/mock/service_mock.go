// Code generated by MockGen. DO NOT EDIT.
// Source: transaction_handler.go

// Package mock_handler is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/micheltank/eth-fee-calculator/internal/domain"
)

// MockTransactionService is a mock of TransactionService interface.
type MockTransactionService struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionServiceMockRecorder
}

// MockTransactionServiceMockRecorder is the mock recorder for MockTransactionService.
type MockTransactionServiceMockRecorder struct {
	mock *MockTransactionService
}

// NewMockTransactionService creates a new mock instance.
func NewMockTransactionService(ctrl *gomock.Controller) *MockTransactionService {
	mock := &MockTransactionService{ctrl: ctrl}
	mock.recorder = &MockTransactionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionService) EXPECT() *MockTransactionServiceMockRecorder {
	return m.recorder
}

// GetTransactionsPerHour mocks base method.
func (m *MockTransactionService) GetTransactionsPerHour(from, to int64) ([]domain.TransactionCostPerHour, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionsPerHour", from, to)
	ret0, _ := ret[0].([]domain.TransactionCostPerHour)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionsPerHour indicates an expected call of GetTransactionsPerHour.
func (mr *MockTransactionServiceMockRecorder) GetTransactionsPerHour(from, to interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionsPerHour", reflect.TypeOf((*MockTransactionService)(nil).GetTransactionsPerHour), from, to)
}