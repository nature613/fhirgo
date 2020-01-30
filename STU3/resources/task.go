package resources

import (
    d "github.com/monarko/fhirgo/STU3/datatypes"
    "github.com/monarko/fhirgo/schema"
)

// Task Resource
type Task struct {
    Domain
    Identifier          []d.Identifier      `json:"identifier,omitempty"`
    DefinitionUri       *d.URI              `json:"definitionUri,omitempty"`
    DefinitionReference *d.Reference        `json:"definitionReference,omitempty"`
    BasedOn             []d.Reference       `json:"basedOn,omitempty"`
    GroupIdentifier     *d.Identifier       `json:"groupIdentifier,omitempty"`
    PartOf              []d.Reference       `json:"partOf,omitempty"`
    Status              *d.Code             `json:"status,omitempty"`
    StatusReason        *d.CodeableConcept  `json:"statusReason,omitempty"`
    BusinessStatus      *d.CodeableConcept  `json:"businessStatus,omitempty"`
    Intent              *d.Code             `json:"intent,omitempty"`
    Priority            *d.Code             `json:"priority,omitempty"`
    Code                *d.CodeableConcept  `json:"code,omitempty"`
    Description         *d.String           `json:"description,omitempty"`
    Focus               *d.Reference        `json:"focus,omitempty"`
    For                 *d.Reference        `json:"for,omitempty"`
    Context             *d.Reference        `json:"context,omitempty"`
    ExecutionPeriod     *d.Period           `json:"executionPeriod,omitempty"`
    AuthoredOn          *d.DateTime         `json:"authoredOn,omitempty"`
    LastModified        *d.DateTime         `json:"lastModified,omitempty"`
    Requester           *d.TaskRequester    `json:"requester,omitempty"`
    PerformerType       []d.CodeableConcept `json:"performerType,omitempty"`
    Owner               *d.Reference        `json:"owner,omitempty"`
    Reason              *d.CodeableConcept  `json:"reason,omitempty"`
    Note                []d.Annotation      `json:"note,omitempty"`
    RelevantHistory     []d.Reference       `json:"relevantHistory,omitempty"`
    Restriction         *d.TaskRestriction  `json:"restriction,omitempty"`
    Input               []d.TaskInput       `json:"input,omitempty"`
    Output              []d.TaskOutput      `json:"output,omitempty"`
}

// Validate returns a check against schema
func (t *Task) Validate() (bool, []error) {
    return schema.ValidateResource(*t, "3")
}
