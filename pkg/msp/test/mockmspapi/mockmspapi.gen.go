// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fiorri/fabric-sdk-go/pkg/msp/api (interfaces: CAClient)

// Package mockmspapi is a generated GoMock package.
package mockmspapi

import (
	reflect "reflect"

	api "github.com/fiorri/fabric-sdk-go/pkg/msp/api"
	gomock "github.com/golang/mock/gomock"
)

// MockCAClient is a mock of CAClient interface
type MockCAClient struct {
	ctrl     *gomock.Controller
	recorder *MockCAClientMockRecorder
}

// MockCAClientMockRecorder is the mock recorder for MockCAClient
type MockCAClientMockRecorder struct {
	mock *MockCAClient
}

// NewMockCAClient creates a new mock instance
func NewMockCAClient(ctrl *gomock.Controller) *MockCAClient {
	mock := &MockCAClient{ctrl: ctrl}
	mock.recorder = &MockCAClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCAClient) EXPECT() *MockCAClientMockRecorder {
	return m.recorder
}

// AddAffiliation mocks base method
func (m *MockCAClient) AddAffiliation(arg0 *api.AffiliationRequest) (*api.AffiliationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAffiliation", arg0)
	ret0, _ := ret[0].(*api.AffiliationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAffiliation indicates an expected call of AddAffiliation
func (mr *MockCAClientMockRecorder) AddAffiliation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAffiliation", reflect.TypeOf((*MockCAClient)(nil).AddAffiliation), arg0)
}

// CreateIdentity mocks base method
func (m *MockCAClient) CreateIdentity(arg0 *api.IdentityRequest) (*api.IdentityResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIdentity", arg0)
	ret0, _ := ret[0].(*api.IdentityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIdentity indicates an expected call of CreateIdentity
func (mr *MockCAClientMockRecorder) CreateIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIdentity", reflect.TypeOf((*MockCAClient)(nil).CreateIdentity), arg0)
}

// Enroll mocks base method
func (m *MockCAClient) Enroll(arg0 *api.EnrollmentRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enroll", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enroll indicates an expected call of Enroll
func (mr *MockCAClientMockRecorder) Enroll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enroll", reflect.TypeOf((*MockCAClient)(nil).Enroll), arg0)
}

// GetAffiliation mocks base method
func (m *MockCAClient) GetAffiliation(arg0, arg1 string) (*api.AffiliationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAffiliation", arg0, arg1)
	ret0, _ := ret[0].(*api.AffiliationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAffiliation indicates an expected call of GetAffiliation
func (mr *MockCAClientMockRecorder) GetAffiliation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAffiliation", reflect.TypeOf((*MockCAClient)(nil).GetAffiliation), arg0, arg1)
}

// GetAllAffiliations mocks base method
func (m *MockCAClient) GetAllAffiliations(arg0 string) (*api.AffiliationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAffiliations", arg0)
	ret0, _ := ret[0].(*api.AffiliationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAffiliations indicates an expected call of GetAllAffiliations
func (mr *MockCAClientMockRecorder) GetAllAffiliations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAffiliations", reflect.TypeOf((*MockCAClient)(nil).GetAllAffiliations), arg0)
}

// GetAllIdentities mocks base method
func (m *MockCAClient) GetAllIdentities(arg0 string) ([]*api.IdentityResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllIdentities", arg0)
	ret0, _ := ret[0].([]*api.IdentityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllIdentities indicates an expected call of GetAllIdentities
func (mr *MockCAClientMockRecorder) GetAllIdentities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllIdentities", reflect.TypeOf((*MockCAClient)(nil).GetAllIdentities), arg0)
}

// GetCAInfo mocks base method
func (m *MockCAClient) GetCAInfo() (*api.GetCAInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCAInfo")
	ret0, _ := ret[0].(*api.GetCAInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCAInfo indicates an expected call of GetCAInfo
func (mr *MockCAClientMockRecorder) GetCAInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCAInfo", reflect.TypeOf((*MockCAClient)(nil).GetCAInfo))
}

// GetIdentity mocks base method
func (m *MockCAClient) GetIdentity(arg0, arg1 string) (*api.IdentityResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIdentity", arg0, arg1)
	ret0, _ := ret[0].(*api.IdentityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIdentity indicates an expected call of GetIdentity
func (mr *MockCAClientMockRecorder) GetIdentity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdentity", reflect.TypeOf((*MockCAClient)(nil).GetIdentity), arg0, arg1)
}

// ModifyAffiliation mocks base method
func (m *MockCAClient) ModifyAffiliation(arg0 *api.ModifyAffiliationRequest) (*api.AffiliationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyAffiliation", arg0)
	ret0, _ := ret[0].(*api.AffiliationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyAffiliation indicates an expected call of ModifyAffiliation
func (mr *MockCAClientMockRecorder) ModifyAffiliation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyAffiliation", reflect.TypeOf((*MockCAClient)(nil).ModifyAffiliation), arg0)
}

// ModifyIdentity mocks base method
func (m *MockCAClient) ModifyIdentity(arg0 *api.IdentityRequest) (*api.IdentityResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyIdentity", arg0)
	ret0, _ := ret[0].(*api.IdentityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyIdentity indicates an expected call of ModifyIdentity
func (mr *MockCAClientMockRecorder) ModifyIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyIdentity", reflect.TypeOf((*MockCAClient)(nil).ModifyIdentity), arg0)
}

// Reenroll mocks base method
func (m *MockCAClient) Reenroll(arg0 *api.ReenrollmentRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reenroll", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reenroll indicates an expected call of Reenroll
func (mr *MockCAClientMockRecorder) Reenroll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reenroll", reflect.TypeOf((*MockCAClient)(nil).Reenroll), arg0)
}

// Register mocks base method
func (m *MockCAClient) Register(arg0 *api.RegistrationRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockCAClientMockRecorder) Register(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockCAClient)(nil).Register), arg0)
}

// RemoveAffiliation mocks base method
func (m *MockCAClient) RemoveAffiliation(arg0 *api.AffiliationRequest) (*api.AffiliationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAffiliation", arg0)
	ret0, _ := ret[0].(*api.AffiliationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveAffiliation indicates an expected call of RemoveAffiliation
func (mr *MockCAClientMockRecorder) RemoveAffiliation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAffiliation", reflect.TypeOf((*MockCAClient)(nil).RemoveAffiliation), arg0)
}

// RemoveIdentity mocks base method
func (m *MockCAClient) RemoveIdentity(arg0 *api.RemoveIdentityRequest) (*api.IdentityResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveIdentity", arg0)
	ret0, _ := ret[0].(*api.IdentityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveIdentity indicates an expected call of RemoveIdentity
func (mr *MockCAClientMockRecorder) RemoveIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveIdentity", reflect.TypeOf((*MockCAClient)(nil).RemoveIdentity), arg0)
}

// Revoke mocks base method
func (m *MockCAClient) Revoke(arg0 *api.RevocationRequest) (*api.RevocationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revoke", arg0)
	ret0, _ := ret[0].(*api.RevocationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Revoke indicates an expected call of Revoke
func (mr *MockCAClientMockRecorder) Revoke(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revoke", reflect.TypeOf((*MockCAClient)(nil).Revoke), arg0)
}
