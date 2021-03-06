// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/nikulnik/weather/domain"

	mock "github.com/stretchr/testify/mock"
)

// WeatherInteractor is an autogenerated mock type for the WeatherInteractor type
type WeatherInteractor struct {
	mock.Mock
}

// GetCurrentWeather provides a mock function with given fields: city, country, forecastDay
func (_m *WeatherInteractor) GetCurrentWeather(city string, country string, forecastDay *int64) (*domain.CurrentWeather, *domain.Forecast, error) {
	ret := _m.Called(city, country, forecastDay)

	var r0 *domain.CurrentWeather
	if rf, ok := ret.Get(0).(func(string, string, *int64) *domain.CurrentWeather); ok {
		r0 = rf(city, country, forecastDay)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.CurrentWeather)
		}
	}

	var r1 *domain.Forecast
	if rf, ok := ret.Get(1).(func(string, string, *int64) *domain.Forecast); ok {
		r1 = rf(city, country, forecastDay)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.Forecast)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, *int64) error); ok {
		r2 = rf(city, country, forecastDay)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
