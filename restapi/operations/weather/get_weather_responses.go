// Code generated by go-swagger; DO NOT EDIT.

package weather

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/nikulnik/weather/models"
)

// GetWeatherOKCode is the HTTP code returned for type GetWeatherOK
const GetWeatherOKCode int = 200

/*GetWeatherOK Success

swagger:response getWeatherOK
*/
type GetWeatherOK struct {

	/*
	  In: Body
	*/
	Payload *models.WeatherWithForecast `json:"body,omitempty"`
}

// NewGetWeatherOK creates GetWeatherOK with default headers values
func NewGetWeatherOK() *GetWeatherOK {

	return &GetWeatherOK{}
}

// WithPayload adds the payload to the get weather o k response
func (o *GetWeatherOK) WithPayload(payload *models.WeatherWithForecast) *GetWeatherOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get weather o k response
func (o *GetWeatherOK) SetPayload(payload *models.WeatherWithForecast) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWeatherOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetWeatherDefault Unexpected error

swagger:response getWeatherDefault
*/
type GetWeatherDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetWeatherDefault creates GetWeatherDefault with default headers values
func NewGetWeatherDefault(code int) *GetWeatherDefault {
	if code <= 0 {
		code = 500
	}

	return &GetWeatherDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get weather default response
func (o *GetWeatherDefault) WithStatusCode(code int) *GetWeatherDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get weather default response
func (o *GetWeatherDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get weather default response
func (o *GetWeatherDefault) WithPayload(payload *models.Error) *GetWeatherDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get weather default response
func (o *GetWeatherDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWeatherDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
