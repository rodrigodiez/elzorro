// Code generated by mockery v1.0.0. DO NOT EDIT.
package datastore

import datastore "cloud.google.com/go/datastore"
import mock "github.com/stretchr/testify/mock"

// Transaction is an autogenerated mock type for the Transaction type
type Transaction struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *Transaction) Get(_a0 *datastore.Key, _a1 interface{}) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datastore.Key, interface{}) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Put provides a mock function with given fields: _a0, _a1
func (_m *Transaction) Put(_a0 *datastore.Key, _a1 interface{}) (*datastore.PendingKey, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *datastore.PendingKey
	if rf, ok := ret.Get(0).(func(*datastore.Key, interface{}) *datastore.PendingKey); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.PendingKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*datastore.Key, interface{}) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
