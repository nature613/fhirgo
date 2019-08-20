package resources

import (
	d "github.com/monarko/fhirgo/STU3/datatypes"
	"github.com/monarko/fhirgo/schema"
)

// RiskAssessment resource
type RiskAssessment struct {
	Domain
	Identifier            *d.Identifier                `json:"identifier,omitempty"`
	BasedOn               *d.Reference                 `json:"basedOn,omitempty"`
	Parent                *d.Reference                 `json:"parent,omitempty"`
	Status                *d.Code                      `json:"status,omitempty"`
	Method                *d.CodeableConcept           `json:"method,omitempty"`
	Code                  *d.CodeableConcept           `json:"code,omitempty"`
	Subject               *d.Reference                 `json:"subject,omitempty"`
	Context               *d.Reference                 `json:"context,omitempty"`
	OccurrenceDateTime    *d.DateTime                  `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod      *d.Period                    `json:"occurrencePeriod,omitempty"`
	Condition             *d.Reference                 `json:"condition,omitempty"`
	Performer             *d.Reference                 `json:"performer,omitempty"`
	ReasonCodeableConcept *d.CodeableConcept           `json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *d.Reference                 `json:"reasonReference,omitempty"`
	Basis                 []d.Reference                `json:"basis,omitempty"`
	Prediction            []d.RiskAssessmentPrediction `json:"prediction,omitempty"`
	Mitigation            *d.String                    `json:"mitigation,omitempty"`
	Comment               *d.String                    `json:"comment,omitempty"`
}

// Validate returns a check against schema
func (r *RiskAssessment) Validate() (bool, []error) {
	return schema.ValidateResource(*r, "3")
}
