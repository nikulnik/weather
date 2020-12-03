// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/caarlos0/env"
	"github.com/nikulnik/weather/cache"
	"github.com/nikulnik/weather/handlers"
	"github.com/nikulnik/weather/interactors"
	"github.com/nikulnik/weather/rest"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/nikulnik/weather/restapi/operations"
	"github.com/nikulnik/weather/restapi/operations/weather"
)

//go:generate swagger generate server --target ..\..\weather --name WeatherAPI --spec ..\swagger.yaml --principal interface{}

type Settings struct {
	OpenWeatherMapKey string `env:"OPENWEATHERMAP_KEY"`
}

func configureFlags(api *operations.WeatherAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.WeatherAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	settings := &Settings{}
	if err := env.Parse(settings); err != nil {
		fmt.Printf("%+v\n", err)
	}
	cache := cache.NewCache(time.Second * 10)
	openWeatherClient := rest.NewOpenWeatherMapClient(settings.OpenWeatherMapKey)
	wi := interactors.NewWeatherInteractor(openWeatherClient, cache)
	wh := handlers.NewWeatherHandler(wi)

	api.WeatherGetWeatherHandler = weather.GetWeatherHandlerFunc(func(params weather.GetWeatherParams) middleware.Responder {
		return wh.GetWeather(params)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
