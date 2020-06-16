package resources

import (
    d "github.com/monarko/fhirgo/STU3/datatypes"
    "github.com/monarko/fhirgo/schema"
)

// AllergyIntolerance resource
type AllergyIntolerance struct {
    Domain
    Identifier         []d.Identifier               `json:"identifier,omitempty"`
    ClinicalStatus     *d.Code                      `json:"clinicalStatus,omitempty"`
    VerificationStatus *d.Code                      `json:"verificationStatus,omitempty"`
    Type               *d.Code                      `json:"type,omitempty"`
    Category           []d.Code                     `json:"category,omitempty"`
    Criticality        *d.Code                      `json:"criticality,omitempty"`
    Code               *d.CodeableConcept           `json:"code,omitempty"`
    Patient            *d.Reference                 `json:"patient,omitempty"`
    OnsetDateTime      *d.DateTime                  `json:"onsetDateTime,omitempty"`
    OnsetAge           *d.Age                       `json:"onsetAge,omitempty"`
    OnsetPeriod        *d.Period                    `json:"onsetPeriod,omitempty"`
    OnsetRange         *d.Range                     `json:"onsetRange,omitempty"`
    OnsetString        *d.String                    `json:"onsetString,omitempty"`
    AssertedDate       *d.DateTime                  `json:"assertedDate,omitempty"`
    Recorder           *d.Reference                 `json:"recorder,omitempty"`
    Asserter           *d.Reference                 `json:"asserter,omitempty"`
    LastOccurrence     *d.DateTime                  `json:"lastOccurrence,omitempty"`
    Note               []d.Annotation               `json:"note,omitempty"`
    Reaction           []AllergyIntoleranceReaction `json:"reaction,omitempty"`
}

// Validate returns a check against schema
func (a *AllergyIntolerance) Validate() (bool, []error) {
    return schema.ValidateResource(*a, "3")
}

// AllergyIntoleranceReaction subResource
type AllergyIntoleranceReaction struct {
    Substance     *d.CodeableConcept  `json:"substance,omitempty"`
    Manifestation []d.CodeableConcept `json:"manifestation,omitempty"`
    Description   *d.String           `json:"description,omitempty"`
    Onset         *d.DateTime         `json:"onset,omitempty"`
    Severity      *d.Code             `json:"severity,omitempty"`
    ExposureRoute *d.CodeableConcept  `json:"exposureRoute,omitempty"`
    Note          []d.Annotation      `json:"note,omitempty"`
}
