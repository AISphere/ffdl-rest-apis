// Code generated by go-swagger; DO NOT EDIT.

package training_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/AISphere/ffdl-rest-apis/api_v1/restmodels"
)

// GetEMetricsOKCode is the HTTP code returned for type GetEMetricsOK
const GetEMetricsOKCode int = 200

/*GetEMetricsOK (streaming responses)

swagger:response getEMetricsOK
*/
type GetEMetricsOK struct {

	/*
	  In: Body
	*/
	Payload *restmodels.V1EMetricsList `json:"body,omitempty"`
}

// NewGetEMetricsOK creates GetEMetricsOK with default headers values
func NewGetEMetricsOK() *GetEMetricsOK {
	return &GetEMetricsOK{}
}

// WithPayload adds the payload to the get e metrics o k response
func (o *GetEMetricsOK) WithPayload(payload *restmodels.V1EMetricsList) *GetEMetricsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get e metrics o k response
func (o *GetEMetricsOK) SetPayload(payload *restmodels.V1EMetricsList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEMetricsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEMetricsUnauthorizedCode is the HTTP code returned for type GetEMetricsUnauthorized
const GetEMetricsUnauthorizedCode int = 401

/*GetEMetricsUnauthorized Unauthorized

swagger:response getEMetricsUnauthorized
*/
type GetEMetricsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewGetEMetricsUnauthorized creates GetEMetricsUnauthorized with default headers values
func NewGetEMetricsUnauthorized() *GetEMetricsUnauthorized {
	return &GetEMetricsUnauthorized{}
}

// WithPayload adds the payload to the get e metrics unauthorized response
func (o *GetEMetricsUnauthorized) WithPayload(payload *restmodels.Error) *GetEMetricsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get e metrics unauthorized response
func (o *GetEMetricsUnauthorized) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEMetricsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
