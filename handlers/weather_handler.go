package handlers

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/nikulnik/weather/domain"
	"github.com/nikulnik/weather/interactors"
	"github.com/nikulnik/weather/models"
	"github.com/nikulnik/weather/restapi/operations/weather"
	"time"
)

var (
	internalErr = `internal error`
)

type WeatherHandler interface {
	GetWeather(params weather.GetWeatherParams) middleware.Responder
}

func NewWeatherHandler(weatherInteractor interactors.WeatherInteractor) WeatherHandler {
	return &weatherHandler{weatherInteractor: weatherInteractor}
}

type weatherHandler struct {
	weatherInteractor interactors.WeatherInteractor
}

func (h *weatherHandler) GetWeather(params weather.GetWeatherParams) middleware.Responder {
	weatherData, forecast, err := h.weatherInteractor.GetCurrentWeather(params.City, params.Country, params.ForecastDay)
	if err != nil {
		return weather.NewGetWeatherDefault(500).WithPayload(&models.Error{Error: err.Error()})
	}
	response := toWeatherResponse(weatherData)
	if forecast != nil {
		response.Forecast = toForecastResponse(forecast)
	}
	return weather.NewGetWeatherOK().WithPayload(response)
}

func toWeatherResponse(domainWeather *domain.CurrentWeather) *models.WeatherWithForecast {
	var weatherModel = &models.WeatherWithForecast{
		LocationName:   fmt.Sprintf("%s, %s", domainWeather.City, domainWeather.Country),
		Temperature:    formatTemp(domainWeather.Temperature.Value),
		Wind:           fmt.Sprintf("%s, %v m/s, %s", domainWeather.Wind.Name, domainWeather.Wind.Speed, domainWeather.Wind.DirectionName),
		Cloudiness:     domainWeather.Cloudiness,
		Pressure:       fmt.Sprintf("%s %s", domainWeather.Pressure.Value, domainWeather.Pressure.Unit),
		Humidity:       fmt.Sprintf(`%s %s`, domainWeather.Humidity.Value, domainWeather.Humidity.Unit),
		Sunrise:        domainWeather.Sunrise.Format("15:04"),
		Sunset:         domainWeather.Sunset.Format("15:04"),
		GeoCoordinates: fmt.Sprintf("[%s, %s]", domainWeather.Lat, domainWeather.Lon),
		RequestedTime:  time.Now().Format("2006-02-01 15:04:05"),
	}
	return weatherModel
}

func toForecastResponse(domainFC *domain.Forecast) *models.Forecast {
	model := &models.Forecast{
		Temperature: formatTemp(domainFC.Temperature),
		Wind:        formatWind(domainFC.WindSpeed, domainFC.WindDegree),
		Pressure:    formatPressure(domainFC.Pressure),
		Humidity:    formatHumidity(domainFC.Humidity),
		Sunrise:     formatSunriseOrSunset(domainFC.Sunrise),
		Sunset:      formatSunriseOrSunset(domainFC.Sunset),
		Date:        time.Unix(int64(domainFC.DateTime), 0).Format("2006-02-01 15:04:05"),
	}
	return model
}

func formatTemp(temp interface{}) string {
	return fmt.Sprintf("%v °C", temp)
}

func formatWind(speed, degree float64) string {
	return fmt.Sprintf("%s, %v m/s, %s", getWindTypeBySpeed(speed), speed, getWindDirectionByDegree(degree))
}

func formatPressure(pressure int) string {
	return fmt.Sprintf("%d hPa", pressure)
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
