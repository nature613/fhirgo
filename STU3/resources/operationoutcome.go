package resources

import (
    d "github.com/monarko/fhirgo/STU3/datatypes"
    "github.com/monarko/fhirgo/schema"
)

// OperationOutcome resource
type OperationOutcome struct {
    Domain
    Issue []d.OperationOutcomeIssue `json:"issue,omitempty"`
}

// Validate returns a check against schema
func (a *OperationOutcome) Validate() (bool, []error) {
    return schema.ValidateResource(*a, "3")
}
