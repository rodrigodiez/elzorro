// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import storage "github.com/rodrigodiez/zorro/pkg/storage"

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// LoadOrStore provides a mock function with given fields: key, value
func (_m *Storage) LoadOrStore(key string, value string) (string, bool) {
	ret := _m.Called(key, value)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string, string) bool); ok {
		r1 = rf(key, value)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Resolve provides a mock function with given fields: value
func (_m *Storage) Resolve(value string) (string, bool) {
	ret := _m.Called(value)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(value)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// WithMetrics provides a mock function with given fields: metrics
func (_m *Storage) WithMetrics(metrics *storage.Metrics) storage.Storage {
	ret := _m.Called(metrics)

	var r0 storage.Storage
	if rf, ok := ret.Get(0).(func(*storage.Metrics) storage.Storage); ok {
		r0 = rf(metrics)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storage.Storage)
		}
	}

	return r0
}
