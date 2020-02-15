package resources

import (
    d "github.com/monarko/fhirgo/STU3/datatypes"
    "github.com/monarko/fhirgo/schema"
)

// CarePlan resource
type CarePlan struct {
    Domain
    Identifier     []d.Identifier      `json:"identifier,omitempty"`
    Definition     []d.Reference       `json:"definition,omitempty"`
    BasedOn        []d.Reference       `json:"basedOn,omitempty"`
    Replaces       []d.Reference       `json:"replaces,omitempty"`
    PartOf         []d.Reference       `json:"partOf,omitempty"`
    Status         *d.Code             `json:"status,omitempty"`
    Intent         *d.Code             `json:"intent,omitempty"`
    Category       []d.CodeableConcept `json:"category,omitempty"`
    Title          *d.String           `json:"title,omitempty"`
    Description    *d.String           `json:"description,omitempty"`
    Subject        *d.Reference        `json:"subject,omitempty"`
    Context        *d.Reference        `json:"context,omitempty"`
    Period         *d.Period           `json:"period,omitempty"`
    Author         []d.Reference       `json:"author,omitempty"`
    CareTeam       []d.Reference       `json:"careTeam,omitempty"`
    Addresses      []d.Reference       `json:"addresses,omitempty"`
    SupportingInfo []d.Reference       `json:"supportingInfo,omitempty"`
    Goal           []d.Reference       `json:"goal,omitempty"`
    Activity       []CarePlanActivity  `json:"activity,omitempty"`
    Note           []d.Annotation      `json:"note,omitempty"`
}

// Validate returns a check against schema
func (c *CarePlan) Validate() (bool, []error) {
    return schema.ValidateResource(*c, "3")
}

// CarePlanActivity subResource
type CarePlanActivity struct {
    OutcomeCodeableConcept []d.CodeableConcept     `json:"outcomeCodeableConcept,omitempty"`
    OutcomeReference       []d.Reference           `json:"outcomeReference,omitempty"`
    Progress               []d.Annotation          `json:"progress,omitempty"`
    Reference              *d.Reference            `json:"reference,omitempty"`
    Detail                 *CarePlanActivityDetail `json:"detail,omitempty"`
}

// CarePlanActivityDetail subResource
type CarePlanActivityDetail struct {
    Category               *d.CodeableConcept  `json:"category,omitempty"`
    Definition             *d.Reference        `json:"definition,omitempty"`
    Code                   *d.CodeableConcept  `json:"code,omitempty"`
    ReasonCode             []d.CodeableConcept `json:"reasonCode,omitempty"`
    ReasonReference        []d.Reference       `json:"reasonReference,omitempty"`
    Goal                   []d.Reference       `json:"goal,omitempty"`
    Status                 *d.Code             `json:"status,omitempty"`
    StatusReason           *d.String           `json:"statusReason,omitempty"`
    Prohibited             *d.Boolean          `json:"prohibited,omitempty"`
    ScheduledTiming        *d.Timing           `json:"scheduledTiming,omitempty"`
    ScheduledPeriod        *d.Period           `json:"scheduledPeriod,omitempty"`
    ScheduledString        *d.String           `json:"scheduledString,omitempty"`
    Location               *d.Reference        `json:"location,omitempty"`
    Performer              []d.Reference       `json:"performer,omitempty"`
    ProductCodeableConcept *d.CodeableConcept  `json:"productCodeableConcept,omitempty"`
    ProductReference       *d.Reference        `json:"productReference,omitempty"`
    DailyAmount            *d.SimpleQuantity   `json:"dailyAmount,omitempty"`
    Quantity               *d.SimpleQuantity   `json:"quantity,omitempty"`
    Description            *d.String           `json:"description,omitempty"`
}
