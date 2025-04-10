// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gosuit/utils/coder (interfaces: Coder)
//
// Generated by this command:
//
//      mockgen github.com/gosuit/utils/coder Coder
//

// Package mock_coder is a generated GoMock package.
package mock_coder

import (
        reflect "reflect"

        gomock "go.uber.org/mock/gomock"
)

// MockCoder is a mock of Coder interface.
type MockCoder struct {
        ctrl     *gomock.Controller
        recorder *MockCoderMockRecorder
        isgomock struct{}
}

// MockCoderMockRecorder is the mock recorder for MockCoder.
type MockCoderMockRecorder struct {
        mock *MockCoder
}

// NewMockCoder creates a new mock instance.
func NewMockCoder(ctrl *gomock.Controller) *MockCoder {
        mock := &MockCoder{ctrl: ctrl}
        mock.recorder = &MockCoderMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCoder) EXPECT() *MockCoderMockRecorder {
        return m.recorder
}

// CompareHash mocks base method.
func (m *MockCoder) CompareHash(hash, text string) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CompareHash", hash, text)
        ret0, _ := ret[0].(error)
        return ret0
}

// CompareHash indicates an expected call of CompareHash.
func (mr *MockCoderMockRecorder) CompareHash(hash, text any) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareHash", reflect.TypeOf((*MockCoder)(nil).CompareHash), hash, text)
}

// Decrypt mocks base method.
func (m *MockCoder) Decrypt(text string) (string, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Decrypt", text)
        ret0, _ := ret[0].(string)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Decrypt indicates an expected call of Decrypt.
func (mr *MockCoderMockRecorder) Decrypt(text any) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrypt", reflect.TypeOf((*MockCoder)(nil).Decrypt), text)
}

// Encrypt mocks base method.
func (m *MockCoder) Encrypt(text string) (string, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Encrypt", text)
        ret0, _ := ret[0].(string)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Encrypt indicates an expected call of Encrypt.
func (mr *MockCoderMockRecorder) Encrypt(text any) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockCoder)(nil).Encrypt), text)
}

// Hash mocks base method.
func (m *MockCoder) Hash(text string) (string, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Hash", text)
        ret0, _ := ret[0].(string)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Hash indicates an expected call of Hash.
func (mr *MockCoderMockRecorder) Hash(text any) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockCoder)(nil).Hash), text)
}