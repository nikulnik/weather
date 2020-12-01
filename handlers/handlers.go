package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nikulnik/weather/interactors"
)

var (
	internalErr = `internal error`
)

type DefaultErrorResponse struct {
	Error string `json:"Error"`
}

type WeatherHandler interface {
	GetWeather(w http.ResponseWriter, r *http.Request)
}

func NewWeatherHandler(weatherIntercator interactors.WeatherInteractor) WeatherHandler {
	return &weatherHandler{weatherInteractor: weatherIntercator}
}

type weatherHandler struct {
	weatherInteractor interactors.WeatherInteractor
}

func (h *weatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	weatherData, err := h.weatherInteractor.GetWeather(vars["city"], vars["country"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(newErrorResponse(err.Error()))
		return
	}
	resp, err := json.Marshal(weatherData)
	w.Write(resp)
}

func newErrorResponse(text string) []byte {
	payload, _ := json.Marshal(&DefaultErrorResponse{Error: text})
	return payload
}
