// Code generated by mockery v1.0.0. DO NOT EDIT.
package boltdb

import boltdb "github.com/rodrigodiez/zorro/pkg/storage/boltdb"
import mock "github.com/stretchr/testify/mock"

// ClientAdapter is an autogenerated mock type for the ClientAdapter type
type ClientAdapter struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *ClientAdapter) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0
func (_m *ClientAdapter) Update(_a0 func(boltdb.Transaction) error) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(boltdb.Transaction) error) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// View provides a mock function with given fields: _a0
func (_m *ClientAdapter) View(_a0 func(boltdb.Transaction) error) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(boltdb.Transaction) error) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}