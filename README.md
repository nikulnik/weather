# Start the application

1. [Install](https://github.com/go-swagger/go-swagger#installing) swagger.
2. Run from the root of the repository directory : `swagger generate server [-f ./swagger.yaml]`
3. Set OPENWEATHERMAP_KEY env to your openweathermap.org key.
4. `go run cmd/weather-api-server/main.go`
5. The app. address and port will be shown in your console.