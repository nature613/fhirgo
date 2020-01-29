package resources

import (
    d "github.com/monarko/fhirgo/STU3/datatypes"
    "github.com/monarko/fhirgo/schema"
)

// PlanDefinition resource
type PlanDefinition struct {
    Domain
    URL             *d.URI                   `json:"url,omitempty"`
    Identifier      []d.Identifier           `json:"identifier,omitempty"`
    Version         *d.String                `json:"version,omitempty"`
    Name            *d.String                `json:"name,omitempty"`
    Title           *d.String                `json:"title,omitempty"`
    Type            *d.CodeableConcept       `json:"type,omitempty"`
    Status          *d.Code                  `json:"status,omitempty"`
    Experimental    *d.Boolean               `json:"experimental,omitempty"`
    Date            *d.DateTime              `json:"date,omitempty"`
    Publisher       *d.String                `json:"publisher,omitempty"`
    Description     *d.Markdown              `json:"description,omitempty"`
    Purpose         *d.Markdown              `json:"purpose,omitempty"`
    Usage           *d.String                `json:"usage,omitempty"`
    ApprovalDate    *d.Date                  `json:"approvalDate,omitempty"`
    LastReviewDate  *d.Date                  `json:"lastReviewDate,omitempty"`
    EffectivePeriod *d.Period                `json:"effectivePeriod,omitempty"`
    UseContext      []d.UsageContext         `json:"useContext,omitempty"`
    Jurisdiction    []d.CodeableConcept      `json:"jurisdiction,omitempty"`
    Topic           []d.CodeableConcept      `json:"topic,omitempty"`
    Contributor     []d.Contributor          `json:"contributor,omitempty"`
    Contact         []d.ContactDetail        `json:"contact,omitempty"`
    Copyright       *d.Markdown              `json:"copyright,omitempty"`
    RelatedArtifact []d.RelatedArtifact      `json:"relatedArtifact,omitempty"`
    Library         []d.Reference            `json:"library,omitempty"`
    Goal            []d.PlanDefinitionGoal   `json:"goal,omitempty"`
    Action          []d.PlanDefinitionAction `json:"action,omitempty"`
}

// Validate returns a check against schema
func (p *PlanDefinition) Validate() (bool, []error) {
    return schema.ValidateResource(*p, "3")
}
