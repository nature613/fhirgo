package resources

import (
	d "github.com/monarko/fhirgo/STU3/datatypes"
	"github.com/monarko/fhirgo/schema"
)

// Observation resource
type Observation struct {
	Domain
	Identifier           []d.Identifier                `json:"identifier,omitempty"`
	BasedOn              []d.Reference                 `json:"basedOn,omitempty"`
	Status               *d.Code                       `json:"status,omitempty"`
	Category             []d.CodeableConcept           `json:"category,omitempty"`
	Code                 *d.CodeableConcept            `json:"code,omitempty"`
	Subject              *d.Reference                  `json:"subject,omitempty"`
	Context              *d.Reference                  `json:"context,omitempty"`
	EffectiveDateTime    *d.DateTime                   `json:"effectiveDateTime"`
	EffectivePeriod      *d.Period                     `json:"effectivePeriod,omitempty"`
	Issued               *d.Instant                    `json:"issued,omitempty"`
	Performer            []d.Reference                 `json:"performer,omitempty"`
	ValueQuantity        *d.Quantity                   `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *d.CodeableConcept            `json:"valueCodeableConcept,omitempty"`
	ValueString          *d.String                     `json:"valueString,omitempty"`
	ValueBoolean         *d.Boolean                    `json:"valueBoolean,omitempty"`
	ValueRange           *d.Range                      `json:"valueRange,omitempty"`
	ValueRatio           *d.Ratio                      `json:"valueRatio,omitempty"`
	ValueSampledData     *d.SampledData                `json:"valueSampledData,omitempty"`
	ValueAttachment      *d.Attachment                 `json:"valueAttachment,omitempty"`
	ValueTime            *d.Time                       `json:"valueTime,omitempty"`
	ValueDateTime        *d.DateTime                   `json:"valueDateTime,omitempty"`
	ValuePeriod          *d.Period                     `json:"valuePeriod,omitempty"`
	DataAbsentReason     *d.CodeableConcept            `json:"dataAbsentReason,omitempty"`
	Interpretation       *d.CodeableConcept            `json:"interpretation,omitempty"`
	Comment              *d.String                     `json:"comment,omitempty"`
	BodySite             *d.CodeableConcept            `json:"bodySite,omitempty"`
	Method               *d.CodeableConcept            `json:"method,omitempty"`
	Specimen             *d.Reference                  `json:"specimen,omitempty"`
	Device               *d.Reference                  `json:"device,omitempty"`
	ReferenceRange       []d.ObservationReferenceRange `json:"referenceRange,omitempty"`
	Related              []d.ObservationRelated        `json:"related,omitempty"`
	Component            []d.ObservationComponent      `json:"component,omitempty"`
}

// Validate returns a check against schema
func (o *Observation) Validate() (bool, []error) {
	return schema.ValidateResource(*o, "3")
}
