package resources

import (
    d "github.com/monarko/fhirgo/STU3/datatypes"
    "github.com/monarko/fhirgo/schema"
)

// PlanDefinition resource
type PlanDefinition struct {
    Domain
    URL             *d.URI                 `json:"url,omitempty"`
    Identifier      []d.Identifier         `json:"identifier,omitempty"`
    Version         *d.String              `json:"version,omitempty"`
    Name            *d.String              `json:"name,omitempty"`
    Title           *d.String              `json:"title,omitempty"`
    Type            *d.CodeableConcept     `json:"type,omitempty"`
    Status          *d.Code                `json:"status,omitempty"`
    Experimental    *d.Boolean             `json:"experimental,omitempty"`
    Date            *d.DateTime            `json:"date,omitempty"`
    Publisher       *d.String              `json:"publisher,omitempty"`
    Description     *d.Markdown            `json:"description,omitempty"`
    Purpose         *d.Markdown            `json:"purpose,omitempty"`
    Usage           *d.String              `json:"usage,omitempty"`
    ApprovalDate    *d.Date                `json:"approvalDate,omitempty"`
    LastReviewDate  *d.Date                `json:"lastReviewDate,omitempty"`
    EffectivePeriod *d.Period              `json:"effectivePeriod,omitempty"`
    UseContext      []d.UsageContext       `json:"useContext,omitempty"`
    Jurisdiction    []d.CodeableConcept    `json:"jurisdiction,omitempty"`
    Topic           []d.CodeableConcept    `json:"topic,omitempty"`
    Contributor     []d.Contributor        `json:"contributor,omitempty"`
    Contact         []d.ContactDetail      `json:"contact,omitempty"`
    Copyright       *d.Markdown            `json:"copyright,omitempty"`
    RelatedArtifact []d.RelatedArtifact    `json:"relatedArtifact,omitempty"`
    Library         []d.Reference          `json:"library,omitempty"`
    Goal            []PlanDefinitionGoal   `json:"goal,omitempty"`
    Action          []PlanDefinitionAction `json:"action,omitempty"`
}

// Validate returns a check against schema
func (p *PlanDefinition) Validate() (bool, []error) {
    return schema.ValidateResource(*p, "3")
}

// PlanDefinitionGoal sub-resource
type PlanDefinitionGoal struct {
    Category      *d.CodeableConcept         `json:"category,omitempty"`
    Description   *d.CodeableConcept         `json:"description,omitempty"`
    Priority      *d.CodeableConcept         `json:"priority,omitempty"`
    Start         *d.CodeableConcept         `json:"start,omitempty"`
    Addresses     []d.CodeableConcept        `json:"addresses,omitempty"`
    Documentation []d.RelatedArtifact        `json:"documentation,omitempty"`
    Target        []PlanDefinitionGoalTarget `json:"target,omitempty"`
}

// PlanDefinitionGoalTarget sub-resource
type PlanDefinitionGoalTarget struct {
    Measure               *d.CodeableConcept `json:"measure,omitempty"`
    DetailQuantity        *d.Quantity        `json:"detailQuantity,omitempty"`
    DetailRange           *d.Range           `json:"detailRange,omitempty"`
    DetailCodeableConcept *d.CodeableConcept `json:"detailCodeableConcept,omitempty"`
    Due                   *d.Duration        `json:"due,omitempty"`
}

// PlanDefinitionAction sub-resource
type PlanDefinitionAction struct {
    Label               *d.String                           `json:"label,omitempty"`
    Title               *d.String                           `json:"title,omitempty"`
    Description         *d.String                           `json:"description,omitempty"`
    TextEquivalent      *d.String                           `json:"textEquivalent,omitempty"`
    Code                []d.CodeableConcept                 `json:"code,omitempty"`
    Reason              []d.CodeableConcept                 `json:"reason,omitempty"`
    Documentation       []d.RelatedArtifact                 `json:"documentation,omitempty"`
    GoalId              []d.ID                              `json:"goalId,omitempty"`
    TriggerDefinition   []d.TriggerDefinition               `json:"triggerDefinition,omitempty"`
    Condition           []PlanDefinitionActionCondition     `json:"condition,omitempty"`
    Input               []d.DataRequirement                 `json:"input,omitempty"`
    Output              []d.DataRequirement                 `json:"output,omitempty"`
    RelatedAction       []PlanDefinitionActionRelatedAction `json:"relatedAction,omitempty"`
    TimingDateTime      *d.DateTime                         `json:"timingDateTime,omitempty"`
    TimingPeriod        *d.Period                           `json:"timingPeriod,omitempty"`
    TimingDuration      *d.Duration                         `json:"timingDuration,omitempty"`
    TimingRange         *d.Range                            `json:"timingRange,omitempty"`
    TimingTiming        *d.Timing                           `json:"timingTiming,omitempty"`
    Participant         []ActivityDefinitionParticipant     `json:"participant,omitempty"`
    Type                *d.Coding                           `json:"type,omitempty"`
    GroupingBehavior    *d.Code                             `json:"groupingBehavior,omitempty"`
    SelectionBehavior   *d.Code                             `json:"selectionBehavior,omitempty"`
    RequiredBehavior    *d.Code                             `json:"requiredBehavior,omitempty"`
    PrecheckBehavior    *d.Code                             `json:"precheckBehavior,omitempty"`
    CardinalityBehavior *d.Code                             `json:"cardinalityBehavior,omitempty"`
    Definition          *d.Reference                        `json:"definition,omitempty"`
    Transform           *d.Reference                        `json:"transform,omitempty"`
    DynamicValue        []ActivityDefinitionDynamicValue    `json:"dynamicValue,omitempty"`
    Action              interface{}                         `json:"action,omitempty"`
}

// PlanDefinitionActionCondition sub-resource
type PlanDefinitionActionCondition struct {
    Kind        *d.Code   `json:"kind,omitempty"`
    Description *d.String `json:"description,omitempty"`
    Language    *d.String `json:"language,omitempty"`
    Expression  *d.String `json:"expression,omitempty"`
}

// PlanDefinitionActionRelatedAction sub-resource
type PlanDefinitionActionRelatedAction struct {
    ActionId       *d.ID       `json:"actionId,omitempty"`
    Relationship   *d.Code     `json:"relationship,omitempty"`
    OffsetDuration *d.Duration `json:"offsetDuration,omitempty"`
    OffsetRange    *d.Range    `json:"offsetRange,omitempty"`
}
