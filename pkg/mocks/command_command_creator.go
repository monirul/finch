// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runfinch/finch/pkg/command (interfaces: Creator)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	command "github.com/runfinch/finch/pkg/command"
)

// CommandCreator is a mock of Creator interface.
type CommandCreator struct {
	ctrl     *gomock.Controller
	recorder *CommandCreatorMockRecorder
}

// CommandCreatorMockRecorder is the mock recorder for CommandCreator.
type CommandCreatorMockRecorder struct {
	mock *CommandCreator
}

// NewCommandCreator creates a new mock instance.
func NewCommandCreator(ctrl *gomock.Controller) *CommandCreator {
	mock := &CommandCreator{ctrl: ctrl}
	mock.recorder = &CommandCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *CommandCreator) EXPECT() *CommandCreatorMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *CommandCreator) Create(arg0 string, arg1 ...string) command.Command {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(command.Command)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *CommandCreatorMockRecorder) Create(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*CommandCreator)(nil).Create), varargs...)
}