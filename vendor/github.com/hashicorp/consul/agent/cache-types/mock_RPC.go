// Code generated by mockery v1.0.0
package cachetype

import mock "github.com/stretchr/testify/mock"

// MockRPC is an autogenerated mock type for the RPC type
type MockRPC struct {
	mock.Mock
}

// RPC provides a mock function with given fields: method, args, reply
func (_m *MockRPC) RPC(method string, args interface{}, reply interface{}) error {
	ret := _m.Called(method, args, reply)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}, interface{}) error); ok {
		r0 = rf(method, args, reply)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
