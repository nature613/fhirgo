package resources

import (
    d "github.com/monarko/fhirgo/STU3/datatypes"
    "github.com/monarko/fhirgo/schema"
)

// ActivityDefinition resource
type ActivityDefinition struct {
    Domain
    URL                    *d.URI                           `json:"url,omitempty"`
    Identifier             []d.Identifier                   `json:"identifier,omitempty"`
    Version                *d.String                        `json:"version,omitempty"`
    Name                   *d.String                        `json:"name,omitempty"`
    Title                  *d.String                        `json:"title,omitempty"`
    Status                 *d.Code                          `json:"status,omitempty"`
    Experimental           *d.Boolean                       `json:"experimental,omitempty"`
    Date                   *d.DateTime                      `json:"date,omitempty"`
    Publisher              *d.String                        `json:"publisher,omitempty"`
    Description            *d.Markdown                      `json:"description,omitempty"`
    Purpose                *d.Markdown                      `json:"purpose,omitempty"`
    Usage                  *d.String                        `json:"usage,omitempty"`
    ApprovalDate           *d.Date                          `json:"approvalDate,omitempty"`
    LastReviewDate         *d.Date                          `json:"lastReviewDate,omitempty"`
    EffectivePeriod        *d.Period                        `json:"effectivePeriod,omitempty"`
    UseContext             []d.UsageContext                 `json:"useContext,omitempty"`
    Jurisdiction           []d.CodeableConcept              `json:"jurisdiction,omitempty"`
    Topic                  []d.CodeableConcept              `json:"topic,omitempty"`
    Contributor            []d.Contributor                  `json:"contributor,omitempty"`
    Contact                []d.ContactDetail                `json:"contact,omitempty"`
    Copyright              *d.Markdown                      `json:"copyright,omitempty"`
    RelatedArtifact        []d.RelatedArtifact              `json:"relatedArtifact,omitempty"`
    Library                []d.Reference                    `json:"library,omitempty"`
    Kind                   *d.Code                          `json:"kind,omitempty"`
    Code                   *d.CodeableConcept               `json:"code,omitempty"`
    TimingTiming           *d.Timing                        `json:"timingTiming,omitempty"`
    TimingDateTime         *d.DateTime                      `json:"timingDateTime,omitempty"`
    TimingPeriod           *d.Period                        `json:"timingPeriod,omitempty"`
    TimingRange            *d.Range                         `json:"timingRange,omitempty"`
    Location               *d.Reference                     `json:"location,omitempty"`
    Participant            []ActivityDefinitionParticipant  `json:"participant,omitempty"`
    ProductReference       *d.Reference                     `json:"productReference,omitempty"`
    ProductCodeableConcept *d.CodeableConcept               `json:"productCodeableConcept,omitempty"`
    Quantity               *d.Quantity                      `json:"quantity,omitempty"`
    Dosage                 []d.Dosage                       `json:"dosage,omitempty"`
    BodySite               []d.CodeableConcept              `json:"bodySite,omitempty"`
    Transform              *d.Reference                     `json:"transform,omitempty"`
    DynamicValue           []ActivityDefinitionDynamicValue `json:"dynamicValue,omitempty"`
}

// Validate returns a check against schema
func (a *ActivityDefinition) Validate() (bool, []error) {
    return schema.ValidateResource(*a, "3")
}

// ActivityDefinitionParticipant sub-resource
type ActivityDefinitionParticipant struct {
    Type *d.Code            `json:"type,omitempty"`
    Role *d.CodeableConcept `json:"role,omitempty"`
}

// ActivityDefinitionDynamicValue sub-resource
type ActivityDefinitionDynamicValue struct {
    Description *d.String `json:"description,omitempty"`
    Path        *d.String `json:"path,omitempty"`
    Language    *d.String `json:"language,omitempty"`
    Expression  *d.String `json:"expression,omitempty"`
}
