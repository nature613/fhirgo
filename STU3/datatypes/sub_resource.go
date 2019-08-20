package datatypes

import (
	"encoding/json"
)

// PatientContact subResource
type PatientContact struct {
	Relationship []CodeableConcept `json:"relationship,omitempty"`
	Name         *HumanName        `json:"name,omitempty"`
	Telecom      []ContactPoint    `json:"telecom,omitempty"`
	Address      *Address          `json:"address,omitempty"`
	Gender       *Code             `json:"gender,omitempty"`
	Organization *Reference        `json:"organization,omitempty"`
	Period       *Period           `json:"period,omitempty"`
}

// PatientCommunication subResource
type PatientCommunication struct {
	Language  *CodeableConcept `json:"language,omitempty"`
	Preferred *Boolean         `json:"preferred,omitempty"`
}

// PatientLink subResource
type PatientLink struct {
	Other *Reference `json:"other,omitempty"`
	Type  *Code      `json:"type,omitempty"`
}

// EncounterStatusHistory subResource
type EncounterStatusHistory struct {
	Status *Code   `json:"status,omitempty"`
	Period *Period `json:"period,omitempty"`
}

// EncounterClassHistory subResource
type EncounterClassHistory struct {
	Class  *Coding `json:"class,omitempty"`
	Period *Period `json:"period,omitempty"`
}

// EncounterParticipant subResource
type EncounterParticipant struct {
	Type       []CodeableConcept `json:"type,omitempty"`
	Period     *Period           `json:"period,omitempty"`
	Individual *Reference        `json:"individual,omitempty"`
}

// EncounterDiagnosis subResource
type EncounterDiagnosis struct {
	Condition *Reference       `json:"condition,omitempty"`
	Use       *CodeableConcept `json:"use,omitempty"`
	Rank      *PositiveInt     `json:"rank,omitempty"`
}

// EncounterHospitalization subResource
type EncounterHospitalization struct {
	PreAdmissionIdentifier *Identifier      `json:"preAdmissionIdentifier,omitempty"`
	Origin                 *Reference       `json:"origin,omitempty"`
	AdmitSource            *CodeableConcept `json:"admitSource,omitempty"`
	ReAdmission            *CodeableConcept `json:"reAdmission,omitempty"`
	DietPreference         *CodeableConcept `json:"dietPreference,omitempty"`
	SpecialCourtesy        *CodeableConcept `json:"specialCourtesy,omitempty"`
	SpecialArrangement     *CodeableConcept `json:"specialArrangement,omitempty"`
	Destination            *Reference       `json:"destination,omitempty"`
	DischargeDisposition   *CodeableConcept `json:"dischargeDisposition,omitempty"`
}

// EncounterLocation subResource
type EncounterLocation struct {
	Location     *Reference       `json:"location,omitempty"`
	Status       *Code            `json:"status,omitempty"`
	PhysicalType *CodeableConcept `json:"physicalType,omitempty"`
	Period       *Period          `json:"period,omitempty"`
}

// ObservationReferenceRange subResource
type ObservationReferenceRange struct {
	Low       *SimpleQuantity   `json:"low,omitempty"`
	High      *SimpleQuantity   `json:"high,omitempty"`
	Type      *CodeableConcept  `json:"type,omitempty"`
	AppliesTo []CodeableConcept `json:"appliesTo,omitempty"`
	Age       *Range            `json:"age,omitempty"`
	Text      *String           `json:"text,omitempty"`
}

