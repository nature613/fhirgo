package resources

import (
	d "github.com/monarko/fhirgo/STU3/datatypes"
	"github.com/monarko/fhirgo/schema"
)

// Specimen resource
type Specimen struct {
	Domain
	Identifier          []d.Identifier         `json:"identifier,omitempty"`
	AccessionIdentifier *d.Identifier          `json:"accessionIdentifier,omitempty"`
	Status              *d.Code                `json:"status,omitempty"`
	Type                *d.CodeableConcept     `json:"type,omitempty"`
	Subject             *d.Reference           `json:"subject,omitempty"`
	ReceivedTime        *d.DateTime            `json:"receivedTime,omitempty"`
	Parent              []d.Reference          `json:"parent,omitempty"`
	Request             []d.Reference          `json:"request,omitempty"`
	Collection          *d.SpecimenCollection  `json:"collection,omitempty"`
	Processing          []d.SpecimenProcessing `json:"processing,omitempty"`
	Container           []d.SpecimenContainer  `json:"container,omitempty"`
	Note                []d.Annotation         `json:"note,omitempty"`
}

// Validate returns a check against schema
func (s *Specimen) Validate() (bool, []error) {
	return schema.ValidateResource(*s, "3")
}
