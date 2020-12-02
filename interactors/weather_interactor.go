package interactors

import (
	"fmt"
	"github.com/nikulnik/weather/domain"
	"github.com/nikulnik/weather/models"
	"github.com/nikulnik/weather/rest"
	"time"
)

type WeatherInteractor interface {
	GetWeather(city, country string, forecastDay *int64) (*models.WeatherWithForecast, error)
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

func (wi *weatherInteractor) GetWeather(city, country string, forecastDay *int64) (*models.WeatherWithForecast, error) {
	weather, err := wi.openWeatherMapClient.GetWeather(city, country)
	if err != nil {
		return nil, err
	}
	if forecastDay != nil {
		forecast, err := wi.openWeatherMapClient.GetForecast(weather.Coord.Lat, weather.Coord.Lon)
		if err != nil {
			return nil, err
		}
		return toWeatherWithForecast(weather, forecast, *forecastDay)
	}
	return toWeatherWithForecast(weather, nil, 0)
}

func toWeatherWithForecast(weather *domain.OpenWeatherResp, forecast *domain.Forecast, forecastDay int64) (*models.WeatherWithForecast, error) {
	sunriseTime := time.Unix(int64(weather.Sys.Sunrise), 0)
	sunsetTime := time.Unix(int64(weather.Sys.Sunset), 0)
	var weatherWithForecast = &models.WeatherWithForecast{
		LocationName:   fmt.Sprintf("%s, %s", weather.Name, weather.Sys.Country),
		Temperature:    fmt.Sprintf("%v °C", weather.Main.Temp),
		Wind:           fmt.Sprintf("%s, %v m/s, %s", getWindTypeBySpeed(weather.Wind.Speed), weather.Wind.Speed, getWindDirectionByDegree(weather.Wind.Deg)),
		Cloudiness:     weather.GetCloudsDescription(),
		Pressure:       fmt.Sprintf("%d hpa", weather.Main.Pressure),
		Humidity:       fmt.Sprintf(`%d %%`, weather.Main.Humidity),
		Sunrise:        sunriseTime.Format("15:04"),
		Sunset:         sunsetTime.Format("15:04"),
		GeoCoordinates: fmt.Sprintf("[%v, %v]", weather.Coord.Lat, weather.Coord.Lon),
		RequestedTime:  time.Now().Format("2006-02-01 15:04:05"),
	}
	if forecast != nil {
		forecastForDay, err := ToForecastResponse(forecast, forecastDay)
		if err != nil {
			return nil, err
		}
		weatherWithForecast.Forecast = &forecastForDay
	}
	return weatherWithForecast, nil
}

func ToForecastResponse(forecast *domain.Forecast, day int64) (models.Forecast, error) {
	if len(forecast.Daily)-1 < int(day) {
		return models.Forecast{}, fmt.Errorf("cannot get forecast for day %d", day)
	}

	dayFC := forecast.Daily[day]
	resp := models.Forecast{
		Temperature: fmt.Sprintf("%v °C", dayFC.Temp.Day),
		Wind:        fmt.Sprintf("%s, %v m/s, %s", getWindTypeBySpeed(dayFC.WindSpeed), dayFC.WindSpeed, getWindDirectionByDegree(dayFC.WindDeg)),
		Pressure:    fmt.Sprintf("%d hpa", dayFC.Pressure),
		Humidity:    fmt.Sprintf("%d %%", dayFC.Humidity),
		Sunrise:     time.Unix(int64(dayFC.Sunrise), 0).Format("15:04"),
		Sunset:      time.Unix(int64(dayFC.Sunset), 0).Format("15:04"),
		Date:        time.Unix(int64(dayFC.Sunrise), 0).Format("2006-02-01 15:04:05"),
	}
	return resp, nil
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
