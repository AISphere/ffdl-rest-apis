// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetEventTypeEndpointsParams creates a new GetEventTypeEndpointsParams object
// with the default values initialized.
func NewGetEventTypeEndpointsParams() GetEventTypeEndpointsParams {
	var ()
	return GetEventTypeEndpointsParams{}
}

// GetEventTypeEndpointsParams contains all the bound params for the get event type endpoints operation
// typically these are obtained from a http.Request
//
// swagger:parameters getEventTypeEndpoints
type GetEventTypeEndpointsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The type of event.
	  Required: true
	  In: path
	*/
	EventType string
	/*The id of the model.
	  Required: true
	  In: path
	*/
	ModelID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetEventTypeEndpointsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rEventType, rhkEventType, _ := route.Params.GetOK("event_type")
	if err := o.bindEventType(rEventType, rhkEventType, route.Formats); err != nil {
		res = append(res, err)
	}

	rModelID, rhkModelID, _ := route.Params.GetOK("model_id")
	if err := o.bindModelID(rModelID, rhkModelID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEventTypeEndpointsParams) bindEventType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.EventType = raw

	return nil
}

func (o *GetEventTypeEndpointsParams) bindModelID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ModelID = raw

	return nil
}