// ObservationComponent subResource
type ObservationComponent struct {
	Code                 *CodeableConcept            `json:"code,omitempty"`
	ValueQuantity        *Quantity                   `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *CodeableConcept            `json:"valueCodeableConcept,omitempty"`
	ValueString          *String                     `json:"valueString,omitempty"`
	ValueBoolean         *Boolean                    `json:"valueBoolean,omitempty"`
	ValueInteger         *Integer                    `json:"valueInteger,omitempty"`
	ValueRange           *Range                      `json:"valueRange,omitempty"`
	ValueRatio           *Ratio                      `json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData                `json:"valueSampledData,omitempty"`
	ValueTime            *Time                       `json:"valueTime,omitempty"`
	ValueDateTime        *DateTime                   `json:"valueDateTime,omitempty"`
	ValuePeriod          *Period                     `json:"valuePeriod,omitempty"`
	DataAbsentReason     *CodeableConcept            `json:"dataAbsentReason,omitempty"`
	Interpretation       []CodeableConcept           `json:"interpretation,omitempty"`
	ReferenceRange       []ObservationReferenceRange `json:"referenceRange,omitempty"`
}

// BundleLink subResource
type BundleLink struct {
	Relation *String `json:"relation,omitempty"`
	URL      *URI    `json:"url,omitempty"`
}

// BundleEntrySearch subResource
type BundleEntrySearch struct {
	Mode  *Code    `json:"mode,omitempty"`
	Score *Decimal `json:"score,omitempty"`
}

// BundleEntryRequest subResource
type BundleEntryRequest struct {
	Method          *Code    `json:"method,omitempty"`
	URL             *URI     `json:"url,omitempty"`
	IfNoneMatch     *String  `json:"ifNoneMatch,omitempty"`
	IfModifiedSince *Instant `json:"ifModifiedSince,omitempty"`
	IfMatch         *String  `json:"ifMatch,omitempty"`
	IfNoneExist     *String  `json:"ifNoneExist,omitempty"`
}

// BundleEntryResponse subResource
type BundleEntryResponse struct {
	Status       *String     `json:"status,omitempty"`
	Location     *URI        `json:"location,omitempty"`
	Etag         *String     `json:"etag,omitempty"`
	LastModified *Instant    `json:"lastModified,omitempty"`
	Outcome      interface{} `json:"outcome,omitempty"` // Resource
}

// BundleEntry subResource
type BundleEntry struct {
	Link     []BundleLink         `json:"link,omitempty"`
	FullURL  *URI                 `json:"fullUrl,omitempty"`
	Resource interface{}          `json:"resource,omitempty"` // A resource in the bundle
	Search   *BundleEntrySearch   `json:"search,omitempty"`
	Request  *BundleEntryRequest  `json:"request,omitempty"`
	Response *BundleEntryResponse `json:"response,omitempty"`
}

// GetResourceType returns the resource type
func (b *BundleEntry) GetResourceType() (string, error) {
	m := make(map[string]interface{})
	bt, err := json.Marshal(b.Resource)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(bt, &m); err != nil {
		return "", err
	}
	return m["resourceType"].(string), nil
}

// AllergyIntoleranceReaction subResource
type AllergyIntoleranceReaction struct {
	Substance     *CodeableConcept  `json:"substance,omitempty"`
	Manifestation []CodeableConcept `json:"manifestation,omitempty"`
	Description   *String           `json:"description,omitempty"`
	Onset         *DateTime         `json:"onset,omitempty"`
	Severity      *Code             `json:"severity,omitempty"`
	ExposureRoute *CodeableConcept  `json:"exposureRoute,omitempty"`
	Note          []Annotation      `json:"note,omitempty"`
}

// FamilyMemberHistoryCondition subResource
type FamilyMemberHistoryCondition struct {
	Code        *CodeableConcept `json:"code,omitempty"`
	Outcome     *CodeableConcept `json:"outcome,omitempty"`
	OnsetAge    *Age             `json:"onsetAge,omitempty"`
	OnsetRange  *Range           `json:"onsetRange,omitempty"`
	OnsetPeriod *Period          `json:"onsetPeriod,omitempty"`
	OnsetString *String          `json:"onsetString,omitempty"`
	Note        []Annotation     `json:"note,omitempty"`
}

// ConditionStage subResource
type ConditionStage struct {
	Summary    *CodeableConcept `json:"summary,omitempty"`
	Assessment []Reference      `json:"assessment,omitempty"`
}

