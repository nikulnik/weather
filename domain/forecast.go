package domain

type Forecast struct {
	Lat            int     `json:"lat"`
	Lon            int     `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset int     `json:"timezone_offset"`
	Daily          []Daily `json:"daily"`
}

type Temp struct {
	Day   float64 `json:"day"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

type FeelsLike struct {
	Day   float64 `json:"day"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

type Daily struct {
	Dt        int       `json:"dt"`
	Sunrise   int       `json:"sunrise"`
	Sunset    int       `json:"sunset"`
	Temp      Temp      `json:"temp"`
	FeelsLike FeelsLike `json:"feels_like"`
	Pressure  int       `json:"pressure"`
	Humidity  int       `json:"humidity"`
	DewPoint  float64   `json:"dew_point"`
	WindSpeed float64   `json:"wind_speed"`
	WindDeg   float64   `json:"wind_deg"`
	Weather   []Weather `json:"weather"`
	Clouds    int       `json:"clouds"`
	Pop       float64   `json:"pop"`
	Uvi       float64   `json:"uvi"`
}
