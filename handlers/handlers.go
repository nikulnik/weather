package handlers

import (
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

func NewWeatherHandler(weatherIntercator interactors.WeatherInteractor) WeatherHandler {
	return &weatherHandler{weatherInteractor: weatherIntercator}
}

type weatherHandler struct {
	weatherInteractor interactors.WeatherInteractor
}

func (h *weatherHandler) GetWeather(params weather.GetWeatherParams) middleware.Responder {
	weatherData, err := h.weatherInteractor.GetWeather(params.City, params.Country, params.ForecastDay)
	if err != nil {
		return weather.NewGetWeatherDefault(500).WithPayload(&models.Error{Error: err.Error()})
	}

	return weather.NewGetWeatherOK().WithPayload(weatherData)
}