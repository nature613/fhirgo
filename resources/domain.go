package resources

import (
	d "github.com/monarko/fhirgo/datatypes"
)

// Domain resource
type Domain struct {
	Base
	Text      *d.Narrative `json:"text,omitempty"`
	Contained []Base       `json:"contained,omitempty"`
}

// NewDomain returns an empty domain resource with resourceType
func NewDomain(resourceType string) (Domain, error) {
	b, err := NewBase(resourceType)
	if err != nil {
		return Domain{}, nil
	}
	return Domain{b, nil, nil}, nil
}
