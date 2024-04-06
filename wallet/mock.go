// Code generated by MockGen. DO NOT EDIT.
// Source: ./wallet/interface.go
//
// Generated by this command:
//
//	mockgen -source=./wallet/interface.go -destination=./wallet/mock.go -package=wallet
//

// Package wallet is a generated GoMock package.
package wallet

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIWallet is a mock of IWallet interface.
type MockIWallet struct {
	ctrl     *gomock.Controller
	recorder *MockIWalletMockRecorder
}

// MockIWalletMockRecorder is the mock recorder for MockIWallet.
type MockIWalletMockRecorder struct {
	mock *MockIWallet
}

// NewMockIWallet creates a new mock instance.
func NewMockIWallet(ctrl *gomock.Controller) *MockIWallet {
	mock := &MockIWallet{ctrl: ctrl}
	mock.recorder = &MockIWalletMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIWallet) EXPECT() *MockIWalletMockRecorder {
	return m.recorder
}

// Address mocks base method.
func (m *MockIWallet) Address() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(string)
	return ret0
}

// Address indicates an expected call of Address.
func (mr *MockIWalletMockRecorder) Address() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockIWallet)(nil).Address))
}

// Balance mocks base method.
func (m *MockIWallet) Balance() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Balance")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Balance indicates an expected call of Balance.
func (mr *MockIWalletMockRecorder) Balance() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Balance", reflect.TypeOf((*MockIWallet)(nil).Balance))
}

// BondTransaction mocks base method.
func (m *MockIWallet) BondTransaction(arg0, arg1, arg2 string, arg3 int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BondTransaction", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BondTransaction indicates an expected call of BondTransaction.
func (mr *MockIWalletMockRecorder) BondTransaction(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BondTransaction", reflect.TypeOf((*MockIWallet)(nil).BondTransaction), arg0, arg1, arg2, arg3)
}

// NewAddress mocks base method.
func (m *MockIWallet) NewAddress(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAddress", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewAddress indicates an expected call of NewAddress.
func (mr *MockIWalletMockRecorder) NewAddress(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAddress", reflect.TypeOf((*MockIWallet)(nil).NewAddress), arg0)
}

// TransferTransaction mocks base method.
func (m *MockIWallet) TransferTransaction(arg0, arg1 string, arg2 int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferTransaction", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferTransaction indicates an expected call of TransferTransaction.
func (mr *MockIWalletMockRecorder) TransferTransaction(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTransaction", reflect.TypeOf((*MockIWallet)(nil).TransferTransaction), arg0, arg1, arg2)
}
