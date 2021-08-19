// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	gateway "go-feedmaker/infrastructure/gateway"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Dialer is an autogenerated mock type for the Dialer type
type Dialer struct {
	mock.Mock
}

// DialTimeout provides a mock function with given fields: addr, timeout
func (_m *Dialer) DialTimeout(addr string, timeout time.Duration) (gateway.FtpConnection, error) {
	ret := _m.Called(addr, timeout)

	var r0 gateway.FtpConnection
	if rf, ok := ret.Get(0).(func(string, time.Duration) gateway.FtpConnection); ok {
		r0 = rf(addr, timeout)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gateway.FtpConnection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, time.Duration) error); ok {
		r1 = rf(addr, timeout)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}