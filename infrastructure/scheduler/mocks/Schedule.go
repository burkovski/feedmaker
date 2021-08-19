// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Schedule is an autogenerated mock type for the Schedule type
type Schedule struct {
	mock.Mock
}

// FireInterval provides a mock function with given fields:
func (_m *Schedule) FireInterval() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// Next provides a mock function with given fields: _a0
func (_m *Schedule) Next(_a0 time.Time) time.Time {
	ret := _m.Called(_a0)

	var r0 time.Time
	if rf, ok := ret.Get(0).(func(time.Time) time.Time); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// StartTimestamp provides a mock function with given fields:
func (_m *Schedule) StartTimestamp() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}