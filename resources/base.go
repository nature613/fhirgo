package resources

import (
	d "github.com/monarko/fhirgo/datatypes"
)

// Base Resource
type Base struct {
	ResourceType  string  `json:"resourceType"`
	ID            *d.ID   `json:"id"`
	Meta          *d.Meta `json:"meta"`
	ImplicitRules *d.URI  `json:"implicitRules"`
	Language      *d.Code `json:"language"`
}

// NewBase returns a empty Base with resourceType
func NewBase(resourceType string) (Base, error) {
	return Base{ResourceType: resourceType}, nil
}
