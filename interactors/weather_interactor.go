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

// GetCurrentWeather returns the current weather. Returns the forecast if forecastDay is not nil
func (wi *weatherInteractor) GetCurrentWeather(city, country string, forecastDay *int64) (*domain.CurrentWeather, *domain.Forecast, error) {
	var weather *domain.CurrentWeather
	var err error

	// Attempt to get weather with forecast from cache
	weatherInt := wi.cache.Get(city + "*" + country)
	if weatherInt == nil {
		// Request weather via openweather API
		weather, err = wi.openWeatherMapClient.GetCurrentWeather(city, country)
		if err != nil {
			return nil, nil, err
		}

		// Store the weather object in the cache
		wi.cache.Set(city+"*"+country, weather)
	}
	if weatherInt != nil {
		weather = weatherInt.(*domain.CurrentWeather)
	}

	if forecastDay != nil {
		// Get forecast from the cache
		forecastInt := wi.cache.Get(weather.Lat + "*" + weather.Lon)
		if forecastInt != nil {
			return weather, forecastInt.(*domain.Forecast), nil
		}

		// Request the forecast
		forecast, err := wi.openWeatherMapClient.GetForecast(weather.Lat, weather.Lon, *forecastDay)
		if err != nil {
			return nil, nil, err
		}

		// Put the forecast to the cache
		wi.cache.Set(weather.Lat+"*"+weather.Lon, forecast)
		return weather, forecast, nil
	}

	return weather, nil, nil
}
