package datatypes

// Element FHIR Type
type Element struct {
	ID        *String     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
}

// Reference FHIR Type
type Reference struct {
	Element
	Reference  *String     `json:"reference,omitempty"`
	Identifier *Identifier `json:"identifier,omitempty"`
	Display    *String     `json:"display,omitempty"`
}

// Meta FHIR Type
type Meta struct {
	Element
	VersionID   *ID      `json:"versionId,omitempty"`
	LastUpdated *Instant `json:"lastUpdated,omitempty"`
	Profile     []URI    `json:"profile,omitempty"`
	Security    []Coding `json:"security,omitempty"`
	Tag         []Coding `json:"tag,omitempty"`
}

// Dosage FHIR Type
type Dosage struct {
	Element
	Sequence                 *Integer          `json:"sequence,omitempty"`
	Text                     *String           `json:"text,omitempty"`
	AdditionalInstruction    []CodeableConcept `json:"additionalInstruction,omitempty"`
	PatientInstruction       *String           `json:"patientInstruction,omitempty"`
	Timing                   *Timing           `json:"timing,omitempty"`
	AsNeededBoolean          *Boolean          `json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept  *CodeableConcept  `json:"asNeededCodeableConcept,omitempty"`
	Site                     *CodeableConcept  `json:"site,omitempty"`
	Route                    *CodeableConcept  `json:"route,omitempty"`
	Method                   *CodeableConcept  `json:"method,omitempty"`
	DoseRange                *Range            `json:"doseRange,omitempty"`
	DoseQuantity             *SimpleQuantity   `json:"doseQuantity,omitempty"`
	MaxDosePerPeriod         *Ratio            `json:"maxDosePerPeriod,omitempty"`
	MaxDosePerAdministration *SimpleQuantity   `json:"maxDosePerAdministration,omitempty"`
	MaxDosePerLifetime       *SimpleQuantity   `json:"maxDosePerLifetime,omitempty"`
	RateRatio                *Ratio            `json:"rateRatio,omitempty"`
	RateRange                *Range            `json:"rateRange,omitempty"`
	RateQuantity             *SimpleQuantity   `json:"rateQuantity,omitempty"`
}

// Narrative for domain resource
type Narrative struct {
	Element
	Status *Code   `json:"status,omitempty"`
	Div    *String `json:"div,omitempty"`
}

// Extension FHIR Type
type Extension struct {
	URL                  *URI             `json:"url,omitempty"`
	ValueInteger         *Integer         `json:"valueInteger,omitempty"`
	ValueDecimal         *Decimal         `json:"valueDecimal,omitempty"`
	ValueDateTime        *DateTime        `json:"valueDateTime,omitempty"`
	ValueDate            *Date            `json:"valueDate,omitempty"`
	ValueInstant         *Instant         `json:"valueInstant,omitempty"`
	ValueString          *String          `json:"valueString,omitempty"`
	ValueURI             *URI             `json:"valueUri,omitempty"`
	ValueBoolean         *Boolean         `json:"valueBoolean,omitempty"`
	ValueCode            *Code            `json:"valueCode,omitempty"`
	ValueBase64Binary    *Base64Binary    `json:"valueBase64Binary,omitempty"`
	ValueCoding          *Coding          `json:"valueCoding,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueAttachment      *Attachment      `json:"valueAttachment,omitempty"`
	ValueIdentifier      *Identifier      `json:"valueIdentifier,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range           `json:"valueRange,omitempty"`
	ValuePeriod          *Period          `json:"valuePeriod,omitempty"`
	ValueRatio           *Ratio           `json:"valueRatio,omitempty"`
	ValueHumanName       *HumanName       `json:"valueHumanName,omitempty"`
	ValueAddress         *Address         `json:"valueAddress,omitempty"`
	ValueContactPoint    *ContactPoint    `json:"valueContactPoint,omitempty"`
	ValueSchedule        *Schedule        `json:"valueSchedule,omitempty"`
	ValueReference       *Reference       `json:"valueReference,omitempty"`
}

// Schedule FHIR type
type Schedule struct {
	Identifier      *Identifier       `json:"identifier,omitempty"`
	Active          *Boolean          `json:"active,omitempty"`
	ServiceCategory *CodeableConcept  `json:"serviceCategory,omitempty"`
	ServiceType     []CodeableConcept `json:"serviceType,omitempty"`
	Specialty       *CodeableConcept  `json:"specialty,omitempty"`
	Actor           *Reference        `json:"actor,omitempty"`
	PlanningHorizon *Period           `json:"planningHorizon,omitempty"`
	Comment         *String           `json:"comment,omitempty"`
}

// // ElementDefinitionSlicingDiscriminator FHIR Sub-Type
// type ElementDefinitionSlicingDiscriminator struct {
// 	Type *Code   `json:"type,omitempty"`
// 	Path *String `json:"path,omitempty"`
// }

// // ElementDefinitionSlicing FHIR Sub-Type
// type ElementDefinitionSlicing struct {
// 	Discriminator []ElementDefinitionSlicingDiscriminator `json:"discriminator,omitempty"`
// 	Description   *String                                 `json:"description,omitempty"`
// 	Ordered       *Boolean                                `json:"ordered,omitempty"`
// 	Rules         *Code                                   `json:"rules,omitempty"`
// }

// // ElementDefinitionBase FHIR Sub-Type
// type ElementDefinitionBase struct {
// 	Path *String      `json:"path,omitempty"`
// 	Min  *UnsignedInt `json:"min,omitempty"`
// 	Max  *String      `json:"max,omitempty"`
// }

// // ElementDefinitionType FHIR Sub-Type
// type ElementDefinitionType struct {
// 	Code          *URI        `json:"code,omitempty"`
// 	Profile       []Canonical `json:"profile,omitempty"`
// 	TargetProfile []Canonical `json:"targetProfile,omitempty"`
// 	Aggregation   []Code      `json:"aggregation,omitempty"`
// 	Versioning    *Code       `json:"versioning,omitempty"`
// }

// // ElementDefinitionExample FHIR Sub-Type
// type ElementDefinitionExample struct {
// 	Label                    *String              `json:"label,omitempty"`
// 	ValueBase64Binary        *Base64Binary        `json:"valueBase64Binary,omitempty"`
// 	ValueBoolean             *Boolean             `json:"valueBoolean,omitempty"`
// 	ValueCanonical           *Canonical           `json:"valueCanonical,omitempty"`
// 	ValueCode                *Code                `json:"valueCode,omitempty"`
// 	ValueDate                *Date                `json:"valueDate,omitempty"`
// 	ValueDateTime            *DateTime            `json:"valueDateTime,omitempty"`
// 	ValueDecimal             *Decimal             `json:"valueDecimal,omitempty"`
// 	ValueID                  *ID                  `json:"valueId,omitempty"`
// 	ValueInstant             *Instant             `json:"valueInstant,omitempty"`
// 	ValueInteger             *Integer             `json:"valueInteger,omitempty"`
// 	ValueMarkdown            *Markdown            `json:"valueMarkdown,omitempty"`
// 	ValueOID                 *OID                 `json:"valueOid,omitempty"`
// 	ValuePositiveInt         *PositiveInt         `json:"valuePositiveInt,omitempty"`
// 	ValueString              *String              `json:"valueString,omitempty"`
// 	ValueTime                *Time                `json:"valueTime,omitempty"`
// 	ValueUnsignedInt         *UnsignedInt         `json:"valueUnsignedInt,omitempty"`
// 	ValueURI                 *URI                 `json:"valueUri,omitempty"`
// 	ValueURL                 *URL                 `json:"valueUrl,omitempty"`
// 	ValueUUID                *UUID                `json:"valueUuid,omitempty"`
// 	ValueAddress             *Address             `json:"valueAddress,omitempty"`
// 	ValueAge                 *Age                 `json:"valueAge,omitempty"`
// 	ValueAnnotation          *Annotation          `json:"valueAnnotation,omitempty"`
// 	ValueAttachment          *Attachment          `json:"valueAttachment,omitempty"`
// 	ValueCodeableConcept     *CodeableConcept     `json:"valueCodeableConcept,omitempty"`
// 	ValueCoding              *Coding              `json:"valueCoding,omitempty"`
// 	ValueContactPoint        *ContactPoint        `json:"valueContactPoint,omitempty"`
// 	ValueCount               *Count               `json:"valueCount,omitempty"`
// 	ValueDistance            *Distance            `json:"valueDistance,omitempty"`
// 	ValueDuration            *Duration            `json:"valueDuration,omitempty"`
// 	ValueHumanName           *HumanName           `json:"valueHumanName,omitempty"`
// 	ValueIdentifier          *Identifier          `json:"valueIdentifier,omitempty"`
// 	ValueMoney               *Money               `json:"valueMoney,omitempty"`
// 	ValuePeriod              *Period              `json:"valuePeriod,omitempty"`
// 	ValueQuantity            *Quantity            `json:"valueQuantity,omitempty"`
// 	ValueRange               *Range               `json:"valueRange,omitempty"`
// 	ValueRatio               *Ratio               `json:"valueRatio,omitempty"`
// 	ValueReference           *Reference           `json:"valueReference,omitempty"`
// 	ValueSampledData         *SampledData         `json:"valueSampledData,omitempty"`
// 	ValueSignature           *Signature           `json:"valueSignature,omitempty"`
// 	ValueTiming              *Timing              `json:"valueTiming,omitempty"`
// 	ValueContactDetail       *ContactDetail       `json:"valueContactDetail,omitempty"`
// 	ValueContributor         *Contributor         `json:"valueContributor,omitempty"`
// 	ValueDataRequirement     *DataRequirement     `json:"valueDataRequirement,omitempty"`
// 	ValueExpression          *Expression          `json:"valueExpression,omitempty"`
// 	ValueParameterDefinition *ParameterDefinition `json:"valueParameterDefinition,omitempty"`
// 	ValueRelatedArtifact     *RelatedArtifact     `json:"valueRelatedArtifact,omitempty"`
// 	ValueTriggerDefinition   *TriggerDefinition   `json:"valueTriggerDefinition,omitempty"`
// 	ValueUsageContext        *UsageContext        `json:"valueUsageContext,omitempty"`
// 	ValueDosage              *Dosage              `json:"valueDosage,omitempty"`
// }

// // ElementDefinitionConstraint FHIR Sub-Type
// type ElementDefinitionConstraint struct {
// 	Key          *ID        `json:"key,omitempty"`
// 	Requirements *String    `json:"requirements,omitempty"`
// 	Severity     *Code      `json:"severity,omitempty"`
// 	Human        *String    `json:"human,omitempty"`
// 	Expression   *String    `json:"expression,omitempty"`
// 	XPath        *String    `json:"xpath,omitempty"`
// 	Source       *Canonical `json:"source,omitempty"`
// }

// // ElementDefinitionBinding FHIR Sub-Type
// type ElementDefinitionBinding struct {
// 	Strength    *Code      `json:"strength,omitempty"`
// 	Description *String    `json:"description,omitempty"`
// 	ValueSet    *Canonical `json:"valueSet,omitempty"`
// }

// // ElementDefinitionMapping FHIR Sub-Type
// type ElementDefinitionMapping struct {
// 	Identity *ID     `json:"identity,omitempty"`
// 	Language *Code   `json:"language,omitempty"`
// 	Map      *String `json:"map,omitempty"`
// 	Comment  *String `json:"comment,omitempty"`
// }

// // ElementDefinition FHIR Type
// type ElementDefinition struct {
// 	Path                            *String                       `json:"path,omitempty"`
// 	Representation                  []Code                        `json:"representation,omitempty"`
// 	SliceName                       *String                       `json:"sliceName,omitempty"`
// 	SliceInConstraining             *Boolean                      `json:"sliceInConstraining,omitempty"`
// 	Label                           *String                       `json:"label,omitempty"`
// 	Code                            []Coding                      `json:"code,omitempty"`
// 	Slicing                         *ElementDefinitionSlicing     `json:"slicing,omitempty"`
// 	Short                           *String                       `json:"short,omitempty"`
// 	Definition                      *Markdown                     `json:"definition,omitempty"`
// 	Comment                         *Markdown                     `json:"comment,omitempty"`
// 	Requirements                    *Markdown                     `json:"requirements,omitempty"`
// 	Alias                           []String                      `json:"alias,omitempty"`
// 	Min                             *UnsignedInt                  `json:"min,omitempty"`
// 	Max                             *String                       `json:"max,omitempty"`
// 	Base                            *ElementDefinitionBase        `json:"base,omitempty"`
// 	ContentReference                *URI                          `json:"contentReference,omitempty"`
// 	Type                            []ElementDefinitionBase       `json:"type,omitempty"`
// 	DefaultValueBase64Binary        *Base64Binary                 `json:"defaultValueBase64Binary,omitempty"`
// 	DefaultValueBoolean             *Boolean                      `json:"defaultValueBoolean,omitempty"`
// 	DefaultValueCanonical           *Canonical                    `json:"defaultValueCanonical,omitempty"`
// 	DefaultValueCode                *Code                         `json:"defaultValueCode,omitempty"`
// 	DefaultValueDate                *Date                         `json:"defaultValueDate,omitempty"`
// 	DefaultValueDateTime            *DateTime                     `json:"defaultValueDateTime,omitempty"`
// 	DefaultValueDecimal             *Decimal                      `json:"defaultValueDecimal,omitempty"`
// 	DefaultValueID                  *ID                           `json:"defaultValueId,omitempty"`
// 	DefaultValueInstant             *Instant                      `json:"defaultValueInstant,omitempty"`
// 	DefaultValueInteger             *Integer                      `json:"defaultValueInteger,omitempty"`
// 	DefaultValueMarkdown            *Markdown                     `json:"defaultValueMarkdown,omitempty"`
// 	DefaultValueOID                 *OID                          `json:"defaultValueOid,omitempty"`
// 	DefaultValuePositiveInt         *PositiveInt                  `json:"defaultValuePositiveInt,omitempty"`
// 	DefaultValueString              *String                       `json:"defaultValueString,omitempty"`
// 	DefaultValueTime                *Time                         `json:"defaultValueTime,omitempty"`
// 	DefaultValueUnsignedInt         *UnsignedInt                  `json:"defaultValueUnsignedInt,omitempty"`
// 	DefaultValueURI                 *URI                          `json:"defaultValueUri,omitempty"`
// 	DefaultValueURL                 *URL                          `json:"defaultValueUrl,omitempty"`
// 	DefaultValueUUID                *UUID                         `json:"defaultValueUuid,omitempty"`
// 	DefaultValueAddress             *Address                      `json:"defaultValueAddress,omitempty"`
// 	DefaultValueAge                 *Age                          `json:"defaultValueAge,omitempty"`
// 	DefaultValueAnnotation          *Annotation                   `json:"defaultValueAnnotation,omitempty"`
// 	DefaultValueAttachment          *Attachment                   `json:"defaultValueAttachment,omitempty"`
// 	DefaultValueCodeableConcept     *CodeableConcept              `json:"defaultValueCodeableConcept,omitempty"`
// 	DefaultValueCoding              *Coding                       `json:"defaultValueCoding,omitempty"`
// 	DefaultValueContactPoint        *ContactPoint                 `json:"defaultValueContactPoint,omitempty"`
// 	DefaultValueCount               *Count                        `json:"defaultValueCount,omitempty"`
// 	DefaultValueDistance            *Distance                     `json:"defaultValueDistance,omitempty"`
// 	DefaultValueDuration            *Duration                     `json:"defaultValueDuration,omitempty"`
// 	DefaultValueHumanName           *HumanName                    `json:"defaultValueHumanName,omitempty"`
// 	DefaultValueIdentifier          *Identifier                   `json:"defaultValueIdentifier,omitempty"`
// 	DefaultValueMoney               *Money                        `json:"defaultValueMoney,omitempty"`
// 	DefaultValuePeriod              *Period                       `json:"defaultValuePeriod,omitempty"`
// 	DefaultValueQuantity            *Quantity                     `json:"defaultValueQuantity,omitempty"`
// 	DefaultValueRange               *Range                        `json:"defaultValueRange,omitempty"`
// 	DefaultValueRatio               *Ratio                        `json:"defaultValueRatio,omitempty"`
// 	DefaultValueReference           *Reference                    `json:"defaultValueReference,omitempty"`
// 	DefaultValueSampledData         *SampledData                  `json:"defaultValueSampledData,omitempty"`
// 	DefaultValueSignature           *Signature                    `json:"defaultValueSignature,omitempty"`
// 	DefaultValueTiming              *Timing                       `json:"defaultValueTiming,omitempty"`
// 	DefaultValueContactDetail       *ContactDetail                `json:"defaultValueContactDetail,omitempty"`
// 	DefaultValueContributor         *Contributor                  `json:"defaultValueContributor,omitempty"`
// 	DefaultValueDataRequirement     *DataRequirement              `json:"defaultValueDataRequirement,omitempty"`
// 	DefaultValueExpression          *Expression                   `json:"defaultValueExpression,omitempty"`
// 	DefaultValueParameterDefinition *ParameterDefinition          `json:"defaultValueParameterDefinition,omitempty"`
// 	DefaultValueRelatedArtifact     *RelatedArtifact              `json:"defaultValueRelatedArtifact,omitempty"`
// 	DefaultValueTriggerDefinition   *TriggerDefinition            `json:"defaultValueTriggerDefinition,omitempty"`
// 	DefaultValueUsageContext        *UsageContext                 `json:"defaultValueUsageContext,omitempty"`
// 	DefaultValueDosage              *Dosage                       `json:"defaultValueDosage,omitempty"`
// 	MeaningWhenMissing              *Markdown                     `json:"meaningWhenMissing,omitempty"`
// 	OrderMeaning                    *String                       `json:"orderMeaning,omitempty"`
// 	FixedBase64Binary               *Base64Binary                 `json:"fixedBase64Binary,omitempty"`
// 	FixedBoolean                    *Boolean                      `json:"fixedBoolean,omitempty"`
// 	FixedCanonical                  *Canonical                    `json:"fixedCanonical,omitempty"`
// 	FixedCode                       *Code                         `json:"fixedCode,omitempty"`
// 	FixedDate                       *Date                         `json:"fixedDate,omitempty"`
// 	FixedDateTime                   *DateTime                     `json:"fixedDateTime,omitempty"`
// 	FixedDecimal                    *Decimal                      `json:"fixedDecimal,omitempty"`
// 	FixedID                         *ID                           `json:"fixedId,omitempty"`
// 	FixedInstant                    *Instant                      `json:"fixedInstant,omitempty"`
// 	FixedInteger                    *Integer                      `json:"fixedInteger,omitempty"`
// 	FixedMarkdown                   *Markdown                     `json:"fixedMarkdown,omitempty"`
// 	FixedOID                        *OID                          `json:"fixedOid,omitempty"`
// 	FixedPositiveInt                *PositiveInt                  `json:"fixedPositiveInt,omitempty"`
// 	FixedString                     *String                       `json:"fixedString,omitempty"`
// 	FixedTime                       *Time                         `json:"fixedTime,omitempty"`
// 	FixedUnsignedInt                *UnsignedInt                  `json:"fixedUnsignedInt,omitempty"`
// 	FixedURI                        *URI                          `json:"fixedUri,omitempty"`
// 	FixedURL                        *URL                          `json:"fixedUrl,omitempty"`
// 	FixedUUID                       *UUID                         `json:"fixedUuid,omitempty"`
// 	FixedAddress                    *Address                      `json:"fixedAddress,omitempty"`
// 	FixedAge                        *Age                          `json:"fixedAge,omitempty"`
// 	FixedAnnotation                 *Annotation                   `json:"fixedAnnotation,omitempty"`
// 	FixedAttachment                 *Attachment                   `json:"fixedAttachment,omitempty"`
// 	FixedCodeableConcept            *CodeableConcept              `json:"fixedCodeableConcept,omitempty"`
// 	FixedCoding                     *Coding                       `json:"fixedCoding,omitempty"`
// 	FixedContactPoint               *ContactPoint                 `json:"fixedContactPoint,omitempty"`
// 	FixedCount                      *Count                        `json:"fixedCount,omitempty"`
// 	FixedDistance                   *Distance                     `json:"fixedDistance,omitempty"`
// 	FixedDuration                   *Duration                     `json:"fixedDuration,omitempty"`
// 	FixedHumanName                  *HumanName                    `json:"fixedHumanName,omitempty"`
// 	FixedIdentifier                 *Identifier                   `json:"fixedIdentifier,omitempty"`
// 	FixedMoney                      *Money                        `json:"fixedMoney,omitempty"`
// 	FixedPeriod                     *Period                       `json:"fixedPeriod,omitempty"`
// 	FixedQuantity                   *Quantity                     `json:"fixedQuantity,omitempty"`
// 	FixedRange                      *Range                        `json:"fixedRange,omitempty"`
// 	FixedRatio                      *Ratio                        `json:"fixedRatio,omitempty"`
// 	FixedReference                  *Reference                    `json:"fixedReference,omitempty"`
// 	FixedSampledData                *SampledData                  `json:"fixedSampledData,omitempty"`
// 	FixedSignature                  *Signature                    `json:"fixedSignature,omitempty"`
// 	FixedTiming                     *Timing                       `json:"fixedTiming,omitempty"`
// 	FixedContactDetail              *ContactDetail                `json:"fixedContactDetail,omitempty"`
// 	FixedContributor                *Contributor                  `json:"fixedContributor,omitempty"`
// 	FixedDataRequirement            *DataRequirement              `json:"fixedDataRequirement,omitempty"`
// 	FixedExpression                 *Expression                   `json:"fixedExpression,omitempty"`
// 	FixedParameterDefinition        *ParameterDefinition          `json:"fixedParameterDefinition,omitempty"`
// 	FixedRelatedArtifact            *RelatedArtifact              `json:"fixedRelatedArtifact,omitempty"`
// 	FixedTriggerDefinition          *TriggerDefinition            `json:"fixedTriggerDefinition,omitempty"`
// 	FixedUsageContext               *UsageContext                 `json:"fixedUsageContext,omitempty"`
// 	FixedDosage                     *Dosage                       `json:"fixedDosage,omitempty"`
// 	PatternBase64Binary             *Base64Binary                 `json:"patternBase64Binary,omitempty"`
// 	PatternBoolean                  *Boolean                      `json:"patternBoolean,omitempty"`
// 	PatternCanonical                *Canonical                    `json:"patternCanonical,omitempty"`
// 	PatternCode                     *Code                         `json:"patternCode,omitempty"`
// 	PatternDate                     *Date                         `json:"patternDate,omitempty"`
// 	PatternDateTime                 *DateTime                     `json:"patternDateTime,omitempty"`
// 	PatternDecimal                  *Decimal                      `json:"patternDecimal,omitempty"`
// 	PatternID                       *ID                           `json:"patternId,omitempty"`
// 	PatternInstant                  *Instant                      `json:"patternInstant,omitempty"`
// 	PatternInteger                  *Integer                      `json:"patternInteger,omitempty"`
// 	PatternMarkdown                 *Markdown                     `json:"patternMarkdown,omitempty"`
// 	PatternOID                      *OID                          `json:"patternOid,omitempty"`
// 	PatternPositiveInt              *PositiveInt                  `json:"patternPositiveInt,omitempty"`
// 	PatternString                   *String                       `json:"patternString,omitempty"`
// 	PatternTime                     *Time                         `json:"patternTime,omitempty"`
// 	PatternUnsignedInt              *UnsignedInt                  `json:"patternUnsignedInt,omitempty"`
// 	PatternURI                      *URI                          `json:"patternUri,omitempty"`
// 	PatternURL                      *URL                          `json:"patternUrl,omitempty"`
// 	PatternUUID                     *UUID                         `json:"patternUuid,omitempty"`
// 	PatternAddress                  *Address                      `json:"patternAddress,omitempty"`
// 	PatternAge                      *Age                          `json:"patternAge,omitempty"`
// 	PatternAnnotation               *Annotation                   `json:"patternAnnotation,omitempty"`
// 	PatternAttachment               *Attachment                   `json:"patternAttachment,omitempty"`
// 	PatternCodeableConcept          *CodeableConcept              `json:"patternCodeableConcept,omitempty"`
// 	PatternCoding                   *Coding                       `json:"patternCoding,omitempty"`
// 	PatternContactPoint             *ContactPoint                 `json:"patternContactPoint,omitempty"`
// 	PatternCount                    *Count                        `json:"patternCount,omitempty"`
// 	PatternDistance                 *Distance                     `json:"patternDistance,omitempty"`
// 	PatternDuration                 *Duration                     `json:"patternDuration,omitempty"`
// 	PatternHumanName                *HumanName                    `json:"patternHumanName,omitempty"`
// 	PatternIdentifier               *Identifier                   `json:"patternIdentifier,omitempty"`
// 	PatternMoney                    *Money                        `json:"patternMoney,omitempty"`
// 	PatternPeriod                   *Period                       `json:"patternPeriod,omitempty"`
// 	PatternQuantity                 *Quantity                     `json:"patternQuantity,omitempty"`
// 	PatternRange                    *Range                        `json:"patternRange,omitempty"`
// 	PatternRatio                    *Ratio                        `json:"patternRatio,omitempty"`
// 	PatternReference                *Reference                    `json:"patternReference,omitempty"`
// 	PatternSampledData              *SampledData                  `json:"patternSampledData,omitempty"`
// 	PatternSignature                *Signature                    `json:"patternSignature,omitempty"`
// 	PatternTiming                   *Timing                       `json:"patternTiming,omitempty"`
// 	PatternContactDetail            *ContactDetail                `json:"patternContactDetail,omitempty"`
// 	PatternContributor              *Contributor                  `json:"patternContributor,omitempty"`
// 	PatternDataRequirement          *DataRequirement              `json:"patternDataRequirement,omitempty"`
// 	PatternExpression               *Expression                   `json:"patternExpression,omitempty"`
// 	PatternParameterDefinition      *ParameterDefinition          `json:"patternParameterDefinition,omitempty"`
// 	PatternRelatedArtifact          *RelatedArtifact              `json:"patternRelatedArtifact,omitempty"`
// 	PatternTriggerDefinition        *TriggerDefinition            `json:"patternTriggerDefinition,omitempty"`
// 	PatternUsageContext             *UsageContext                 `json:"patternUsageContext,omitempty"`
// 	PatternDosage                   *Dosage                       `json:"patternDosage,omitempty"`
// 	Example                         []ElementDefinitionExample    `json:"example,omitempty"`
// 	MinValueDate                    *Date                         `json:"minValueDate,omitempty"`
// 	MinValueDateTime                *DateTime                     `json:"minValueDateTime,omitempty"`
// 	MinValueInstant                 *Instant                      `json:"minValueInstant,omitempty"`
// 	MinValueTime                    *Time                         `json:"minValueTime,omitempty"`
// 	MinValueDecimal                 *Decimal                      `json:"minValueDecimal,omitempty"`
// 	MinValueInteger                 *Integer                      `json:"minValueInteger,omitempty"`
// 	MinValuePositiveInt             *PositiveInt                  `json:"minValuePositiveInt,omitempty"`
// 	MinValueUnsignedInt             *UnsignedInt                  `json:"minValueUnsignedInt,omitempty"`
// 	MinValueQuantity                *Quantity                     `json:"minValueQuantity,omitempty"`
// 	MaxValueDate                    *Date                         `json:"maxValueDate,omitempty"`
// 	MaxValueDateTime                *DateTime                     `json:"maxValueDateTime,omitempty"`
// 	MaxValueInstant                 *Instant                      `json:"maxValueInstant,omitempty"`
// 	MaxValueTime                    *Time                         `json:"maxValueTime,omitempty"`
// 	MaxValueDecimal                 *Decimal                      `json:"maxValueDecimal,omitempty"`
// 	MaxValueInteger                 *Integer                      `json:"maxValueInteger,omitempty"`
// 	MaxValuePositiveInt             *PositiveInt                  `json:"maxValuePositiveInt,omitempty"`
// 	MaxValueUnsignedInt             *UnsignedInt                  `json:"maxValueUnsignedInt,omitempty"`
// 	MaxValueQuantity                *Quantity                     `json:"maxValueQuantity,omitempty"`
// 	MaxLength                       *Integer                      `json:"maxLength,omitempty"`
// 	Condition                       []ID                          `json:"condition,omitempty"`
// 	Constraint                      []ElementDefinitionConstraint `json:"constraint,omitempty"`
// 	MustSupport                     *Boolean                      `json:"mustSupport,omitempty"`
// 	IsModifier                      *Boolean                      `json:"isModifier,omitempty"`
// 	IsModifierReason                *String                       `json:"isModifierReason,omitempty"`
// 	IsSummary                       *Boolean                      `json:"isSummary,omitempty"`
// 	Binding                         *ElementDefinitionBinding     `json:"binding,omitempty"`
// 	Mapping                         []ElementDefinitionMapping    `json:"mapping,omitempty"`
// 	BackboneElement
// }
