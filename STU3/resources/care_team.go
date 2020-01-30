package resources

import (
    d "github.com/monarko/fhirgo/STU3/datatypes"
    "github.com/monarko/fhirgo/schema"
)

// CareTeam resource
type CareTeam struct {
    Domain
    Identifier           []d.Identifier          `json:"identifier,omitempty"`
    Status               *d.Code                 `json:"status,omitempty"`
    Category             []d.CodeableConcept     `json:"category,omitempty"`
    Name                 *d.String               `json:"name,omitempty"`
    Subject              *d.Reference            `json:"subject,omitempty"`
    Context              *d.Reference            `json:"context,omitempty"`
    Period               *d.Period               `json:"period,omitempty"`
    Participant          []d.CareTeamParticipant `json:"participant,omitempty"`
    ReasonCode           []d.CodeableConcept     `json:"reasonCode,omitempty"`
    ReasonReference      []d.Reference           `json:"reasonReference,omitempty"`
    ManagingOrganization []d.Reference           `json:"managingOrganization,omitempty"`
    Note                 []d.Annotation          `json:"note,omitempty"`
}

// Validate returns a check against schema
func (p *CareTeam) Validate() (bool, []error) {
    return schema.ValidateResource(*p, "3")
}
