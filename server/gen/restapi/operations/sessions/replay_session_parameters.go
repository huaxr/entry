// Code generated by go-swagger; DO NOT EDIT.

package sessions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewReplaySessionParams creates a new ReplaySessionParams object
// no default values defined in spec.
func NewReplaySessionParams() ReplaySessionParams {

	return ReplaySessionParams{}
}

// ReplaySessionParams contains all the bound params for the replay session operation
// typically these are obtained from a http.Request
//
// swagger:parameters replaySession
type ReplaySessionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	SessionID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewReplaySessionParams() beforehand.
func (o *ReplaySessionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rSessionID, rhkSessionID, _ := route.Params.GetOK("session_id")
	if err := o.bindSessionID(rSessionID, rhkSessionID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ReplaySessionParams) bindSessionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("session_id", "path", "int64", raw)
	}
	o.SessionID = value

	return nil
}
