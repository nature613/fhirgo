package resources

import (
	d "github.com/monarko/fhirgo/STU3/datatypes"
	"github.com/monarko/fhirgo/schema"
)

// FamilyMemberHistory resource
type FamilyMemberHistory struct {
	Domain
	Identifier      []d.Identifier                   `json:"identifier,omitempty"`
	Definition      []d.Reference                    `json:"definition,omitempty"`
	Status          *d.Code                          `json:"status,omitempty"`
	NotDone         *d.Boolean                       `json:"notDone,omitempty"`
	NotDoneReason   *d.CodeableConcept               `json:"notDoneReason,omitempty"`
	Patient         *d.Reference                     `json:"patient,omitempty"`
	Date            *d.DateTime                      `json:"date,omitempty"`
	Name            *d.String                        `json:"name,omitempty"`
	Relationship    *d.CodeableConcept               `json:"relationship,omitempty"`
	Gender          *d.Code                          `json:"gender,omitempty"`
	BornPeriod      *d.Period                        `json:"bornPeriod,omitempty"`
	BornDate        *d.Date                          `json:"bornDate,omitempty"`
	BornString      *d.String                        `json:"bornString,omitempty"`
	AgeAge          *d.Age                           `json:"ageAge,omitempty"`
	AgeRange        *d.Range                         `json:"ageRange,omitempty"`
	AgeString       *d.String                        `json:"ageString,omitempty"`
	EstimatedAge    *d.Boolean                       `json:"estimatedAge,omitempty"`
	DeceasedBoolean *d.Boolean                       `json:"deceasedBoolean,omitempty"`
	DeceasedAge     *d.Age                           `json:"deceasedAge,omitempty"`
	DeceasedRange   *d.Range                         `json:"deceasedRange,omitempty"`
	DeceasedDate    *d.Date                          `json:"deceasedDate,omitempty"`
	DeceasedString  *d.String                        `json:"deceasedString,omitempty"`
	ReasonCode      []d.CodeableConcept              `json:"reasonCode,omitempty"`
	ReasonReference []d.Reference                    `json:"reasonReference,omitempty"`
	Note            []d.Annotation                   `json:"note,omitempty"`
	Condition       []d.FamilyMemberHistoryCondition `json:"condition,omitempty"`
}

// Validate returns a check against schema
func (f *FamilyMemberHistory) Validate() (bool, []error) {
	return schema.ValidateResource(*f, "3")
}
