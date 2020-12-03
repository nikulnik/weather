package domain

import "time"

type CurrentWeather struct {
	Country     string
	City        string
	Cloudiness  string
	Humidity    Humidity
	Pressure    Pressure
	Sunrise     time.Time
	Sunset      time.Time
	Temperature Temperature
	Wind        Wind
	Lat         string
	Lon         string
}

type Wind struct {
	Speed         string
	Unit          string
	Name          string
	DirectionName string
}

type Temperature struct {
	Value string
	Unit  string
}
type Pressure struct {
	Value string
	Unit  string
}
type Humidity struct {
	Value string
	Unit  string
}
