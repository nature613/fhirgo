package resources

import (
	d "github.com/monarko/fhirgo/STU3/datatypes"
	"github.com/monarko/fhirgo/schema"
)

// GuidanceResponse resource
type GuidanceResponse struct {
	Domain
	RequestID             *d.ID               `json:"requestId,omitempty"`
	Identifier            *d.Identifier       `json:"identifier,omitempty"`
	Module                *d.Reference        `json:"module,omitempty"`
	Status                *d.Code             `json:"status,omitempty"`
	Subject               *d.Reference        `json:"subject,omitempty"`
	Context               *d.Reference        `json:"context,omitempty"`
	OccurrenceDateTime    *d.DateTime         `json:"occurrenceDateTime,omitempty"`
	Performer             *d.Reference        `json:"performer,omitempty"`
	ReasonCodeableConcept *d.CodeableConcept  `json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *d.Reference        `json:"reasonReference,omitempty"`
	Note                  []d.Annotation      `json:"note,omitempty"`
	EvaluationMessage     []d.Reference       `json:"evaluationMessage,omitempty"`
	OutputParameters      *d.Reference        `json:"outputParameters,omitempty"`
	Result                *d.Reference        `json:"result,omitempty"`
	DataRequirement       []d.DataRequirement `json:"dataRequirement,omitempty"`
}

// Validate returns a check against schema
func (r *GuidanceResponse) Validate() (bool, []error) {
	return schema.ValidateResource(*r, "3")
}
