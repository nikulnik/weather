package interactors

import (
	"github.com/nikulnik/weather/domain"
	"github.com/nikulnik/weather/rest"
)

type WeatherInteractor interface {
	GetWeather(city, country string) (*domain.OpenWeatherResp, error)
}

type weatherInteractor struct {
	cache                string
	openWeatherMapClient rest.OpenWeatherMapClient
}

func NewWeatherInteractor(client rest.OpenWeatherMapClient) WeatherInteractor {
	return &weatherInteractor{
		cache:                "",
		openWeatherMapClient: client,
	}
}

func (wi *weatherInteractor) GetWeather(city, country string) (*domain.OpenWeatherResp, error) {
	return wi.openWeatherMapClient.GetSomething(city, country)
}
