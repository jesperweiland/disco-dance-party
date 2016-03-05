// Automatically generated by MockGen. DO NOT EDIT!
// Source: hue/hue.go

package mockHue

import (
	message "github.com/drombosky/disco-dance-party/hue/message"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

// Recorder for MockClient (not exported)
type _MockClientRecorder struct {
	mock *MockClient
}

func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

func (_m *MockClient) EXPECT() *_MockClientRecorder {
	return _m.recorder
}

func (_m *MockClient) Do(method string, address string, message []byte, resp interface{}) error {
	ret := _m.ctrl.Call(_m, "Do", method, address, message, resp)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) Do(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Do", arg0, arg1, arg2, arg3)
}

// Mock of Lights interface
type MockLights struct {
	ctrl     *gomock.Controller
	recorder *_MockLightsRecorder
}

// Recorder for MockLights (not exported)
type _MockLightsRecorder struct {
	mock *MockLights
}

func NewMockLights(ctrl *gomock.Controller) *MockLights {
	mock := &MockLights{ctrl: ctrl}
	mock.recorder = &_MockLightsRecorder{mock}
	return mock
}

func (_m *MockLights) EXPECT() *_MockLightsRecorder {
	return _m.recorder
}

func (_m *MockLights) GetAll() (map[string]message.Light, error) {
	ret := _m.ctrl.Call(_m, "GetAll")
	ret0, _ := ret[0].(map[string]message.Light)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLightsRecorder) GetAll() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAll")
}

func (_m *MockLights) GetNew() (*message.GetNewResp, error) {
	ret := _m.ctrl.Call(_m, "GetNew")
	ret0, _ := ret[0].(*message.GetNewResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLightsRecorder) GetNew() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetNew")
}

func (_m *MockLights) Get(id string) (*message.Light, error) {
	ret := _m.ctrl.Call(_m, "Get", id)
	ret0, _ := ret[0].(*message.Light)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLightsRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockLights) Rename(id string, name string) error {
	ret := _m.ctrl.Call(_m, "Rename", id, name)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLightsRecorder) Rename(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Rename", arg0, arg1)
}

func (_m *MockLights) Set(id string, state message.NewLightState) error {
	ret := _m.ctrl.Call(_m, "Set", id, state)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLightsRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Set", arg0, arg1)
}

func (_m *MockLights) Delete(id string) error {
	ret := _m.ctrl.Call(_m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLightsRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0)
}
