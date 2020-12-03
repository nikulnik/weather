package rest

import (
	"encoding/json"
	"fmt"
	"github.com/nikulnik/weather/domain"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	openWeatherCurrentWeatherURLFmt = "http://api.openweathermap.org/data/2.5/weather?q=%s,%s&units=metric&mode=xml&appid=%s"
	openWeatherForecastURLFmt       = "http://api.openweathermap.org/data/2.5/onecall?lat=%v&lon=%v&exclude=current,minutely,hourly,alerts&units=metric&appid=%s"

	openWeatherErrRespFmt = "openweathermap API responded with error: %s"
	)

type OpenWeatherMapClient interface {
	GetWeather(city, countryCode string) (*domain.CurrentWeather, error)
	GetForecast(lat, lon float64, day int64) (*domain.Forecast, error)
}

type openWeatherMapClient struct {
	ApiKey     string
	httpClient http.Client
}

func NewOpenWeatherMapClient(apiKey string) OpenWeatherMapClient {
	client := &openWeatherMapClient{
		ApiKey:     apiKey,
		httpClient: http.Client{},
	}
	return client
}

func (c *openWeatherMapClient) GetWeather(city, countryCode string) (*domain.CurrentWeather, error) {
	resp, err := http.Get(
		fmt.Sprintf(openWeatherCurrentWeatherURLFmt, city, countryCode, c.ApiKey),
	)
	if err != nil {
		return nil, err
	}
	data := &OpenWeatherResp{}
	if resp.StatusCode == 200 {
		decoder := json.NewDecoder(ioutil.NopCloser(resp.Body))
		decoder.DisallowUnknownFields()
		err = decoder.Decode(data)
		if err != nil {
			return nil, err
		}
		return toWeatherDomain(data), nil
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

func (c *openWeatherMapClient) GetForecast(lat, lon float64, day int64) (*domain.Forecast, error) {
	resp, err := http.Get(
		fmt.Sprintf(openWeatherForecastURLFmt, lat, lon, c.ApiKey),
	)
	if err != nil {
		return nil, err
	}

	data := &ForecastResp{}
	if resp.StatusCode == 200 {
		decoder := json.NewDecoder(ioutil.NopCloser(resp.Body))
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

func toWeatherDomain(weather *OpenWeatherResp) *domain.CurrentWeather {
	var weatherModel = &domain.CurrentWeather{
		Temperature:   weather.Main.Temp,
		WindSpeed:     weather.Wind.Speed,
		WindDegree:    weather.Wind.Deg,
		Cloudiness:    weather.GetCloudsDescription(),
		Pressure:      weather.Main.Pressure,
		Humidity:      weather.Main.Humidity,
		Sunrise:       weather.Sys.Sunrise,
		Sunset:        weather.Sys.Sunset,
		RequestedTime: time.Now(),
		Lat:           weather.Coord.Lat,
		Lon:           weather.Coord.Lon,
		City:          weather.Name,
		Country:       weather.Sys.Country,
	}
	return weatherModel
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

func formatTemp(temp float64) string {
	return fmt.Sprintf("%v °C", temp)
}

func formatWind(speed, degree float64) string {
	return fmt.Sprintf("%s, %v m/s, %s", getWindTypeBySpeed(speed), speed, getWindDirectionByDegree(degree))
}

func formatPressure(pressure int) string {
	return fmt.Sprintf("%d hpa", pressure)
}
func formatHumidity(humidity int) string {
	return fmt.Sprintf(`%d %%`, humidity)
}

func formatSunriseOrSunset(v int) string {
	return time.Unix(int64(v), 0).Format("15:04")
}
func getWindDirectionByDegree(degree float64) string {
	if degree < 11.25 {
		return "north"
	}
	if degree < 33.75 {
		return "north-northeast"
	}
	if degree < 56.25 {
		return "northeast"
	}
	if degree < 78.75 {
		return "east-northeast"
	}
	if degree < 101.25 {
		return "east"
	}
	if degree < 123.75 {
		return "east-southeast"
	}
	if degree < 146.25 {
		return "southeast"
	}
	if degree < 168.75 {
		return "south-southeast"
	}
	if degree < 191.25 {
		return "south"
	}
	if degree < 213.75 {
		return "south-southwest"
	}
	if degree < 236.25 {
		return "southwest"
	}
	if degree < 258.75 {
		return "west-southwest"
	}
	if degree < 281.25 {
		return "west"
	}
	if degree < 303.75 {
		return "west-northwest"
	}
	if degree < 326.25 {
		return "northwest"
	}
	if degree < 348.25 {
		return "north-northwest"
	}
	return "north"
}

// m/s
func getWindTypeBySpeed(speed float64) string {
	if speed < 0.5 {
		return "Calm"
	}
	if speed < 1.5 {
		return "Light air"
	}
	if speed < 3 {
		return "Light breeze"
	}
	if speed < 5 {
		return "Gentle breeze"
	}
	if speed < 8 {
		return "Moderate breeze"
	}
	if speed < 10.5 {
		return "Fresh breeze"
	}
	if speed < 13.5 {
		return "Strong breeze"
	}
	if speed < 16.5 {
		return "Moderate gale"
	}
	if speed < 20 {
		return "Fresh gale"
	}
	if speed < 23.5 {
		return "Strong gale"
	}
	if speed < 27.5 {
		return "Whole gale"
	}
	if speed < 31.5 {
		return "Storm"
	}
	return "Hurricane"
}
