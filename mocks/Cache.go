// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/nikulnik/weather/domain"
	mock "github.com/stretchr/testify/mock"
)

// Cache is an autogenerated mock type for the Cache type
type Cache struct {
	mock.Mock
}

// GetCurrentWeather provides a mock function with given fields: city, country
func (_m *Cache) GetCurrentWeather(city string, country string) *domain.CurrentWeather {
	ret := _m.Called(city, country)

	var r0 *domain.CurrentWeather
	if rf, ok := ret.Get(0).(func(string, string) *domain.CurrentWeather); ok {
		r0 = rf(city, country)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.CurrentWeather)
		}
	}

	return r0
}

// GetForecast provides a mock function with given fields: lat, lon
func (_m *Cache) GetForecast(lat string, lon string) *domain.Forecast {
	ret := _m.Called(lat, lon)

	var r0 *domain.Forecast
	if rf, ok := ret.Get(0).(func(string, string) *domain.Forecast); ok {
		r0 = rf(lat, lon)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Forecast)
		}
	}

	return r0
}

// SetForecast provides a mock function with given fields: lat, lon, value
func (_m *Cache) SetForecast(lat string, lon string, value *domain.Forecast) {
	_m.Called(lat, lon, value)
}

// SetWeather provides a mock function with given fields: city, country, value
func (_m *Cache) SetWeather(city string, country string, value *domain.CurrentWeather) {
	_m.Called(city, country, value)
}
