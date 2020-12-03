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

	interactorMock.On("GetCurrentWeather", params.City, params.Country, params.ForecastDay).Return(nil, nil, expectedError)
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
		Sunset: time.Now(),
		Humidity: domain.Humidity{
			Value: "1",
			Unit:  "%",
		},
	}
	forecastData := &domain.Forecast{
		Pressure: 45,
	}

	interactorMock.On("GetCurrentWeather", params.City, params.Country, params.ForecastDay).Return(weatherData, forecastData, nil)
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
		Sunset: time.Now(),
		Humidity: domain.Humidity{
			Value: "1",
			Unit:  "%",
		},
	}

	interactorMock.On("GetCurrentWeather", params.City, params.Country, params.ForecastDay).Return(weatherData, nil, nil)
	r := handler.GetWeather(params)
	resp, ok := r.(*weather.GetWeatherOK)
	assert.True(t, ok)
	assert.Equal(t, toWeatherResponse(weatherData), resp.Payload)
}

func TestToResponseWeather_ReturnsResponse(t *testing.T) {
	domainWeather := &domain.CurrentWeather{
		Country:     "country",
		City:        "city",
		Cloudiness:  "broken clouds",
		Humidity:    domain.Humidity{Value: "12", Unit: "%"},
		Pressure:    domain.Pressure{Value: "31", Unit: "hPa"},
		Sunrise:     time.Now(),
		Sunset:      time.Now(),
		Temperature: domain.Temperature{Value: "44", Unit: "celsius"},
		Wind:        domain.Wind{Speed: "", Unit: "", Name: "light breeze", DirectionName: "south"},
		Lat:         "55",
		Lon:         "55",
	}

	resp := toWeatherResponse(domainWeather)

	assert.Equal(t, fmt.Sprintf("%s, %s", domainWeather.City, domainWeather.Country), resp.LocationName)
	assert.Equal(t, formatTemp(domainWeather.Temperature.Value), resp.Temperature)
	assert.Equal(t, fmt.Sprintf("%s, %v m/s, %s", domainWeather.Wind.Name, domainWeather.Wind.Speed, domainWeather.Wind.DirectionName), resp.Wind)
	assert.Equal(t, domainWeather.Cloudiness, resp.Cloudiness)
	assert.Equal(t, fmt.Sprintf("%s %s", domainWeather.Pressure.Value, domainWeather.Pressure.Unit), resp.Pressure)
	assert.Equal(t, fmt.Sprintf(`%s %s`, domainWeather.Humidity.Value, domainWeather.Humidity.Unit), resp.Humidity)
	assert.Equal(t, domainWeather.Sunrise.Format("15:04"), resp.Sunrise)
	assert.Equal(t, domainWeather.Sunset.Format("15:04"), resp.Sunset)
	assert.Equal(t, fmt.Sprintf("[%s, %s]", domainWeather.Lat, domainWeather.Lon), resp.GeoCoordinates)
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
		Date:        time.Unix(int64(domainFC.DateTime), 0).Format("2006-02-01 15:04:05"),
		Humidity:    formatHumidity(domainFC.Humidity),
		Pressure:    formatPressure(domainFC.Pressure),
		Sunrise:     formatSunriseOrSunset(domainFC.Sunrise),
		Sunset:      formatSunriseOrSunset(domainFC.Sunset),
		Temperature: formatTemp(domainFC.Temperature),
		Wind:        formatWind(domainFC.WindSpeed, domainFC.WindDegree),
	}
	respFC := toForecastResponse(domainFC)
	assert.Equal(t, expected, respFC)
}

func TestGetWindDirectionByDegree(t *testing.T) {
	var tests = []struct {
		deg       float64
		direction string
	}{
		{5, "north"},
		{355, "north"},
		{22, "north-northeast"},
		{44, "northeast"},
		{66, "east-northeast"},
		{133, "southeast"},
		{177, "south"},
		{200, "south-southwest"},
		{222, "southwest"},
		{244, "west-southwest"},
		{290, "west-northwest"},
		{309, "northwest"},
		{333, "north-northwest"},
	}

	for _, tt := range tests {
		t.Run(tt.direction, func(t *testing.T) {
			assert.Equal(t, tt.direction, getWindDirectionByDegree(tt.deg))
		})
	}
}

func TestGetWindTypeBySpeed(t *testing.T) {
	var tests = []struct {
		speed    float64
		windType string
	}{
		{0.3, "Calm"},
		{1.2, "Light air"},
		{2.2, "Light breeze"},
		{4, "Gentle breeze"},
		{6, "Moderate breeze"},
		{9, "Fresh breeze"},
		{12, "Strong breeze"},
		{14, "Moderate gale"},
		{17, "Fresh gale"},
		{22, "Strong gale"},
		{25, "Whole gale"},
		{29, "Storm"},
		{50, "Hurricane"},
	}

	for _, tt := range tests {
		t.Run(tt.windType, func(t *testing.T) {
			assert.Equal(t, tt.windType, getWindTypeBySpeed(tt.speed))
		})
	}
}
