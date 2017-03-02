// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/amazon-ecs-cni-plugins/pkg/oswrapper (interfaces: OS,OSProcess)

package mock_oswrapper

import (
	oswrapper "github.com/aws/amazon-ecs-cni-plugins/pkg/oswrapper"
	gomock "github.com/golang/mock/gomock"
)

// Mock of OS interface
type MockOS struct {
	ctrl     *gomock.Controller
	recorder *_MockOSRecorder
}

// Recorder for MockOS (not exported)
type _MockOSRecorder struct {
	mock *MockOS
}

func NewMockOS(ctrl *gomock.Controller) *MockOS {
	mock := &MockOS{ctrl: ctrl}
	mock.recorder = &_MockOSRecorder{mock}
	return mock
}

func (_m *MockOS) EXPECT() *_MockOSRecorder {
	return _m.recorder
}

func (_m *MockOS) FindProcess(_param0 int) (oswrapper.OSProcess, error) {
	ret := _m.ctrl.Call(_m, "FindProcess", _param0)
	ret0, _ := ret[0].(oswrapper.OSProcess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockOSRecorder) FindProcess(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindProcess", arg0)
}

// Mock of OSProcess interface
type MockOSProcess struct {
	ctrl     *gomock.Controller
	recorder *_MockOSProcessRecorder
}

// Recorder for MockOSProcess (not exported)
type _MockOSProcessRecorder struct {
	mock *MockOSProcess
}

func NewMockOSProcess(ctrl *gomock.Controller) *MockOSProcess {
	mock := &MockOSProcess{ctrl: ctrl}
	mock.recorder = &_MockOSProcessRecorder{mock}
	return mock
}

func (_m *MockOSProcess) EXPECT() *_MockOSProcessRecorder {
	return _m.recorder
}

func (_m *MockOSProcess) Kill() error {
	ret := _m.ctrl.Call(_m, "Kill")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockOSProcessRecorder) Kill() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Kill")
}