// ConditionEvidence subResource
type ConditionEvidence struct {
	Code   []CodeableConcept `json:"code,omitempty"`
	Detail []Reference       `json:"detail,omitempty"`
}

// DiagnosticReportPerformer subResource
type DiagnosticReportPerformer struct {
	Role  *CodeableConcept `json:"role,omitempty"`
	Actor *Reference       `json:"actor,omitempty"`
}

// DiagnosticReportImage subResource
type DiagnosticReportImage struct {
	Comment *String    `json:"comment,omitempty"`
	Link    *Reference `json:"link,omitempty"`
}

// SpecimenCollection subResource
type SpecimenCollection struct {
	Collector         *Reference       `json:"collector,omitempty"`
	CollectedDateTime *DateTime        `json:"collectedDateTime,omitempty"`
	CollectedPeriod   *Period          `json:"collectedPeriod,omitempty"`
	Quantity          *SimpleQuantity  `json:"quantity,omitempty"`
	Method            *CodeableConcept `json:"method,omitempty"`
	BodySite          *CodeableConcept `json:"bodySite,omitempty"`
}

// SpecimenProcessing subResource
type SpecimenProcessing struct {
	Description  *String          `json:"description,omitempty"`
	Procedure    *CodeableConcept `json:"procedure,omitempty"`
	Additive     []Reference      `json:"additive,omitempty"`
	TimeDateTime *DateTime        `json:"timeDateTime,omitempty"`
	TimePeriod   *Period          `json:"timePeriod,omitempty"`
}

// SpecimenContainer subResource
type SpecimenContainer struct {
	Identifier              []Identifier     `json:"identifier,omitempty"`
	Description             *String          `json:"description,omitempty"`
	Type                    *CodeableConcept `json:"type,omitempty"`
	Capacity                *SimpleQuantity  `json:"capacity,omitempty"`
	SpecimenQuantity        *SimpleQuantity  `json:"specimenQuantity,omitempty"`
	AdditiveCodeableConcept *CodeableConcept `json:"additiveCodeableConcept,omitempty"`
	AdditiveReference       *Reference       `json:"additiveReference,omitempty"`
}

// CarePlanActivity subResource
type CarePlanActivity struct {
	OutcomeCodeableConcept []CodeableConcept       `json:"outcomeCodeableConcept,omitempty"`
	OutcomeReference       []Reference             `json:"outcomeReference,omitempty"`
	Progress               []Annotation            `json:"progress,omitempty"`
	Reference              *Reference              `json:"reference,omitempty"`
	Detail                 *CarePlanActivityDetail `json:"detail,omitempty"`
}

// CarePlanActivityDetail subResource
type CarePlanActivityDetail struct {
	Category               *CodeableConcept  `json:"category,omitempty"`
	Definition             *Reference        `json:"definition,omitempty"`
	Code                   *CodeableConcept  `json:"code,omitempty"`
	ReasonCode             []CodeableConcept `json:"reasonCode,omitempty"`
	ReasonReference        []Reference       `json:"reasonReference,omitempty"`
	Goal                   []Reference       `json:"goal,omitempty"`
	Status                 *Code             `json:"status,omitempty"`
	StatusReason           *String           `json:"statusReason,omitempty"`
	Prohibited             *Boolean          `json:"prohibited,omitempty"`
	ScheduledTiming        *Timing           `json:"scheduledTiming,omitempty"`
	ScheduledPeriod        *Period           `json:"scheduledPeriod,omitempty"`
	ScheduledString        *String           `json:"scheduledString,omitempty"`
	Location               *Reference        `json:"location,omitempty"`
	Performer              []Reference       `json:"performer,omitempty"`
	ProductCodeableConcept *CodeableConcept  `json:"productCodeableConcept,omitempty"`
	ProductReference       *Reference        `json:"productReference,omitempty"`
	DailyAmount            *SimpleQuantity   `json:"dailyAmount,omitempty"`
	Quantity               *SimpleQuantity   `json:"quantity,omitempty"`
	Description            *String           `json:"description,omitempty"`
}
