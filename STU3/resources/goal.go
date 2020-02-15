package resources

import (
    d "github.com/monarko/fhirgo/STU3/datatypes"
    "github.com/monarko/fhirgo/schema"
)

// Goal resource
type Goal struct {
    Domain
    Identifier           []d.Identifier      `json:"identifier,omitempty"`
    Status               *d.Code             `json:"status,omitempty"`
    Category             []d.CodeableConcept `json:"category,omitempty"`
    Priority             *d.CodeableConcept  `json:"priority,omitempty"`
    Description          *d.CodeableConcept  `json:"description,omitempty"`
    Subject              *d.Reference        `json:"subject,omitempty"`
    StartDate            *d.Date             `json:"startDate,omitempty"`
    StartCodeableConcept *d.CodeableConcept  `json:"startCodeableConcept,omitempty"`
    Target               *GoalTarget         `json:"target,omitempty"`
    StatusDate           *d.Date             `json:"statusDate,omitempty"`
    StatusReason         *d.String           `json:"statusReason,omitempty"`
    ExpressedBy          *d.Reference        `json:"expressedBy,omitempty"`
    Addresses            []d.Reference       `json:"addresses,omitempty"`
    Note                 []d.Annotation      `json:"note,omitempty"`
    OutcomeCode          []d.CodeableConcept `json:"outcomeCode,omitempty"`
    OutcomeReference     []d.Reference       `json:"outcomeReference,omitempty"`
}

// Validate returns a check against schema
func (g *Goal) Validate() (bool, []error) {
    return schema.ValidateResource(*g, "3")
}

// GoalTarget subResource
type GoalTarget struct {
    Measure               *d.CodeableConcept `json:"measure,omitempty"`
    DetailQuantity        *d.Quantity        `json:"detailQuantity,omitempty"`
    DetailRange           *d.Range           `json:"detailRange,omitempty"`
    DetailCodeableConcept *d.CodeableConcept `json:"detailCodeableConcept,omitempty"`
    DueDate               *d.Date            `json:"dueDate,omitempty"`
    DueDuration           *d.Duration        `json:"dueDuration,omitempty"`
}
