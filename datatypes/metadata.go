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
	Type    *Code           `json:"type,omitempty"`
	Name    *String         `json:"name,omitempty"`
	Contact []ContactDetail `json:"contact,omitempty"`
}

// DataRequirementCodeFilter FHIR Sub-Type
type DataRequirementCodeFilter struct {
	Path        *String    `json:"path,omitempty"`
	SearchParam *String    `json:"searchParam,omitempty"`
	ValueSet    *Canonical `json:"valueSet,omitempty"`
	Code        []Coding   `json:"code,omitempty"`
}

// DataRequirementDateFilter FHIR Sub-Type
type DataRequirementDateFilter struct {
	Path          *String   `json:"path,omitempty"`
	SearchParam   *String   `json:"searchParam,omitempty"`
	ValueDateTime *DateTime `json:"valueDateTime,omitempty"`
	ValuePeriod   *Period   `json:"valuePeriod,omitempty"`
	ValueDuration *Duration `json:"valueDuration,omitempty"`
}

// DataRequirementSort FHIR Sub-Type
type DataRequirementSort struct {
	Path      *String `json:"path,omitempty"`
	Direction *Code   `json:"direction,omitempty"`
}

// DataRequirement FHIR Type
type DataRequirement struct {
	Element
	Type                   *Code                       `json:"type,omitempty"`
	Profile                []Canonical                 `json:"profile,omitempty"`
	SubjectCodeableConcept *CodeableConcept            `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference                  `json:"subjectReference,omitempty"`
	MustSupport            []String                    `json:"mustSupport,omitempty"`
	CodeFilter             []DataRequirementCodeFilter `json:"codeFilter,omitempty"`
	DateFilter             []DataRequirementDateFilter `json:"dateFilter,omitempty"`
	Limit                  *PositiveInt                `json:"limit,omitempty"`
	Sort                   []DataRequirementSort       `json:"sort,omitempty"`
}

// RelatedArtifact FHIR Type
type RelatedArtifact struct {
	Element
	Type     *Code       `json:"type,omitempty"`
	Label    *String     `json:"label,omitempty"`
	Display  *String     `json:"display,omitempty"`
	Citation *Markdown   `json:"citation,omitempty"`
	URL      *URL        `json:"url,omitempty"`
	Document *Attachment `json:"document,omitempty"`
	Resource *Canonical  `json:"resource,omitempty"`
}

// UsageContext FHIR Type
type UsageContext struct {
	Element
	Code                 *Coding          `json:"coding,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range           `json:"valueRange,omitempty"`
	ValueReference       *Reference       `json:"valueReference,omitempty"`
}

// ParameterDefinition FHIR Type
type ParameterDefinition struct {
	Element
	Name          *Code      `json:"name,omitempty"`
	Use           *Code      `json:"use,omitempty"`
	Min           *Integer   `json:"min,omitempty"`
	Max           *String    `json:"max,omitempty"`
	Documentation *String    `json:"documentation,omitempty"`
	Type          *Code      `json:"type,omitempty"`
	Profile       *Canonical `json:"profile,omitempty"`
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
	Type            *Code             `json:"type,omitempty"`
	Name            *String           `json:"name,omitempty"`
	TimingTiming    *Timing           `json:"timingTiming,omitempty"`
	TimingReference *Reference        `json:"timingReference,omitempty"`
	TimingDate      *Date             `json:"timingDate,omitempty"`
	TimingDateTime  *DateTime         `json:"timingDateTime,omitempty"`
	Data            []DataRequirement `json:"data,omitempty"`
	Condition       *Expression       `json:"condition,omitempty"`
}
