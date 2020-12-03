package rest

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/nikulnik/weather/domain"
	"net/http"
	"time"
)

const (
	openWeatherCurrentWeatherURLFmt = "http://api.openweathermap.org/data/2.5/weather?q=%s,%s&units=metric&mode=xml&appid=%s"
	openWeatherForecastURLFmt       = "http://api.openweathermap.org/data/2.5/onecall?lat=%s&lon=%s&exclude=current,minutely,hourly,alerts&units=metric&appid=%s"

	openWeatherErrRespFmt = "openweathermap API responded with error: %s"
)

// OpenWeatherMapClient - Openweathermap http client
type OpenWeatherMapClient interface {
	// GetCurrentWeather gets current weather
	GetCurrentWeather(city, countryCode string) (*domain.CurrentWeather, error)
	// GetForecast gets the forecast by latitude and longitude for the given day. 0 - today
	GetForecast(lat, lon string, day int64) (*domain.Forecast, error)
}

type openWeatherMapClient struct {
	ApiKey     string
	httpClient http.Client
}

// NewOpenWeatherMapClient creates a new client with the given API key
func NewOpenWeatherMapClient(apiKey string) OpenWeatherMapClient {
	client := &openWeatherMapClient{
		ApiKey:     apiKey,
		httpClient: http.Client{},
	}
	return client
}

// GetCurrentWeather gets current weather
func (c *openWeatherMapClient) GetCurrentWeather(city, countryCode string) (*domain.CurrentWeather, error) {
	resp, err := http.Get(
		fmt.Sprintf(openWeatherCurrentWeatherURLFmt, city, countryCode, c.ApiKey),
	)
	if err != nil {
		return nil, err
	}
	data := &OpenWeatherRespXML{}
	if resp.StatusCode == 200 {
		decoder := xml.NewDecoder(resp.Body)

		err = decoder.Decode(data)
		if err != nil {
			return nil, err
		}
		return toWeatherDomain(data)
	}

	errResp := &OpenWeatherErrorResp{}
	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(errResp)
	if err != nil {
		return nil, err
	}

	return nil, fmt.Errorf(openWeatherErrRespFmt, errResp.Message)
}

// GetForecast gets the forecast by latitude and longitude for the given day. 0 - today
func (c *openWeatherMapClient) GetForecast(lat, lon string, day int64) (*domain.Forecast, error) {
	resp, err := http.Get(
		fmt.Sprintf(openWeatherForecastURLFmt, lat, lon, c.ApiKey),
	)
	if err != nil {
		return nil, err
	}

	data := &ForecastResp{}
	if resp.StatusCode == 200 {
		decoder := json.NewDecoder(resp.Body)
		decoder.DisallowUnknownFields()
		err = decoder.Decode(data)
		if err != nil {
			return nil, err
		}
		return toForecastDomain(data, day)
	}

	errResp := &OpenWeatherErrorResp{}
	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(errResp)
	if err != nil {
		return nil, err
	}

	return nil, fmt.Errorf(openWeatherErrRespFmt, errResp.Message)
}

func toWeatherDomain(weather *OpenWeatherRespXML) (*domain.CurrentWeather, error) {
	sunRiseTime, err := time.Parse("2006-01-02T15:04:05", weather.City.Sun.Rise)
	if err != nil {
		return nil, fmt.Errorf("failed to parse sunrise time: %s", err)
	}
	sunSetTime, err := time.Parse("2006-01-02T15:04:05", weather.City.Sun.Set)
	if err != nil {
		return nil, fmt.Errorf("failed to parse sunset time: %s", err)
	}
	var weatherModel = &domain.CurrentWeather{
		Temperature: domain.Temperature{
			Value: weather.Temperature.Value,
			Unit:  weather.Temperature.Unit,
		},
		Cloudiness: weather.Clouds.Name,
		Pressure: domain.Pressure{
			Value: weather.Pressure.Value,
			Unit:  weather.Pressure.Unit,
		},
		Humidity: domain.Humidity{
			Value: weather.Humidity.Value,
			Unit:  weather.Humidity.Unit,
		},
		Wind: domain.Wind{
			Speed:         weather.Wind.Speed.Value,
			Unit:          weather.Wind.Speed.Unit,
			Name:          weather.Wind.Speed.Name,
			DirectionName: weather.Wind.Direction.Name,
		},
		Sunrise: sunRiseTime,
		Sunset:  sunSetTime,
		Lat:     weather.City.Coord.Lat,
		Lon:     weather.City.Coord.Lon,
		City:    weather.City.Name,
		Country: weather.City.Country,
	}
	return weatherModel, nil
}

func toForecastDomain(forecast *ForecastResp, day int64) (*domain.Forecast, error) {
	if len(forecast.Daily)-1 < int(day) {
		return nil, fmt.Errorf("cannot get forecast for day %d", day)
	}

	dayFC := forecast.Daily[day]
	resp := &domain.Forecast{
		Temperature: dayFC.Temp.Day,
		WindDegree:  dayFC.WindDeg,
		WindSpeed:   dayFC.WindSpeed,
		Pressure:    dayFC.Pressure,
		Humidity:    dayFC.Humidity,
		Sunrise:     dayFC.Sunrise,
		Sunset:      dayFC.Sunset,
		DateTime:    dayFC.Dt,
	}
	return resp, nil
}
