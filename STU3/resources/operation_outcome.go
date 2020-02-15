package resources

import (
    d "github.com/monarko/fhirgo/STU3/datatypes"
    "github.com/monarko/fhirgo/schema"
)

// OperationOutcome resource
type OperationOutcome struct {
    Domain
    Issue []OperationOutcomeIssue `json:"issue,omitempty"`
}

// Validate returns a check against schema
func (a *OperationOutcome) Validate() (bool, []error) {
    return schema.ValidateResource(*a, "3")
}

// OperationOutcomeIssue sub-resource
type OperationOutcomeIssue struct {
    Severity    *d.Code            `json:"severity,omitempty"`
    Code        *d.Code            `json:"code,omitempty"`
    Details     *d.CodeableConcept `json:"details,omitempty"`
    Diagnostics *d.String          `json:"diagnostics,omitempty"`
    Location    []d.String         `json:"location,omitempty"`
    Expression  []d.String         `json:"expression,omitempty"`
}
