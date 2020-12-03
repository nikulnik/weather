package rest

// OpenWeatherErrorResp represents the structure of errors returned by Openweathermap API
type OpenWeatherErrorResp struct {
	Cod     int    `json:"cod"`
	Message string `json:"message"`
}
