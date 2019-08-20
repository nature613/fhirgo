package resources

import (
	d "github.com/monarko/fhirgo/STU3/datatypes"
	"github.com/monarko/fhirgo/schema"
)

// CarePlan resource
type CarePlan struct {
	Domain
	Identifier     []d.Identifier       `json:"identifier,omitempty"`
	Definition     []d.Reference        `json:"definition,omitempty"`
	BasedOn        []d.Reference        `json:"basedOn,omitempty"`
	Replaces       []d.Reference        `json:"replaces,omitempty"`
	PartOf         []d.Reference        `json:"partOf,omitempty"`
	Status         *d.Code              `json:"status,omitempty"`
	Intent         *d.Code              `json:"intent,omitempty"`
	Category       []d.CodeableConcept  `json:"category,omitempty"`
	Title          *d.String            `json:"title,omitempty"`
	Description    *d.String            `json:"description,omitempty"`
	Subject        *d.Reference         `json:"subject,omitempty"`
	Context        *d.Reference         `json:"context,omitempty"`
	Period         *d.Period            `json:"period,omitempty"`
	Author         []d.Reference        `json:"author,omitempty"`
	CareTeam       []d.Reference        `json:"careTeam,omitempty"`
	Addresses      []d.Reference        `json:"addresses,omitempty"`
	SupportingInfo []d.Reference        `json:"supportingInfo,omitempty"`
	Goal           []d.Reference        `json:"goal,omitempty"`
	Activity       []d.CarePlanActivity `json:"activity,omitempty"`
	Note           []d.Annotation       `json:"note,omitempty"`
}

// Validate returns a check against schema
func (c *CarePlan) Validate() (bool, []error) {
	return schema.ValidateResource(*c, "3")
}
