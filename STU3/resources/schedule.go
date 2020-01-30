package resources

import (
    d "github.com/monarko/fhirgo/STU3/datatypes"
    "github.com/monarko/fhirgo/schema"
)

// Schedule resource
type Schedule struct {
    Domain
    Identifier      []d.Identifier      `json:"identifier,omitempty"`
    Active          *d.Boolean          `json:"active,omitempty"`
    ServiceCategory *d.CodeableConcept  `json:"serviceCategory,omitempty"`
    ServiceType     []d.CodeableConcept `json:"serviceType,omitempty"`
    Specialty       []d.CodeableConcept `json:"specialty,omitempty"`
    Actor           []d.Reference       `json:"actor,omitempty"`
    PlanningHorizon *d.Period           `json:"planningHorizon,omitempty"`
    Comment         *d.String           `json:"comment,omitempty"`
}

// Validate returns a check against schema
func (s *Schedule) Validate() (bool, []error) {
    return schema.ValidateResource(*s, "3")
}
