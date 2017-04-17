package arimocks

import ari "github.com/CyCoreSystems/ari"
import mock "github.com/stretchr/testify/mock"

// BridgeHandle is an autogenerated mock type for the BridgeHandle type
type BridgeHandle struct {
	mock.Mock
}

// AddChannel provides a mock function with given fields: channelID
func (_m *BridgeHandle) AddChannel(channelID string) error {
	ret := _m.Called(channelID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(channelID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Data provides a mock function with given fields:
func (_m *BridgeHandle) Data() (*ari.BridgeData, error) {
	ret := _m.Called()

	var r0 *ari.BridgeData
	if rf, ok := ret.Get(0).(func() *ari.BridgeData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.BridgeData)
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

// Delete provides a mock function with given fields:
func (_m *BridgeHandle) Delete() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *BridgeHandle) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Match provides a mock function with given fields: e
func (_m *BridgeHandle) Match(e ari.Event) bool {
	ret := _m.Called(e)

	var r0 bool
	if rf, ok := ret.Get(0).(func(ari.Event) bool); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Play provides a mock function with given fields: id, mediaURI
func (_m *BridgeHandle) Play(id string, mediaURI string) (ari.PlaybackHandle, error) {
	ret := _m.Called(id, mediaURI)

	var r0 ari.PlaybackHandle
	if rf, ok := ret.Get(0).(func(string, string) ari.PlaybackHandle); ok {
		r0 = rf(id, mediaURI)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.PlaybackHandle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, mediaURI)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Record provides a mock function with given fields: name, opts
func (_m *BridgeHandle) Record(name string, opts *ari.RecordingOptions) (ari.LiveRecordingHandle, error) {
	ret := _m.Called(name, opts)

	var r0 ari.LiveRecordingHandle
	if rf, ok := ret.Get(0).(func(string, *ari.RecordingOptions) ari.LiveRecordingHandle); ok {
		r0 = rf(name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.LiveRecordingHandle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *ari.RecordingOptions) error); ok {
		r1 = rf(name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveChannel provides a mock function with given fields: channelID
func (_m *BridgeHandle) RemoveChannel(channelID string) error {
	ret := _m.Called(channelID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(channelID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: n
func (_m *BridgeHandle) Subscribe(n ...string) ari.Subscription {
	_va := make([]interface{}, len(n))
	for _i := range n {
		_va[_i] = n[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ari.Subscription
	if rf, ok := ret.Get(0).(func(...string) ari.Subscription); ok {
		r0 = rf(n...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.Subscription)
		}
	}

	return r0
}
