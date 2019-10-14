package datatypes

// ContactDetail FHIR Type
type ContactDetail struct {
	Element
	Name    *String        `json:"name,omitempty"`
	Telecom []ContactPoint `json:"telecom,omitempty"`
}

// Contributor FHIR Type
type Contributor struct {
	Element
	ResourceType string          `json:"resourceType"`
	Type         *Code           `json:"type,omitempty"`
	Name         *String         `json:"name,omitempty"`
	Contact      []ContactDetail `json:"contact,omitempty"`
}

// DataRequirementCodeFilter FHIR Sub-Type
type DataRequirementCodeFilter struct {
	Path                 *String           `json:"path,omitempty"`
	ValueSetString       *String           `json:"valueSetString,omitempty"`
	ValueSetReference    *Reference        `json:"valueSetReference,omitempty"`
	ValueCode            []Code            `json:"valueCode,omitempty"`
	ValueCoding          []Coding          `json:"valueCoding,omitempty"`
	ValueCodeableConcept []CodeableConcept `json:"valueCodeableConcept,omitempty"`
}

// DataRequirementDateFilter FHIR Sub-Type
type DataRequirementDateFilter struct {
	Path          *String   `json:"path,omitempty"`
	ValueDateTime *DateTime `json:"valueDateTime,omitempty"`
	ValuePeriod   *Period   `json:"valuePeriod,omitempty"`
	ValueDuration *Duration `json:"valueDuration,omitempty"`
}

// DataRequirement FHIR Type
type DataRequirement struct {
	Element
	Type        *Code                       `json:"type,omitempty"`
	Profile     []URI                       `json:"profile,omitempty"`
	MustSupport []String                    `json:"mustSupport,omitempty"`
	CodeFilter  []DataRequirementCodeFilter `json:"codeFilter,omitempty"`
	DateFilter  []DataRequirementDateFilter `json:"dateFilter,omitempty"`
}

// RelatedArtifact FHIR Type
type RelatedArtifact struct {
	Element
	ResourceType string      `json:"resourceType"`
	Type         *Code       `json:"type,omitempty"`
	Display      *String     `json:"display,omitempty"`
	Citation     *String     `json:"citation,omitempty"`
	URL          *URI        `json:"url,omitempty"`
	Document     *Attachment `json:"document,omitempty"`
	Resource     *Reference  `json:"resource,omitempty"`
}

// UsageContext FHIR Type
type UsageContext struct {
	Element
	ResourceType         string           `json:"resourceType"`
	Code                 *Coding          `json:"coding,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range           `json:"valueRange,omitempty"`
}

// ParameterDefinition FHIR Type
type ParameterDefinition struct {
	Element
	Name          *Code    `json:"name,omitempty"`
	Use           *Code    `json:"use,omitempty"`
	Min           *Integer `json:"min,omitempty"`
	Max           *String  `json:"max,omitempty"`
	Documentation *String  `json:"documentation,omitempty"`
	Type          *Code    `json:"type,omitempty"`
	Profile       *URI     `json:"profile,omitempty"`
}

// Expression FHIR Type
type Expression struct {
	Element
	Description *String `json:"description,omitempty"`
	Name        *ID     `json:"name,omitempty"`
	Language    *Code   `json:"language,omitempty"`
	Expression  *String `json:"expression,omitempty"`
	Reference   *URI    `json:"reference,omitempty"`
}

// TriggerDefinition FHIR Type
type TriggerDefinition struct {
	Element
	Type                 *Code             `json:"type,omitempty"`
	EventName            *String           `json:"eventName,omitempty"`
	EventTimingTiming    *Timing           `json:"eventTimingTiming,omitempty"`
	EventTimingReference *Reference        `json:"eventTimingReference,omitempty"`
	EventTimingDate      *Date             `json:"eventTimingDate,omitempty"`
	EventTimingDateTime  *DateTime         `json:"eventTimingDateTime,omitempty"`
	EventData            []DataRequirement `json:"eventData,omitempty"`
}
