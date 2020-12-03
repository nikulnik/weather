package domain

type Forecast struct {
	DateTime    int
	Humidity    int
	Pressure    int
	Sunrise     int
	Sunset      int
	Temperature float64
	WindSpeed   float64
	WindDegree  float64
}
