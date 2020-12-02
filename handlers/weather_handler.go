package handlers

import (
	"github.com/nikulnik/weather/domain"
	"github.com/nikulnik/weather/interactors"
	"github.com/nikulnik/weather/models"
	"github.com/nikulnik/weather/restapi/operations/weather"

	"github.com/go-openapi/runtime/middleware"
)

var (
	internalErr = `internal error`
)

type WeatherHandler interface {
	GetWeather(params weather.GetWeatherParams) middleware.Responder
}

func NewWeatherHandler(weatherInteractor interactors.WeatherInteractor) WeatherHandler {
	return &weatherHandler{weatherInteractor: weatherInteractor}
}

type weatherHandler struct {
	weatherInteractor interactors.WeatherInteractor
}

func (h *weatherHandler) GetWeather(params weather.GetWeatherParams) middleware.Responder {
	weatherData, err := h.weatherInteractor.GetWeather(params.City, params.Country, params.ForecastDay)
	if err != nil {
		return weather.NewGetWeatherDefault(500).WithPayload(&models.Error{Error: err.Error()})
	}
	return weather.NewGetWeatherOK().WithPayload(toResponseWeather(weatherData))
}

func toResponseWeather(weatherDomain *domain.WeatherWithForecast) *models.WeatherWithForecast {
	resp := &models.WeatherWithForecast{
		Cloudiness:     weatherDomain.Cloudiness,
		GeoCoordinates: weatherDomain.GeoCoordinates,
		Humidity:       weatherDomain.Humidity,
		LocationName:   weatherDomain.LocationName,
		Pressure:       weatherDomain.Pressure,
		RequestedTime:  weatherDomain.RequestedTime,
		Sunrise:        weatherDomain.Sunrise,
		Sunset:         weatherDomain.Sunset,
		Temperature:    weatherDomain.Temperature,
		Wind:           weatherDomain.Wind,
	}
	if weatherDomain.Forecast != nil {
		var mf = models.Forecast(*weatherDomain.Forecast)
		resp.Forecast = &mf
	}
	return resp
}
