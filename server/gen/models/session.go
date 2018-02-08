// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Session session
// swagger:model session
type Session struct {

	// app name
	AppName string `json:"app_name,omitempty"`

	// container id
	ContainerID string `json:"container_id,omitempty"`

	// Unix timestamp(unit: second)
	CreatedAt int64 `json:"created_at,omitempty"`

	// Unix timestamp(unit: second)
	EndedAt int64 `json:"ended_at,omitempty"`

	// instance no
	InstanceNo string `json:"instance_no,omitempty"`

	// node ip
	NodeIP string `json:"node_ip,omitempty"`

	// proc name
	ProcName string `json:"proc_name,omitempty"`

	// session id
	// Read Only: true
	SessionID int64 `json:"session_id,omitempty"`

	// session type
	SessionType string `json:"session_type,omitempty"`

	// source ip
	SourceIP string `json:"source_ip,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// user
	User string `json:"user,omitempty"`
}

// Validate validates this session
func (m *Session) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Session) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Session) UnmarshalBinary(b []byte) error {
	var res Session
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
