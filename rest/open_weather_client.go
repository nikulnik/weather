package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nikulnik/weather/domain"
)

const openWeatherMapURLFmt = "http://api.openweathermap.org/data/2.5/weather?q=%s,%s&appid=%s"

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

func (c *OpenWeatherMapClient) GetSomething(city, countryCode string) (*domain.OpenWeatherResp, error) {
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
