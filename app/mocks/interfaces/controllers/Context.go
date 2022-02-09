// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Context is an autogenerated mock type for the Context type
type Context struct {
	mock.Mock
}

// AbortWithStatus provides a mock function with given fields: code
func (_m *Context) AbortWithStatus(code int) {
	_m.Called(code)
}

// Get provides a mock function with given fields: key
func (_m *Context) Get(key string) (interface{}, bool) {
	ret := _m.Called(key)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// JSON provides a mock function with given fields: code, obj
func (_m *Context) JSON(code int, obj interface{}) {
	_m.Called(code, obj)
}

// MustGet provides a mock function with given fields: key
func (_m *Context) MustGet(key string) interface{} {
	ret := _m.Called(key)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Param provides a mock function with given fields: key
func (_m *Context) Param(key string) string {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PostForm provides a mock function with given fields: key
func (_m *Context) PostForm(key string) string {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
