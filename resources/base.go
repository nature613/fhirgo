package resources

import (
	"strings"

	d "github.com/monarko/fhirgo/datatypes"
)

// Base Resource
type Base struct {
	ResourceType  string  `json:"resourceType"`
	ID            *d.ID   `json:"id,omitempty"`
	Meta          *d.Meta `json:"meta,omitempty"`
	ImplicitRules *d.URI  `json:"implicitRules,omitempty"`
	Language      *d.Code `json:"language,omitempty"`
}

// NewBase returns a empty Base with resourceType
func NewBase(resourceType string) (Base, error) {
	return Base{ResourceType: strings.Title(resourceType)}, nil
}
