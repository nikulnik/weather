package cache

import (
	"testing"
	"time"

	"github.com/nikulnik/weather/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewCache_ReturnsCache(t *testing.T) {
	cache := NewCache(time.Second)
	assert.NotNil(t, cache)
}

func TestSetAndGetWeather_GetsCorrectWeather(t *testing.T) {
	var city, country = "city", "country"
	cache := NewCache(time.Minute)
	value := &domain.WeatherWithForecast{Sunset: "6:45"}
	cache.SetWeather(city, country, value)
	assert.Equal(t, value, cache.GetWeather(city, country))
}

func TestSetAndGetForecast_ReturnsCorrectForecast(t *testing.T) {
	var lat, lon float64 = 1, 1
	cache := NewCache(time.Minute)
	value := &domain.Forecast{Sunset: "6:45"}
	cache.SetForecast(lat, lon, value)
	assert.Equal(t, value, cache.GetForecast(lat, lon))
}

func TestGetForecast_ReturnsNilAfterTimeout(t *testing.T) {
	var lat, lon float64 = 1, 1
	cache := NewCache(time.Millisecond * 3)
	value := &domain.Forecast{Sunset: "6:45"}
	cache.SetForecast(lat, lon, value)
	time.Sleep(time.Millisecond * 4)
	assert.Nil(t, cache.GetForecast(lat, lon))
}

func TestGetWeather_ReturnsNilAfterTimeout(t *testing.T) {
	var city, country = "city", "country"
	cache := NewCache(time.Millisecond * 3)
	value := &domain.WeatherWithForecast{Sunset: "6:45"}
	cache.SetWeather(city, country, value)
	time.Sleep(time.Millisecond * 4)
	assert.Nil(t, cache.GetWeather(city, country))
}

func TestCreateForecastKey_ReturnsCorrectKey(t *testing.T) {
	var lat, lon = 1.123, 1.41424
	cache := cache{}
	assert.Equal(t, "1.123*1.41424", cache.createForecastKey(lat, lon))
}
