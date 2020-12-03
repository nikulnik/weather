package handlers

import (
	"errors"
	"fmt"
	"github.com/nikulnik/weather/domain"
	"github.com/nikulnik/weather/mocks"
	"github.com/nikulnik/weather/models"
	"github.com/nikulnik/weather/restapi/operations/weather"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewWeatherHandler_ReturnsHandler(t *testing.T) {
	interactorMock := &mocks.WeatherInteractor{}
	assert.NotNil(t, NewWeatherHandler(interactorMock))
}

func TestGetWeather_WhenGetWeatherInteractorReturnsError_ReturnsDefaultResponder(t *testing.T) {
	params := weather.GetWeatherParams{
		City:        "city",
		Country:     "country",
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

func TestGetWeather_HappyPathWithCurrWeatherAndForecast(t *testing.T) {
	params := weather.GetWeatherParams{
		City:        "city",
		Country:     "country",
		ForecastDay: (*int64)(nil),
	}
	interactorMock := &mocks.WeatherInteractor{}
	handler := NewWeatherHandler(interactorMock)

	weatherData := &domain.CurrentWeather{
		Sunset:   123,
		Humidity: 42,
	}
	forecastData := &domain.Forecast{
		Pressure: 45,
	}

	interactorMock.On("GetWeather", params.City, params.Country, params.ForecastDay).Return(weatherData, forecastData, nil)
	r := handler.GetWeather(params)
	resp, ok := r.(*weather.GetWeatherOK)
	assert.True(t, ok)
	expectedResp := toWeatherResponse(weatherData)
	expectedResp.Forecast = toForecastResponse(forecastData)
	assert.Equal(t, expectedResp, resp.Payload)
}

func TestGetWeather_HappyPathWithCurrWeather(t *testing.T) {
	params := weather.GetWeatherParams{
		City:        "city",
		Country:     "country",
		ForecastDay: (*int64)(nil),
	}
	interactorMock := &mocks.WeatherInteractor{}
	handler := NewWeatherHandler(interactorMock)

	weatherData := &domain.CurrentWeather{
		Sunset:   123,
		Humidity: 42,
	}

	interactorMock.On("GetWeather", params.City, params.Country, params.ForecastDay).Return(weatherData, nil, nil)
	r := handler.GetWeather(params)
	resp, ok := r.(*weather.GetWeatherOK)
	assert.True(t, ok)
	assert.Equal(t, toWeatherResponse(weatherData), resp.Payload)
}

func TestToResponseWeather_ReturnsResponse(t *testing.T) {
	domainWeather := &domain.CurrentWeather{
		Cloudiness: 	"value",
		Humidity:       56,
		Pressure:       11,
		RequestedTime:  time.Now(),
		Sunrise:        2,
		Sunset:         2,
		Temperature:    15,
		Lat:            15,
		Lon:            56,
		City: "city",
		Country: "country",
	}
	expected := &models.WeatherWithForecast{
		Cloudiness:domainWeather.Cloudiness,
		Forecast: (*models.Forecast)(nil),
		GeoCoordinates:fmt.Sprintf("[%v, %v]", domainWeather.Lat, domainWeather.Lon),
		Humidity:formatHumidity(domainWeather.Humidity),
		Pressure:formatPressure(domainWeather.Pressure),
		RequestedTime: domainWeather.RequestedTime.Format("2006-02-01 15:04:05"),
		Sunrise:formatSunriseOrSunset(domainWeather.Sunrise),
		Sunset:formatSunriseOrSunset(domainWeather.Sunset),
		Temperature:formatTemp(domainWeather.Temperature),
		Wind:formatWind(domainWeather.WindSpeed, domainWeather.WindDegree),
		LocationName:fmt.Sprintf("%s, %s", domainWeather.City, domainWeather.Country),
	}
	respWeather := toWeatherResponse(domainWeather)
	assert.Equal(t, expected, respWeather)
}

func TestToForecastResponse_ReturnsResponse(t *testing.T) {
	domainFC := &domain.Forecast{
		DateTime:    45,
		Humidity:    45,
		Pressure:    4,
		Sunrise:     46455,
		Sunset:      53566,
		Temperature: 23,
		WindSpeed:   1,
		WindDegree:  1,
	}
	expected := &models.Forecast{
		Date:time.Unix(int64(domainFC.DateTime), 0).Format("2006-02-01 15:04:05"),
		Humidity:formatHumidity(domainFC.Humidity),
		Pressure:formatPressure(domainFC.Pressure),
		Sunrise:formatSunriseOrSunset(domainFC.Sunrise),
		Sunset:formatSunriseOrSunset(domainFC.Sunset),
		Temperature:formatTemp(domainFC.Temperature),
		Wind:formatWind(domainFC.WindSpeed, domainFC.WindDegree),
	}
	respFC := toForecastResponse(domainFC)
	assert.Equal(t, expected, respFC)
}
