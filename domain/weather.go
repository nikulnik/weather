package domain

type WeatherWithForecast struct {

	// cloudiness
	Cloudiness string

	// forecast
	Forecast *Forecast

	// geo coordinates
	GeoCoordinates string

	// humidity
	Humidity string

	// location name
	LocationName string

	// pressure
	Pressure string

	// requested time
	RequestedTime string

	// sunrise
	Sunrise string

	// sunset
	Sunset string

	// temperature
	Temperature string

	// wind
	Wind string

	Lat float64

	Lon float64
}
