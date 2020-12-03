package rest

type OpenWeatherErrorResp struct {
	Cod     int    `json:"cod"`
	Message string `json:"message"`
}
