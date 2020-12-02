package cache

import (
	"fmt"
	"github.com/nikulnik/weather/domain"
	"sync"
	"time"
)

type Cache interface {
	SetWeather(city, country string, value *domain.WeatherWithForecast)
	SetForecast(lat, lon float64, value *domain.Forecast)

	GetWeather(city, country string) *domain.WeatherWithForecast
	GetForecast(lat, lon float64) *domain.Forecast
}

type cache struct {
	dataWeather  map[string]*domain.WeatherWithForecast
	dataForecast map[string]*domain.Forecast
	mux          *sync.Mutex
	ttl          time.Duration
}

func NewCache(ttl time.Duration) Cache {
	return &cache{
		dataWeather:  make(map[string]*domain.WeatherWithForecast),
		dataForecast: make(map[string]*domain.Forecast),
		mux:          &sync.Mutex{},
		ttl:          ttl,
	}
}

func (c cache) SetWeather(city, country string, value *domain.WeatherWithForecast) {
	key := city + "*" + country
	c.mux.Lock()
	c.dataWeather[key] = value
	c.mux.Unlock()
	time.AfterFunc(c.ttl, func() {
		c.mux.Lock()
		delete(c.dataWeather, key)
		c.mux.Unlock()
	})
}

func (c cache) SetForecast(lat, lon float64, value *domain.Forecast) {
	key := c.createForecastKey(lat, lon)
	c.mux.Lock()
	c.dataForecast[key] = value
	c.mux.Unlock()
	time.AfterFunc(c.ttl, func() {
		c.mux.Lock()
		delete(c.dataForecast, key)
		c.mux.Unlock()
	})
}

func (c cache) GetWeather(city, country string) *domain.WeatherWithForecast {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.dataWeather[city+"*"+country]
}

func (c cache) GetForecast(lat, lon float64) *domain.Forecast {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.dataForecast[c.createForecastKey(lat, lon)]
}

func (c cache) createForecastKey(lat, lon float64) string {
	return fmt.Sprintf("%v*%v", lat, lon)
}
