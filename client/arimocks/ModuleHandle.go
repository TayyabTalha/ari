package arimocks

import ari "github.com/CyCoreSystems/ari"
import mock "github.com/stretchr/testify/mock"

// ModuleHandle is an autogenerated mock type for the ModuleHandle type
type ModuleHandle struct {
	mock.Mock
}

// Data provides a mock function with given fields:
func (_m *ModuleHandle) Data() (*ari.ModuleData, error) {
	ret := _m.Called()

	var r0 *ari.ModuleData
	if rf, ok := ret.Get(0).(func() *ari.ModuleData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.ModuleData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ID provides a mock function with given fields:
func (_m *ModuleHandle) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Load provides a mock function with given fields:
func (_m *ModuleHandle) Load() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reload provides a mock function with given fields:
func (_m *ModuleHandle) Reload() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unload provides a mock function with given fields:
func (_m *ModuleHandle) Unload() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
