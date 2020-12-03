package interactors

import (
	"errors"
	"github.com/nikulnik/weather/domain"
	"github.com/nikulnik/weather/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

var city, country = "city", "country"
var forecastDay int64 = 3

func TestGetWeather_WhenWeatherExistsInCacheAndDayNotProvided_ReturnsWeather(t *testing.T) {
	weather := &domain.CurrentWeather{
		City: "city",
	}
	cacheMock := &mocks.Cache{}
	clientMock := &mocks.OpenWeatherMapClient{}
	interactor := NewWeatherInteractor(clientMock, cacheMock)
	cacheMock.On("GetCurrentWeather", city, country).Return(weather)
	currWeather, forecast, err := interactor.GetCurrentWeather(city, country, nil)
	assert.Nil(t, err)
	assert.NotNil(t, currWeather)
	assert.Nil(t, forecast)
}

func TestGetWeather_WhenOpenWeatherClientReturnsErrorAndNothingInCache_ReturnsError(t *testing.T) {
	cacheMock := &mocks.Cache{}
	clientMock := &mocks.OpenWeatherMapClient{}
	interactor := NewWeatherInteractor(clientMock, cacheMock)
	clientMock.On("GetCurrentWeather", city, country).Return(nil, errors.New("some error"))
	cacheMock.On("GetCurrentWeather", city, country).Return(nil)
	currWeather, forecast, err := interactor.GetCurrentWeather(city, country, nil)
	assert.Error(t, err)
	assert.Nil(t, currWeather)
	assert.Nil(t, forecast)
}

func TestGetWeather_WhenForecastExistsInCache_ReturnsForecast(t *testing.T) {
	forecastCache := &domain.Forecast{
		DateTime: 12,
		Humidity: 21,
	}
	weather := &domain.CurrentWeather{
		Lat: "11",
		Lon: "11",
	}
	cacheMock := &mocks.Cache{}
	clientMock := &mocks.OpenWeatherMapClient{}
	interactor := NewWeatherInteractor(clientMock, cacheMock)
	clientMock.On("GetCurrentWeather", city, country).Return(weather, nil)
	cacheMock.On("GetCurrentWeather", city, country).Return(nil)
	cacheMock.On("GetForecast", weather.Lat, weather.Lon).Return(forecastCache)
	cacheMock.On("SetWeather", city, country, weather).Return()
	currWeather, forecast, err := interactor.GetCurrentWeather(city, country, &forecastDay)
	assert.Nil(t, err)
	assert.Equal(t, weather, currWeather)
	assert.Equal(t, forecastCache, forecast)
}

func TestGetWeather_WhenCannotRequestForecast_ReturnsError(t *testing.T) {
	forecast := &domain.Forecast{
		DateTime: 12,
		Humidity: 21,
	}
	weather := &domain.CurrentWeather{
		Lat: "11",
		Lon: "11",
	}
	cacheMock := &mocks.Cache{}
	clientMock := &mocks.OpenWeatherMapClient{}
	interactor := NewWeatherInteractor(clientMock, cacheMock)
	clientMock.On("GetCurrentWeather", city, country).Return(weather, nil)
	cacheMock.On("GetCurrentWeather", city, country).Return(nil)
	cacheMock.On("GetForecast", weather.Lat, weather.Lon).Return(nil)
	cacheMock.On("SetWeather", city, country, weather).Return()
	clientMock.On("GetForecast", weather.Lat, weather.Lon, forecastDay).Return(nil, errors.New("some error"))
	currWeather, forecast, err := interactor.GetCurrentWeather(city, country, &forecastDay)
	assert.Error(t, err)
	assert.Nil(t, currWeather)
	assert.Nil(t, forecast)
}

func TestGetWeather_WhenForecastIsNotInCache_ReturnsForecast(t *testing.T) {
	forecast := &domain.Forecast{
		DateTime: 12,
		Humidity: 21,
	}
	weather := &domain.CurrentWeather{
		Lat: "11",
		Lon: "11",
	}
	cacheMock := &mocks.Cache{}
	clientMock := &mocks.OpenWeatherMapClient{}
	interactor := NewWeatherInteractor(clientMock, cacheMock)
	clientMock.On("GetCurrentWeather", city, country).Return(weather, nil)
	cacheMock.On("GetCurrentWeather", city, country).Return(nil)
	cacheMock.On("GetForecast", weather.Lat, weather.Lon).Return(nil)
	cacheMock.On("SetWeather", city, country, weather).Return()
	cacheMock.On("SetForecast", weather.Lat, weather.Lon, forecast).Return()
	clientMock.On("GetForecast", weather.Lat, weather.Lon, forecastDay).Return(forecast, nil)
	resultWeather, resultForecast, err := interactor.GetCurrentWeather(city, country, &forecastDay)
	assert.Nil(t, err)
	assert.Equal(t, weather, resultWeather)
	assert.Equal(t, forecast, resultForecast)
}
