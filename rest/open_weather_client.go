package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nikulnik/weather/domain"
)

const openWeatherMapURLFmt = "http://api.openweathermap.org/data/2.5/weather?q=%s,%s&units=metric&appid=%s"
const openWeatherForecastURLFmt = "http://api.openweathermap.org/data/2.5/onecall?lat=%v&lon=%v&exclude=current,minutely,hourly,alerts&units=metric&appid=%s"

type OpenWeatherMapClient struct {
	ApiKey     string
	httpClient http.Client
}

func NewOpenWeatherMapClient(apiKey string) OpenWeatherMapClient {
	client := OpenWeatherMapClient{
		ApiKey:     apiKey,
		httpClient: http.Client{},
	}
	return client
}

func (c *OpenWeatherMapClient) GetWeather(city, countryCode string) (*domain.OpenWeatherResp, error) {
	resp, err := http.Get(
		fmt.Sprintf(openWeatherMapURLFmt, city, countryCode, c.ApiKey),
	)
	if err != nil {
		return nil, err
	}
	data := &domain.OpenWeatherResp{}
	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *OpenWeatherMapClient) GetForecast(lat, lon float64) (*domain.Forecast, error) {
	resp, err := http.Get(
		fmt.Sprintf(openWeatherForecastURLFmt, lat, lon, c.ApiKey),
	)
	if err != nil {
		return nil, err
	}
	dataForecast := &domain.Forecast{}
	err = json.NewDecoder(resp.Body).Decode(dataForecast)
	if err != nil {
		return nil, err
	}
	return dataForecast, nil
}
