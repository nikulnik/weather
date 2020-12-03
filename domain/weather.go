package domain

import "time"

type CurrentWeather struct {
	Country       string
	City          string
	Cloudiness    string
	Humidity      int
	Pressure      int
	RequestedTime time.Time
	Sunrise       int
	Sunset        int
	Temperature   float64
	WindSpeed     float64
	WindDegree    float64
	Lat           float64
	Lon           float64
}
