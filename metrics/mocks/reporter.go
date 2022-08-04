// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/topfreegames/pitaya/v2/metrics (interfaces: Reporter)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockReporter is a mock of Reporter interface
type MockReporter struct {
	ctrl     *gomock.Controller
	recorder *MockReporterMockRecorder
}

// MockReporterMockRecorder is the mock recorder for MockReporter
type MockReporterMockRecorder struct {
	mock *MockReporter
}

// NewMockReporter creates a new mock instance
func NewMockReporter(ctrl *gomock.Controller) *MockReporter {
	mock := &MockReporter{ctrl: ctrl}
	mock.recorder = &MockReporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReporter) EXPECT() *MockReporterMockRecorder {
	return m.recorder
}

// ReportCount mocks base method
func (m *MockReporter) ReportCount(arg0 string, arg1 map[string]string, arg2 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReportCount", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReportCount indicates an expected call of ReportCount
func (mr *MockReporterMockRecorder) ReportCount(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportCount", reflect.TypeOf((*MockReporter)(nil).ReportCount), arg0, arg1, arg2)
}

// ReportGauge mocks base method
func (m *MockReporter) ReportGauge(arg0 string, arg1 map[string]string, arg2 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReportGauge", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReportGauge indicates an expected call of ReportGauge
func (mr *MockReporterMockRecorder) ReportGauge(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportGauge", reflect.TypeOf((*MockReporter)(nil).ReportGauge), arg0, arg1, arg2)
}

// ReportHistogram mocks base method
func (m *MockReporter) ReportHistogram(arg0 string, arg1 map[string]string, arg2 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReportHistogram", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReportHistogram indicates an expected call of ReportHistogram
func (mr *MockReporterMockRecorder) ReportHistogram(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportHistogram", reflect.TypeOf((*MockReporter)(nil).ReportHistogram), arg0, arg1, arg2)
}

// ReportSummary mocks base method
func (m *MockReporter) ReportSummary(arg0 string, arg1 map[string]string, arg2 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReportSummary", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReportSummary indicates an expected call of ReportSummary
func (mr *MockReporterMockRecorder) ReportSummary(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportSummary", reflect.TypeOf((*MockReporter)(nil).ReportSummary), arg0, arg1, arg2)
}
