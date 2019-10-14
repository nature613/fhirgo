package resources

import (
	d "github.com/monarko/fhirgo/STU3/datatypes"
	"github.com/monarko/fhirgo/schema"
)

// ReferralRequest resource
type ReferralRequest struct {
	Domain
	Identifier         []d.Identifier              `json:"identifier,omitempty"`
	Definition         []d.Reference               `json:"definition,omitempty"`
	BasedOn            []d.Reference               `json:"basedOn,omitempty"`
	Replaces           []d.Reference               `json:"replaces,omitempty"`
	GroupIdentifier    *d.Identifier               `json:"groupIdentifier,omitempty"`
	Status             *d.Code                     `json:"status,omitempty"`
	Intent             *d.Code                     `json:"intent,omitempty"`
	Type               *d.CodeableConcept          `json:"type,omitempty"`
	Priority           *d.Code                     `json:"priority,omitempty"`
	ServiceRequested   []d.CodeableConcept         `json:"serviceRequested,omitempty"`
	Subject            *d.Reference                `json:"subject,omitempty"`
	Context            *d.Reference                `json:"context,omitempty"`
	OccurrenceDateTime *d.DateTime                 `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *d.Period                   `json:"occurrencePeriod,omitempty"`
	AuthoredOn         *d.DateTime                 `json:"authoredOn,omitempty"`
	Requester          *d.ReferralRequestRequester `json:"requester,omitempty"`
	Specialty          *d.CodeableConcept          `json:"specialty,omitempty"`
	Recipient          []d.Reference               `json:"recipient,omitempty"`
	ReasonCode         []d.CodeableConcept         `json:"reasonCode,omitempty"`
	ReasonReference    []d.Reference               `json:"reasonReference,omitempty"`
	Description        *d.String                   `json:"description,omitempty"`
	SupportingInfo     []d.Reference               `json:"supportingInfo,omitempty"`
	Note               []d.Annotation              `json:"note,omitempty"`
	RelevantHistory    []d.Reference               `json:"relevantHistory,omitempty"`
}

// Validate returns a check against schema
func (r *ReferralRequest) Validate() (bool, []error) {
	return schema.ValidateResource(*r, "3")
}
