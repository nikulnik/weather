package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
	"github.com/nikulnik/weather/handlers"
	"github.com/nikulnik/weather/interactors"
	"github.com/nikulnik/weather/rest"
	"log"
	"net/http"
)

type Settings struct {
	AppPort           string `env:"APP_PORT"`
	OpenWeatherMapKey string `env:"OPENWEATHERMAP_KEY"`
}

func main() {
	settings := &Settings{}
	if err := env.Parse(settings); err != nil {
		fmt.Printf("%+v\n", err)
	}

	openWeatherClient := rest.NewOpenWeatherMapClient(settings.OpenWeatherMapKey)
	wi := interactors.NewWeatherInteractor(openWeatherClient)
	wh := handlers.NewWeatherHandler(wi)
	r := mux.NewRouter()
	r.HandleFunc("/weather", wh.GetWeather).Queries("city", "{city}", "country", "{country}")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+settings.AppPort, nil))
}
