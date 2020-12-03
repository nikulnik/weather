package interactors

import (
	"github.com/nikulnik/weather/cache"
	"github.com/nikulnik/weather/domain"
	"github.com/nikulnik/weather/rest"
)

type WeatherInteractor interface {
	GetCurrentWeather(city, country string, forecastDay *int64) (*domain.CurrentWeather, *domain.Forecast, error)
}

type weatherInteractor struct {
	cache                cache.Cache
	openWeatherMapClient rest.OpenWeatherMapClient
}

func NewWeatherInteractor(client rest.OpenWeatherMapClient, cache cache.Cache) WeatherInteractor {
	return &weatherInteractor{
		cache:                cache,
		openWeatherMapClient: client,
	}
}

func (wi *weatherInteractor) GetCurrentWeather(city, country string, forecastDay *int64) (*domain.CurrentWeather, *domain.Forecast, error) {
	var weather *domain.CurrentWeather
	var err error

	// Attempt to get weather with forecast from cache
	weather = wi.cache.GetCurrentWeather(city, country)
	if weather == nil {
		// Request weather
		weather, err = wi.openWeatherMapClient.GetCurrentWeather(city, country)
		if err != nil {
			return nil, nil, err
		}

		// Store weather object in the cache
		wi.cache.SetWeather(city, country, weather)
	}

	// If day for forecast is provided
	if forecastDay != nil {
		// Get forecast from the cache
		forecast := wi.cache.GetForecast(weather.Lat, weather.Lon)
		if forecast != nil {
			return weather, forecast, nil
		}

		// Request forecast
		forecast, err := wi.openWeatherMapClient.GetForecast(weather.Lat, weather.Lon, *forecastDay)
		if err != nil {
			return nil, nil, err
		}

		// Put the forecast to the cache
		wi.cache.SetForecast(weather.Lat, weather.Lon, forecast)
		return weather, forecast, nil
	}

	return weather, nil, nil
}
