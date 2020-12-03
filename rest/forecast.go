package rest

// ForecastResp represents the forecast response of Openweathermap API
type ForecastResp struct {
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
type WeatherForecast struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
type Daily struct {
	Dt        int               `json:"dt"`
	Sunrise   int               `json:"sunrise"`
	Sunset    int               `json:"sunset"`
	Temp      Temp              `json:"temp"`
	FeelsLike FeelsLike         `json:"feels_like"`
	Pressure  int               `json:"pressure"`
	Humidity  int               `json:"humidity"`
	DewPoint  float64           `json:"dew_point"`
	WindSpeed float64           `json:"wind_speed"`
	WindDeg   float64           `json:"wind_deg"`
	Weather   []WeatherForecast `json:"weather"`
	Clouds    float64           `json:"clouds"`
	Pop       float64           `json:"pop"`
	Uvi       float64           `json:"uvi"`
	Rain      float64           `json:"rain"`
	Snow      float64           `json:"snow"`
}
