// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory-am/fosite (interfaces: AccessResponder)

package internal

import (
	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory-am/fosite"
	time "time"
)

// Mock of AccessResponder interface
type MockAccessResponder struct {
	ctrl     *gomock.Controller
	recorder *_MockAccessResponderRecorder
}

// Recorder for MockAccessResponder (not exported)
type _MockAccessResponderRecorder struct {
	mock *MockAccessResponder
}

func NewMockAccessResponder(ctrl *gomock.Controller) *MockAccessResponder {
	mock := &MockAccessResponder{ctrl: ctrl}
	mock.recorder = &_MockAccessResponderRecorder{mock}
	return mock
}

func (_m *MockAccessResponder) EXPECT() *_MockAccessResponderRecorder {
	return _m.recorder
}

func (_m *MockAccessResponder) GetAccessToken() string {
	ret := _m.ctrl.Call(_m, "GetAccessToken")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockAccessResponderRecorder) GetAccessToken() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAccessToken")
}

func (_m *MockAccessResponder) GetExtra(_param0 string) interface{} {
	ret := _m.ctrl.Call(_m, "GetExtra", _param0)
	ret0, _ := ret[0].(interface{})
	return ret0
}

func (_mr *_MockAccessResponderRecorder) GetExtra(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetExtra", arg0)
}

func (_m *MockAccessResponder) GetTokenType() string {
	ret := _m.ctrl.Call(_m, "GetTokenType")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockAccessResponderRecorder) GetTokenType() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTokenType")
}

func (_m *MockAccessResponder) SetAccessToken(_param0 string) {
	_m.ctrl.Call(_m, "SetAccessToken", _param0)
}

func (_mr *_MockAccessResponderRecorder) SetAccessToken(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetAccessToken", arg0)
}

func (_m *MockAccessResponder) SetExpiresIn(_param0 time.Duration) {
	_m.ctrl.Call(_m, "SetExpiresIn", _param0)
}

func (_mr *_MockAccessResponderRecorder) SetExpiresIn(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetExpiresIn", arg0)
}

func (_m *MockAccessResponder) SetExtra(_param0 string, _param1 interface{}) {
	_m.ctrl.Call(_m, "SetExtra", _param0, _param1)
}

func (_mr *_MockAccessResponderRecorder) SetExtra(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetExtra", arg0, arg1)
}

func (_m *MockAccessResponder) SetScopes(_param0 fosite.Arguments) {
	_m.ctrl.Call(_m, "SetScopes", _param0)
}

func (_mr *_MockAccessResponderRecorder) SetScopes(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetScopes", arg0)
}

func (_m *MockAccessResponder) SetTokenType(_param0 string) {
	_m.ctrl.Call(_m, "SetTokenType", _param0)
}

func (_mr *_MockAccessResponderRecorder) SetTokenType(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetTokenType", arg0)
}

func (_m *MockAccessResponder) ToMap() map[string]interface{} {
	ret := _m.ctrl.Call(_m, "ToMap")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

func (_mr *_MockAccessResponderRecorder) ToMap() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ToMap")
}
