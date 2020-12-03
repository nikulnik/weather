package rest

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/nikulnik/weather/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const testAPPID = "12345"

var openWeatherValidResponseCurrentWeather = []byte(`<current>
<city id="3688689" name="Bogotá">
<coord lon="-74.08" lat="4.61"/>
<country>CO</country>
<timezone>-18000</timezone>
<sun rise="2020-12-03T10:50:29" set="2020-12-03T22:42:36"/>
</city>
<temperature value="16" min="16" max="16" unit="celsius"/>
<feels_like value="14.14" unit="celsius"/>
<humidity value="72" unit="%"/>
<pressure value="1028" unit="hPa"/>
<wind>
<speed value="3.1" unit="m/s" name="Light breeze"/>
<gusts/>
<direction value="40" code="NE" name="NorthEast"/>
</wind>
<clouds value="75" name="broken clouds"/>
<visibility value="10000"/>
<precipitation mode="no"/>
<weather number="803" value="broken clouds" icon="04d"/>
<lastupdate value="2020-12-03T15:24:33"/>
</current>`)

var openWeatherValidResponseForecast = []byte(`{"lat":40,"lon":40,"timezone":"Europe/Istanbul","timezone_offset":10800,"daily":[{"dt":1606899600,"sunrise":1606883042,"sunset":1606917322,"temp":{"day":274.38,"min":269.92,"max":282.15,"night":269.92,"eve":277.28,"morn":272.29},"feels_like":{"day":270.18,"night":266.77,"eve":274.56,"morn":268.26},"pressure":1024,"humidity":89,"dew_point":271.67,"wind_speed":3.09,"wind_deg":143,"weather":[{"id":600,"main":"Snow","description":"light snow","icon":"13d"}],"clouds":100,"pop":0.52,"snow":0.82,"uvi":1.27},{"dt":1606986000,"sunrise":1606969500,"sunset":1607003711,"temp":{"day":274.28,"min":268.34,"max":275.16,"night":269.4,"eve":269.29,"morn":268.48},"feels_like":{"day":271.01,"night":265.97,"eve":266.23,"morn":265.23},"pressure":1024,"humidity":87,"dew_point":270.17,"wind_speed":1.67,"wind_deg":159,"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":4,"pop":0,"uvi":2.13},{"dt":1607072400,"sunrise":1607055958,"sunset":1607090101,"temp":{"day":274.17,"min":267.5,"max":275,"night":268.93,"eve":269.25,"morn":267.57},"feels_like":{"day":270.79,"night":265.67,"eve":266.16,"morn":264.07},"pressure":1022,"humidity":84,"dew_point":268,"wind_speed":1.71,"wind_deg":161,"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}],"clouds":86,"pop":0,"uvi":2.03},{"dt":1607158800,"sunrise":1607142414,"sunset":1607176494,"temp":{"day":274.99,"min":266.85,"max":276.18,"night":269.38,"eve":269.56,"morn":267.05},"feels_like":{"day":272.4,"night":266.41,"eve":266.52,"morn":263.93},"pressure":1019,"humidity":77,"dew_point":266.76,"wind_speed":0.52,"wind_deg":146,"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":0,"pop":0,"uvi":1.93},{"dt":1607245200,"sunrise":1607228868,"sunset":1607262889,"temp":{"day":274.91,"min":267.08,"max":275.99,"night":269.18,"eve":269.41,"morn":267.08},"feels_like":{"day":272.09,"night":265.97,"eve":266.42,"morn":263.77},"pressure":1019,"humidity":76,"dew_point":265.75,"wind_speed":0.8,"wind_deg":286,"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":0,"pop":0,"uvi":1.75},{"dt":1607331600,"sunrise":1607315322,"sunset":1607349286,"temp":{"day":273.95,"min":266.79,"max":274.77,"night":269.42,"eve":269.44,"morn":266.79},"feels_like":{"day":270.5,"night":265.78,"eve":265.95,"morn":263.18},"pressure":1022,"humidity":78,"dew_point":263.67,"wind_speed":1.59,"wind_deg":159,"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":0,"pop":0,"uvi":2},{"dt":1607418000,"sunrise":1607401774,"sunset":1607435685,"temp":{"day":274.03,"min":267.03,"max":274.91,"night":269.79,"eve":269.62,"morn":267.03},"feels_like":{"day":270.3,"night":266.19,"eve":266.18,"morn":263.19},"pressure":1021,"humidity":75,"dew_point":262.33,"wind_speed":1.91,"wind_deg":156,"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":0,"pop":0,"uvi":2},{"dt":1607504400,"sunrise":1607488225,"sunset":1607522087,"temp":{"day":273.38,"min":268.54,"max":274.21,"night":271.07,"eve":271.13,"morn":268.54},"feels_like":{"day":269.67,"night":268.02,"eve":268.01,"morn":264.67},"pressure":1020,"humidity":81,"dew_point":263.92,"wind_speed":1.96,"wind_deg":152,"weather":[{"id":600,"main":"Snow","description":"light snow","icon":"13d"}],"clouds":100,"pop":0.32,"snow":1.06,"uvi":2}]}`)

