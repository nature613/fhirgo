package resources

import (
	d "github.com/monarko/fhirgo/datatypes"
)

// Narrative for domain resource
type Narrative struct {
	Status *d.Code   `json:"status"`
	Div    *d.String `json:"div"`
}

// Domain resource
type Domain struct {
	Base
	Text      *Narrative `json:"text"`
	Contained []Base     `json:"contained"`
}

// NewDomain returns an empty domain resource with resourceType
func NewDomain(resourceType string) (Domain, error) {
	b, err := NewBase(resourceType)
	if err != nil {
		return Domain{}, nil
	}
	return Domain{b, nil, nil}, nil
}
