package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const testAPPID = "12345"

var openWeatherValidResponseCurrentWeather = []byte(`{"coord":{"lon":-74.08,"lat":4.61},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03n"}],"base":"stations","main":{"temp":285.15,"feels_like":283.98,"temp_min":285.15,"temp_max":285.15,"pressure":1025,"humidity":93},"visibility":10000,"wind":{"speed":2.1,"deg":50},"clouds":{"all":40},"dt":1606904259,"sys":{"type":1,"id":8582,"country":"CO","sunrise":1606906202,"sunset":1606948935},"timezone":-18000,"id":3688689,"name":"Bogotá","cod":200}`)
var openWeatherValidResponseForecast = []byte(`{"lat":40,"lon":40,"timezone":"Europe/Istanbul","timezone_offset":10800,"daily":[{"dt":1606899600,"sunrise":1606883042,"sunset":1606917322,"temp":{"day":274.38,"min":269.92,"max":282.15,"night":269.92,"eve":277.28,"morn":272.29},"feels_like":{"day":270.18,"night":266.77,"eve":274.56,"morn":268.26},"pressure":1024,"humidity":89,"dew_point":271.67,"wind_speed":3.09,"wind_deg":143,"weather":[{"id":600,"main":"Snow","description":"light snow","icon":"13d"}],"clouds":100,"pop":0.52,"snow":0.82,"uvi":1.27},{"dt":1606986000,"sunrise":1606969500,"sunset":1607003711,"temp":{"day":274.28,"min":268.34,"max":275.16,"night":269.4,"eve":269.29,"morn":268.48},"feels_like":{"day":271.01,"night":265.97,"eve":266.23,"morn":265.23},"pressure":1024,"humidity":87,"dew_point":270.17,"wind_speed":1.67,"wind_deg":159,"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":4,"pop":0,"uvi":2.13},{"dt":1607072400,"sunrise":1607055958,"sunset":1607090101,"temp":{"day":274.17,"min":267.5,"max":275,"night":268.93,"eve":269.25,"morn":267.57},"feels_like":{"day":270.79,"night":265.67,"eve":266.16,"morn":264.07},"pressure":1022,"humidity":84,"dew_point":268,"wind_speed":1.71,"wind_deg":161,"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}],"clouds":86,"pop":0,"uvi":2.03},{"dt":1607158800,"sunrise":1607142414,"sunset":1607176494,"temp":{"day":274.99,"min":266.85,"max":276.18,"night":269.38,"eve":269.56,"morn":267.05},"feels_like":{"day":272.4,"night":266.41,"eve":266.52,"morn":263.93},"pressure":1019,"humidity":77,"dew_point":266.76,"wind_speed":0.52,"wind_deg":146,"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":0,"pop":0,"uvi":1.93},{"dt":1607245200,"sunrise":1607228868,"sunset":1607262889,"temp":{"day":274.91,"min":267.08,"max":275.99,"night":269.18,"eve":269.41,"morn":267.08},"feels_like":{"day":272.09,"night":265.97,"eve":266.42,"morn":263.77},"pressure":1019,"humidity":76,"dew_point":265.75,"wind_speed":0.8,"wind_deg":286,"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":0,"pop":0,"uvi":1.75},{"dt":1607331600,"sunrise":1607315322,"sunset":1607349286,"temp":{"day":273.95,"min":266.79,"max":274.77,"night":269.42,"eve":269.44,"morn":266.79},"feels_like":{"day":270.5,"night":265.78,"eve":265.95,"morn":263.18},"pressure":1022,"humidity":78,"dew_point":263.67,"wind_speed":1.59,"wind_deg":159,"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":0,"pop":0,"uvi":2},{"dt":1607418000,"sunrise":1607401774,"sunset":1607435685,"temp":{"day":274.03,"min":267.03,"max":274.91,"night":269.79,"eve":269.62,"morn":267.03},"feels_like":{"day":270.3,"night":266.19,"eve":266.18,"morn":263.19},"pressure":1021,"humidity":75,"dew_point":262.33,"wind_speed":1.91,"wind_deg":156,"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":0,"pop":0,"uvi":2},{"dt":1607504400,"sunrise":1607488225,"sunset":1607522087,"temp":{"day":273.38,"min":268.54,"max":274.21,"night":271.07,"eve":271.13,"morn":268.54},"feels_like":{"day":269.67,"night":268.02,"eve":268.01,"morn":264.67},"pressure":1020,"humidity":81,"dew_point":263.92,"wind_speed":1.96,"wind_deg":152,"weather":[{"id":600,"main":"Snow","description":"light snow","icon":"13d"}],"clouds":100,"pop":0.32,"snow":1.06,"uvi":2}]}`)

func TestGetWeather_WhenOpenWeatherReturnsError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	var city, country = "city", "country"
	path := fmt.Sprintf(openWeatherCurrentWeatherURLFmt, city, country, testAPPID)
	expectedErr := errors.New("expected err")
	httpmock.RegisterResponder("GET", path, httpmock.NewErrorResponder(expectedErr))

	client := NewOpenWeatherMapClient(testAPPID)
	weather, err := client.GetWeather(city, country)
	assert.Error(t, err)
	assert.Nil(t, weather)
}

