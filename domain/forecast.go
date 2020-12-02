package domain

type Forecast struct {

	// date
	Date string `json:"date,omitempty"`

	// humidity
	Humidity string `json:"humidity,omitempty"`

	// pressure
	Pressure string `json:"pressure,omitempty"`

	// sunrise
	Sunrise string `json:"sunrise,omitempty"`

	// sunset
	Sunset string `json:"sunset,omitempty"`

	// temperature
	Temperature string `json:"temperature,omitempty"`

	// wind
	Wind string `json:"wind,omitempty"`
}
