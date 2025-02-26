// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	core "github.com/LambdaTest/test-at-scale/pkg/core"

	mock "github.com/stretchr/testify/mock"
)

// ExecutionManager is an autogenerated mock type for the ExecutionManager type
type ExecutionManager struct {
	mock.Mock
}

// ExecuteInternalCommands provides a mock function with given fields: ctx, commandType, commands, cwd, envMap, secretData
func (_m *ExecutionManager) ExecuteInternalCommands(ctx context.Context, commandType core.CommandType, commands []string, cwd string, envMap map[string]string, secretData map[string]string) error {
	ret := _m.Called(ctx, commandType, commands, cwd, envMap, secretData)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, core.CommandType, []string, string, map[string]string, map[string]string) error); ok {
		r0 = rf(ctx, commandType, commands, cwd, envMap, secretData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExecuteUserCommands provides a mock function with given fields: ctx, commandType, payload, runConfig, secretData
func (_m *ExecutionManager) ExecuteUserCommands(ctx context.Context, commandType core.CommandType, payload *core.Payload, runConfig *core.Run, secretData map[string]string) error {
	ret := _m.Called(ctx, commandType, payload, runConfig, secretData)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, core.CommandType, *core.Payload, *core.Run, map[string]string) error); ok {
		r0 = rf(ctx, commandType, payload, runConfig, secretData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetEnvVariables provides a mock function with given fields: envMap, secretData
func (_m *ExecutionManager) GetEnvVariables(envMap map[string]string, secretData map[string]string) ([]string, error) {
	ret := _m.Called(envMap, secretData)

	var r0 []string
	if rf, ok := ret.Get(0).(func(map[string]string, map[string]string) []string); ok {
		r0 = rf(envMap, secretData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]string, map[string]string) error); ok {
		r1 = rf(envMap, secretData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreCommandLogs provides a mock function with given fields: ctx, blobPath, reader
func (_m *ExecutionManager) StoreCommandLogs(ctx context.Context, blobPath string, reader io.Reader) <-chan error {
	ret := _m.Called(ctx, blobPath, reader)

	var r0 <-chan error
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader) <-chan error); ok {
		r0 = rf(ctx, blobPath, reader)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan error)
		}
	}

	return r0
}
