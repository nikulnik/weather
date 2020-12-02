package handlers

import (
	"errors"
	"github.com/nikulnik/weather/domain"
	"github.com/nikulnik/weather/mocks"
	"github.com/nikulnik/weather/models"
	"github.com/nikulnik/weather/restapi/operations/weather"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWeatherHandler_ReturnsHandler(t *testing.T) {
	interactorMock := &mocks.WeatherInteractor{}
	assert.NotNil(t, NewWeatherHandler(interactorMock))
}

func TestGetWeather_WhenGetWeatherInteractorReturnsError_ReturnsDefaultResponder(t *testing.T) {
	params := weather.GetWeatherParams{
		City: "city",
		Country: "country",
		ForecastDay: (*int64)(nil),
	}
	interactorMock := &mocks.WeatherInteractor{}
	handler := NewWeatherHandler(interactorMock)

	expectedError := errors.New("some error")

	interactorMock.On("GetWeather", params.City, params.Country, params.ForecastDay).Return(nil, expectedError)
	r := handler.GetWeather(params)
	resp, ok := r.(*weather.GetWeatherDefault)
	assert.True(t, ok)
	assert.Equal(t, expectedError.Error(), resp.Payload.Error)
}

func TestGetWeather_HappyPath(t *testing.T) {
	params := weather.GetWeatherParams{
		City: "city",
		Country: "country",
		ForecastDay: (*int64)(nil),
	}
	interactorMock := &mocks.WeatherInteractor{}
	handler := NewWeatherHandler(interactorMock)

	expectedResponse := &domain.WeatherWithForecast{
		Sunset: "6:54",
		Humidity: "32%",
		Forecast: &domain.Forecast{
			Date:        "2011-11-11",
		},
	}

	interactorMock.On("GetWeather", params.City, params.Country, params.ForecastDay).Return(expectedResponse, nil)
	r := handler.GetWeather(params)
	resp, ok := r.(*weather.GetWeatherOK)
	assert.True(t, ok)
	assert.Equal(t, toResponseWeather(expectedResponse), resp.Payload)
}

func TestToResponseWeather_ReturnsResponse(t *testing.T) {
	domainWeather := &domain.WeatherWithForecast{
		Cloudiness:     "1",
		Forecast:      &domain.Forecast{
			Date:        "1",
			Humidity:    "1",
			Pressure:    "1",
			Sunrise:     "1",
			Sunset:      "1",
			Temperature: "1",
			Wind:        "1",
		},
		GeoCoordinates: "1",
		Humidity:       "1",
		LocationName:   "1",
		Pressure:       "1",
		RequestedTime:  "1",
		Sunrise:        "1",
		Sunset:         "1",
		Temperature:    "1",
		Wind:           "1",
		Lat:            1,
		Lon:            1,
	}
	expected :=  &models.WeatherWithForecast{
		Cloudiness:     "1",
		Forecast:      &models.Forecast{
			Date:        "1",
			Humidity:    "1",
			Pressure:    "1",
			Sunrise:     "1",
			Sunset:      "1",
			Temperature: "1",
			Wind:        "1",
		},
		GeoCoordinates: "1",
		Humidity:       "1",
		LocationName:   "1",
		Pressure:       "1",
		RequestedTime:  "1",
		Sunrise:        "1",
		Sunset:         "1",
		Temperature:    "1",
		Wind:           "1",
	}
	respWeather := toResponseWeather(domainWeather)
	assert.Equal(t, expected, respWeather)
}