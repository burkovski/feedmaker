// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Connection is an autogenerated mock type for the Connection type
type Connection struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Connection) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Do provides a mock function with given fields: commandName, args
func (_m *Connection) Do(commandName string, args ...interface{}) (interface{}, error) {
	var _ca []interface{}
	_ca = append(_ca, commandName)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, ...interface{}) interface{}); ok {
		r0 = rf(commandName, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(commandName, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Err provides a mock function with given fields:
func (_m *Connection) Err() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Flush provides a mock function with given fields:
func (_m *Connection) Flush() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Receive provides a mock function with given fields:
func (_m *Connection) Receive() (interface{}, error) {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
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

// Send provides a mock function with given fields: commandName, args
func (_m *Connection) Send(commandName string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, commandName)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) error); ok {
		r0 = rf(commandName, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
