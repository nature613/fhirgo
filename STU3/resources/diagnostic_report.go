package resources

import (
	d "github.com/monarko/fhirgo/STU3/datatypes"
	"github.com/monarko/fhirgo/schema"
)

// DiagnosticReport resource
type DiagnosticReport struct {
	Domain
	Identifier        []d.Identifier                `json:"identifier,omitempty"`
	BasedOn           []d.Reference                 `json:"basedOn,omitempty"`
	Status            *d.Code                       `json:"status,omitempty"`
	Category          *d.CodeableConcept            `json:"category,omitempty"`
	Code              *d.CodeableConcept            `json:"code,omitempty"`
	Subject           *d.Reference                  `json:"subject,omitempty"`
	Context           *d.Reference                  `json:"context,omitempty"`
	EffectiveDateTime *d.DateTime                   `json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *d.Period                     `json:"effectivePeriod,omitempty"`
	Issued            *d.Instant                    `json:"issued,omitempty"`
	Performer         []d.DiagnosticReportPerformer `json:"performer,omitempty"`
	Specimen          []d.Reference                 `json:"specimen,omitempty"`
	Result            []d.Reference                 `json:"result,omitempty"`
	ImagingStudy      []d.Reference                 `json:"imagingStudy,omitempty"`
	Image             []d.DiagnosticReportImage     `json:"image,omitempty"`
	Conclusion        *d.String                     `json:"conclusion,omitempty"`
	CodedDiagnosis    []d.CodeableConcept           `json:"codedDiagnosis,omitempty"`
	PresentedForm     []d.Attachment                `json:"presentedForm,omitempty"`
}

// Validate returns a check against schema
func (d *DiagnosticReport) Validate() (bool, []error) {
	return schema.ValidateResource(*d, "3")
}
