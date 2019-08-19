package resources

import (
	d "github.com/monarko/fhirgo/STU3/datatypes"
)

// Domain resource
type Domain struct {
	Base
	Text      *d.Narrative `json:"text,omitempty"`
	Contained []Base       `json:"contained,omitempty"`
}