func TestGetWeather_WhenOpenWeatherReturnsError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	var city, country = "city", "country"
	path := fmt.Sprintf(openWeatherCurrentWeatherURLFmt, city, country, testAPPID)
	expectedErr := errors.New("expected err")
	httpmock.RegisterResponder("GET", path, httpmock.NewErrorResponder(expectedErr))

	client := NewOpenWeatherMapClient(testAPPID)
	weather, err := client.GetCurrentWeather(city, country)
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
	weather, err := client.GetCurrentWeather(city, country)
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
	weather, err := client.GetCurrentWeather(city, country)
	assert.Nil(t, err)
	assert.NotNil(t, weather)

	resp := &OpenWeatherRespXML{}
	assert.Nil(t, xml.Unmarshal(openWeatherValidResponseCurrentWeather, resp))
	result, err := toWeatherDomain(resp)
	assert.Nil(t, err)
	assert.Equal(t, result, weather)
}

func TestGetForecast_WhenOpenWeatherReturnsError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	var lat, lon = "5.5", "5.5"
	var day int64 = 5
	path := fmt.Sprintf(openWeatherForecastURLFmt, lat, lon, testAPPID)
	expectedErr := errors.New("expected err")
	httpmock.RegisterResponder("GET", path, httpmock.NewErrorResponder(expectedErr))

	client := NewOpenWeatherMapClient(testAPPID)
	forecast, err := client.GetForecast("lat", "lon", day)
	assert.Error(t, err)
	assert.Nil(t, forecast)
}

func TestGetForecast_WhenOpenWeatherReturnsUnexpectedJSON(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	var lat, lon = "5.5", "5.5"
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
	var lat, lon = "5.5", "5.5"
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
	weather := &OpenWeatherRespXML{
		XMLName: xml.Name{
			Space: "",
			Local: "",
		},
		Text: "",
		City: City{
			Text: "",
			ID:   "123",
			Name: "city",
			Coord: Coord{
				Lon: "55",
				Lat: "55",
			},
			Country:  "country",
			Timezone: "",
			Sun: Sun{
				Rise: "2020-12-03T10:50:29",
				Set:  "2020-12-03T10:50:29",
			},
		},
		Temperature: Temperature{
			Text:  "",
			Value: "44",
			Unit:  "celsius",
		},
		Humidity: Humidity{
			Value: "12",
			Unit:  "%",
		},
		Pressure: Pressure{
			Value: "31",
			Unit:  "hPa",
		},
		Wind: Wind{
			Speed: Speed{
				Name: "light breeze",
			},
			Direction: Direction{
				Name: "south",
			},
		},
		Clouds: Clouds{
			Name: "broken clouds",
		},
	}

	tsr, _ := time.Parse("2006-01-02T15:04:05", weather.City.Sun.Rise)
	tss, _ := time.Parse("2006-01-02T15:04:05", weather.City.Sun.Rise)
	expected := &domain.CurrentWeather{
		Country:     "country",
		City:        "city",
		Cloudiness:  "broken clouds",
		Humidity:    domain.Humidity{Value: "12", Unit: "%"},
		Pressure:    domain.Pressure{Value: "31", Unit: "hPa"},
		Sunrise:     tsr,
		Sunset:      tss,
		Temperature: domain.Temperature{Value: "44", Unit: "celsius"},
		Wind:        domain.Wind{Speed: "", Unit: "", Name: "light breeze", DirectionName: "south"},
		Lat:         "55",
		Lon:         "55",
	}

	result, err := toWeatherDomain(weather)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

}

func TestToWeatherDomain_WhenSunriseTimeIsIncorrect_ReturnsError(t *testing.T) {
	weather := &OpenWeatherRespXML{
		City: City{
			Text: "",
			ID:   "123",
			Name: "city",
			Coord: Coord{
				Lon: "55",
				Lat: "55",
			},
			Country:  "country",
			Timezone: "",
			Sun: Sun{
				Rise: "2020:12:03T10:50:29",
				Set:  "2020:12:03T10:50:29",
			},
		},
	}

	_, err := toWeatherDomain(weather)
	assert.Error(t, err)
}