func TestGetWeather_WhenOpenWeatherReturnsUnexpectedJSON(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	var city, country = "city", "country"
	path := fmt.Sprintf(openWeatherCurrentWeatherURLFmt, city, country, testAPPID)
	type wrongResp struct {
		Field string
	}

	httpmock.RegisterResponder("GET", path, httpmock.NewJsonResponderOrPanic(200, &wrongResp{"123"}))

	client := NewOpenWeatherMapClient(testAPPID)
	weather, err := client.GetWeather(city, country)
	assert.Error(t, err)
	assert.Nil(t, weather)
}

func TestGetWeather_WhenOpenWeatherReturnsValidData_ReturnsWeather(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	var city, country = "city", "country"
	path := fmt.Sprintf(openWeatherCurrentWeatherURLFmt, city, country, testAPPID)

	httpmock.RegisterResponder("GET", path, httpmock.NewBytesResponder(200, openWeatherValidResponseCurrentWeather))

	client := NewOpenWeatherMapClient(testAPPID)
	weather, err := client.GetWeather(city, country)
	assert.Nil(t, err)
	assert.NotNil(t, weather)

	resp := &OpenWeatherResp{}
	assert.Nil(t, json.Unmarshal(openWeatherValidResponseCurrentWeather, resp))
	assert.Equal(t, toWeatherDomain(resp), weather)
}

func TestGetForecast_WhenOpenWeatherReturnsError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	var lat, lon = 5.5, 5.5
	var day int64 = 5
	path := fmt.Sprintf(openWeatherForecastURLFmt, lat, lon, testAPPID)
	expectedErr := errors.New("expected err")
	httpmock.RegisterResponder("GET", path, httpmock.NewErrorResponder(expectedErr))

	client := NewOpenWeatherMapClient(testAPPID)
	forecast, err := client.GetForecast(lat, lon, day)
	assert.Error(t, err)
	assert.Nil(t, forecast)
}

func TestGetForecast_WhenOpenWeatherReturnsUnexpectedJSON(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	var lat, lon = 5.5, 5.5
	var day int64 = 5
	path := fmt.Sprintf(openWeatherForecastURLFmt, lat, lon, testAPPID)
	type wrongResp struct {
		Field string
	}

	httpmock.RegisterResponder("GET", path, httpmock.NewJsonResponderOrPanic(200, &wrongResp{"123"}))

	client := NewOpenWeatherMapClient(testAPPID)
	forecast, err := client.GetForecast(lat, lon, day)
	assert.Error(t, err)
	assert.Nil(t, forecast)
}

func TestGetForecast_WhenOpenWeatherReturnsValidData_ReturnsForecast(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	var lat, lon = 5.5, 5.5
	var day int64 = 5
	path := fmt.Sprintf(openWeatherForecastURLFmt, lat, lon, testAPPID)

	httpmock.RegisterResponder("GET", path, httpmock.NewBytesResponder(200, openWeatherValidResponseForecast))

	client := NewOpenWeatherMapClient(testAPPID)
	forecast, err := client.GetForecast(lat, lon, day)
	assert.Nil(t, err)
	assert.NotNil(t, forecast)

	resp := &ForecastResp{}
	assert.Nil(t, json.Unmarshal(openWeatherValidResponseForecast, resp))

	expected, err := toForecastDomain(resp, day)
	assert.Nil(t, err)
	assert.Equal(t, expected, forecast)
}

func TestToWeatherDomain_ReturnsDomainRepresentation(t *testing.T) {
	weather := &OpenWeatherResp{
		Coord: Coord{4, 5},
		Weather: []Weather{
			{ID: 1, Main: "clouds", Description: "description"},
		},
		Base: "stations",
		Main: Main{Temp: 285.15,
			FeelsLike: 283.98,
			TempMin:   285.15,
			TempMax:   285.15,
			Pressure:  1025,
			Humidity:  93},
		Visibility: 31,
		Wind:       Wind{},
		Clouds:     Clouds{},
		Dt:         321,
		Sys: Sys{
			Type:    1,
			ID:      8582,
			Country: "CO",
			Sunrise: 1606906202,
			Sunset:  1606948935,
		},
		Timezone: 2,
		ID:       1,
		Name:     "Some City",
		Cod:      200,
	}
	result := toWeatherDomain(weather)

	assert.Equal(t, fmt.Sprintf("%v °C", weather.Main.Temp), result.Temperature)
	assert.Equal(t, weather.GetCloudsDescription(), result.Cloudiness)
	assert.Equal(t, fmt.Sprintf("%d hpa", weather.Main.Pressure), result.Pressure)
	assert.Equal(t, time.Unix(int64(weather.Sys.Sunrise), 0).Format("15:04"), result.Sunrise)
	assert.Equal(t, time.Unix(int64(weather.Sys.Sunset), 0).Format("15:04"), result.Sunset)
	assert.Equal(t, time.Now().Format("2006-02-01 15:04:05"), result.RequestedTime)
	assert.Equal(t, weather.Coord.Lat, result.Lat)
	assert.Equal(t, weather.Coord.Lon, result.Lon)
}
