// Code generated by go-swagger; DO NOT EDIT.

package restmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MetricData metric data
// swagger:model MetricData

type MetricData struct {

	// map of key/values, that describe evaluation metrics
	Values map[string]interface{} `json:"Values,omitempty"`

	// Current iteration number be processed.
	Iteration int32 `json:"iteration,omitempty"`

	// Timestamp of the metric. Format: yyyy-MM-dd'T'HH:mm:ss.SSS'Z'
	//
	Timestamp string `json:"timestamp,omitempty"`

	// The type of metrics data
	Type string `json:"type,omitempty"`
}

/* polymorph MetricData Values false */

/* polymorph MetricData iteration false */

/* polymorph MetricData timestamp false */

/* polymorph MetricData type false */

// Validate validates this metric data
func (m *MetricData) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *MetricData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricData) UnmarshalBinary(b []byte) error {
	var res MetricData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
