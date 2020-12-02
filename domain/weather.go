package domain

type WeatherWithForecast struct {
	Cloudiness string
	Forecast *Forecast
	GeoCoordinates string
	Humidity string
	LocationName string
	Pressure string
	RequestedTime string
	Sunrise string
	Sunset string
	Temperature string
	Wind string
	Lat float64
	Lon float64
}
