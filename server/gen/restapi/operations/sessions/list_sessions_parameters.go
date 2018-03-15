// Code generated by go-swagger; DO NOT EDIT.

package sessions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListSessionsParams creates a new ListSessionsParams object
// with the default values initialized.
func NewListSessionsParams() ListSessionsParams {

	var (
		// initialize parameters with default values

		limitDefault  = int64(20)
		offsetDefault = int64(0)
		sinceDefault  = int64(0)
	)

	return ListSessionsParams{
		Limit: &limitDefault,

		Offset: &offsetDefault,

		Since: &sinceDefault,
	}
}

// ListSessionsParams contains all the bound params for the list sessions operation
// typically these are obtained from a http.Request
//
// swagger:parameters listSessions
type ListSessionsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Cookie with access_token
	  Required: true
	  In: header
	*/
	Cookie string
	/*MySQL LIKE pattern match
	  In: query
	*/
	AppName *string
	/*
	  In: query
	  Default: 20
	*/
	Limit *int64
	/*
	  In: query
	  Default: 0
	*/
	Offset *int64
	/*Unix timestamp(unit: second)
	  In: query
	  Default: 0
	*/
	Since *int64
	/*MySQL LIKE pattern match
	  In: query
	*/
	User *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListSessionsParams() beforehand.
func (o *ListSessionsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindCookie(r.Header[http.CanonicalHeaderKey("Cookie")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	qAppName, qhkAppName, _ := qs.GetOK("app_name")
	if err := o.bindAppName(qAppName, qhkAppName, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qOffset, qhkOffset, _ := qs.GetOK("offset")
	if err := o.bindOffset(qOffset, qhkOffset, route.Formats); err != nil {
		res = append(res, err)
	}

	qSince, qhkSince, _ := qs.GetOK("since")
	if err := o.bindSince(qSince, qhkSince, route.Formats); err != nil {
		res = append(res, err)
	}

	qUser, qhkUser, _ := qs.GetOK("user")
	if err := o.bindUser(qUser, qhkUser, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListSessionsParams) bindCookie(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("Cookie", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("Cookie", "header", raw); err != nil {
		return err
	}

	o.Cookie = raw

	return nil
}

func (o *ListSessionsParams) bindAppName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.AppName = &raw

	return nil
}

func (o *ListSessionsParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewListSessionsParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	return nil
}

func (o *ListSessionsParams) bindOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewListSessionsParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("offset", "query", "int64", raw)
	}
	o.Offset = &value

	return nil
}

func (o *ListSessionsParams) bindSince(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewListSessionsParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("since", "query", "int64", raw)
	}
	o.Since = &value

	return nil
}

func (o *ListSessionsParams) bindUser(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.User = &raw

	return nil
}
