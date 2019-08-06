package resources

import (
	d "github.com/monarko/fhirgo/datatypes"
	"github.com/monarko/fhirgo/schema"
)

// Observation resource
type Observation struct {
	Domain
	Identifier           []d.Identifier                `json:"identifier,omitempty"`
	BasedOn              *d.Reference                  `json:"basedOn,omitempty"`
	PartOf               *d.Reference                  `json:"partOf,omitempty"`
	Status               *d.Code                       `json:"status,omitempty"`
	Category             *d.CodeableConcept            `json:"category,omitempty"`
	Code                 *d.CodeableConcept            `json:"code,omitempty"`
	Subject              *d.Reference                  `json:"subject,omitempty"`
	Focus                *d.Reference                  `json:"focus,omitempty"`
	Encounter            *d.Reference                  `json:"encounter,omitempty"`
	EffectiveDateTime    *d.DateTime                   `json:"effectiveDateTime"`
	EffectivePeriod      *d.Period                     `json:"effectivePeriod,omitempty"`
	EffectiveTiming      *d.Timing                     `json:"effectiveTiming,omitempty"`
	EffectiveInstant     *d.Instant                    `json:"effectiveInstant,omitempty"`
	Issued               *d.Instant                    `json:"issued,omitempty"`
	Performer            *d.Reference                  `json:"performer,omitempty"`
	ValueQuantity        *d.Quantity                   `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *d.CodeableConcept            `json:"valueCodeableConcept,omitempty"`
	ValueString          *d.String                     `json:"valueString,omitempty"`
	ValueBoolean         *d.Boolean                    `json:"valueBoolean,omitempty"`
	ValueInteger         *d.Integer                    `json:"valueInteger,omitempty"`
	ValueRange           *d.Range                      `json:"valueRange,omitempty"`
	ValueRatio           *d.Ratio                      `json:"valueRatio,omitempty"`
	ValueSampledData     *d.SampledData                `json:"valueSampledData,omitempty"`
	ValueTime            *d.Time                       `json:"valueTime,omitempty"`
	ValueDateTime        *d.DateTime                   `json:"valueDateTime,omitempty"`
	ValuePeriod          *d.Period                     `json:"valuePeriod,omitempty"`
	DataAbsentReason     *d.CodeableConcept            `json:"dataAbsentReason,omitempty"`
	Interpretation       []d.CodeableConcept           `json:"interpretation,omitempty"`
	Note                 []d.Annotation                `json:"note,omitempty"`
	BodySite             *d.CodeableConcept            `json:"bodySite,omitempty"`
	Method               *d.CodeableConcept            `json:"method,omitempty"`
	Specimen             *d.Reference                  `json:"specimen,omitempty"`
	Device               *d.Reference                  `json:"device,omitempty"`
	ReferenceRange       []d.ObservationReferenceRange `json:"referenceRange,omitempty"`
	HasMember            []d.Reference                 `json:"hasMember,omitempty"`
	DerivedFrom          []d.Reference                 `json:"derivedFrom,omitempty"`
	Component            []d.ObservationComponent      `json:"component,omitempty"`
}

// Validate returns a check against schema
func (o *Observation) Validate() (bool, []error) {
	return schema.ValidateResource(*o, "4.0")
}
