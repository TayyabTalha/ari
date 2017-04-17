package arimocks

import ari "github.com/CyCoreSystems/ari"
import mock "github.com/stretchr/testify/mock"

// Modules is an autogenerated mock type for the Modules type
type Modules struct {
	mock.Mock
}

// Data provides a mock function with given fields: name
func (_m *Modules) Data(name string) (*ari.ModuleData, error) {
	ret := _m.Called(name)

	var r0 *ari.ModuleData
	if rf, ok := ret.Get(0).(func(string) *ari.ModuleData); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.ModuleData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: name
func (_m *Modules) Get(name string) ari.ModuleHandle {
	ret := _m.Called(name)

	var r0 ari.ModuleHandle
	if rf, ok := ret.Get(0).(func(string) ari.ModuleHandle); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.ModuleHandle)
		}
	}

	return r0
}

// List provides a mock function with given fields:
func (_m *Modules) List() ([]ari.ModuleHandle, error) {
	ret := _m.Called()

	var r0 []ari.ModuleHandle
	if rf, ok := ret.Get(0).(func() []ari.ModuleHandle); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ari.ModuleHandle)
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

// Load provides a mock function with given fields: name
func (_m *Modules) Load(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reload provides a mock function with given fields: name
func (_m *Modules) Reload(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unload provides a mock function with given fields: name
func (_m *Modules) Unload(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
