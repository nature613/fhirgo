package resources

import (
	d "github.com/monarko/fhirgo/STU3/datatypes"
	"github.com/monarko/fhirgo/schema"
)

// Task Resource
type Task struct {
	Domain
	Identifier          []d.Identifier      `json:"identifier,omitempty"`
	DefinitionURI       *d.URI              `json:"definitionUri,omitempty"`
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
	Requester           *TaskRequester      `json:"requester,omitempty"`
	PerformerType       []d.CodeableConcept `json:"performerType,omitempty"`
	Owner               *d.Reference        `json:"owner,omitempty"`
	Reason              *d.CodeableConcept  `json:"reason,omitempty"`
	Note                []d.Annotation      `json:"note,omitempty"`
	RelevantHistory     []d.Reference       `json:"relevantHistory,omitempty"`
	Restriction         *TaskRestriction    `json:"restriction,omitempty"`
	Input               []TaskInputOutput   `json:"input,omitempty"`
	Output              []TaskInputOutput   `json:"output,omitempty"`
}

// Validate returns a check against schema
func (t *Task) Validate() (bool, []error) {
	return schema.ValidateResource(*t, "3")
}

// TaskRequester sub-resource
type TaskRequester struct {
	Agent      *d.Reference `json:"agent,omitempty"`
	OnBehalfOf *d.Reference `json:"onBehalfOf,omitempty"`
}

// TaskRestriction sub-resource
type TaskRestriction struct {
	Repetitions *d.PositiveInt `json:"repetitions,omitempty"`
	Period      *d.Period      `json:"period,omitempty"`
	Recipient   []d.Reference  `json:"recipient,omitempty"`
}

// TaskInputOutput sub-resource
type TaskInputOutput struct {
	Type                 *d.CodeableConcept `json:"type,omitempty"`
	ValueInteger         *d.Integer         `json:"valueInteger,omitempty"`
	ValueDecimal         *d.Decimal         `json:"valueDecimal,omitempty"`
	ValueDateTime        *d.DateTime        `json:"valueDateTime,omitempty"`
	ValueDate            *d.Date            `json:"valueDate,omitempty"`
	ValueInstant         *d.Instant         `json:"valueInstant,omitempty"`
	ValueString          *d.String          `json:"valueString,omitempty"`
	ValueURI             *d.URI             `json:"valueUri,omitempty"`
	ValueBoolean         *d.Boolean         `json:"valueBoolean,omitempty"`
	ValueCode            *d.Code            `json:"valueCode,omitempty"`
	ValueBase64Binary    *d.Base64Binary    `json:"valueBase64Binary,omitempty"`
	ValueCoding          *d.Coding          `json:"valueCoding,omitempty"`
	ValueCodeableConcept *d.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueAttachment      *d.Attachment      `json:"valueAttachment,omitempty"`
	ValueIdentifier      *d.Identifier      `json:"valueIdentifier,omitempty"`
	ValueQuantity        *d.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *d.Range           `json:"valueRange,omitempty"`
	ValuePeriod          *d.Period          `json:"valuePeriod,omitempty"`
	ValueRatio           *d.Ratio           `json:"valueRatio,omitempty"`
	ValueHumanName       *d.HumanName       `json:"valueHumanName,omitempty"`
	ValueAddress         *d.Address         `json:"valueAddress,omitempty"`
	ValueContactPoint    *d.ContactPoint    `json:"valueContactPoint,omitempty"`
	ValueSchedule        interface{}        `json:"valueSchedule,omitempty"`
	ValueReference       *d.Reference       `json:"valueReference,omitempty"`
}
