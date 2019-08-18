package resources

import (
	d "github.com/monarko/fhirgo/R4/datatypes"
)

// Base Resource
type Base struct {
	ResourceType  string  `json:"resourceType"`
	ID            *d.ID   `json:"id,omitempty"`
	Meta          *d.Meta `json:"meta,omitempty"`
	ImplicitRules *d.URI  `json:"implicitRules,omitempty"`
	Language      *d.Code `json:"language,omitempty"`
}
