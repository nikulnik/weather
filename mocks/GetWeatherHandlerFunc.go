// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	middleware "github.com/go-openapi/runtime/middleware"
	mock "github.com/stretchr/testify/mock"

	weather "github.com/nikulnik/weather/restapi/operations/weather"
)

// GetWeatherHandlerFunc is an autogenerated mock type for the GetWeatherHandlerFunc type
type GetWeatherHandlerFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *GetWeatherHandlerFunc) Execute(_a0 weather.GetWeatherParams) middleware.Responder {
	ret := _m.Called(_a0)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(weather.GetWeatherParams) middleware.Responder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}