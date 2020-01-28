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

// PatientAnimal subResource
type PatientAnimal struct {
    Species      *CodeableConcept `json:"species,omitempty"`
    Breed        *CodeableConcept `json:"breed,omitempty"`
    GenderStatus *CodeableConcept `json:"genderStatus,omitempty"`
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
    Role      *CodeableConcept `json:"role,omitempty"`
    Rank      *PositiveInt     `json:"rank,omitempty"`
}

// EncounterHospitalization subResource
type EncounterHospitalization struct {
    PreAdmissionIdentifier *Identifier       `json:"preAdmissionIdentifier,omitempty"`
    Origin                 *Reference        `json:"origin,omitempty"`
    AdmitSource            *CodeableConcept  `json:"admitSource,omitempty"`
    ReAdmission            *CodeableConcept  `json:"reAdmission,omitempty"`
    DietPreference         []CodeableConcept `json:"dietPreference,omitempty"`
    SpecialCourtesy        []CodeableConcept `json:"specialCourtesy,omitempty"`
    SpecialArrangement     []CodeableConcept `json:"specialArrangement,omitempty"`
    Destination            *Reference        `json:"destination,omitempty"`
    DischargeDisposition   *CodeableConcept  `json:"dischargeDisposition,omitempty"`
}

// EncounterLocation subResource
type EncounterLocation struct {
    Location *Reference `json:"location,omitempty"`
    Status   *Code      `json:"status,omitempty"`
    Period   *Period    `json:"period,omitempty"`
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

// ObservationRelated subResource
type ObservationRelated struct {
    Type   *Code      `json:"type,omitempty"`
    Target *Reference `json:"target,omitempty"`
}

// ObservationComponent subResource
type ObservationComponent struct {
    Code                 *CodeableConcept            `json:"code,omitempty"`
    ValueQuantity        *Quantity                   `json:"valueQuantity,omitempty"`
    ValueCodeableConcept *CodeableConcept            `json:"valueCodeableConcept,omitempty"`
    ValueString          *String                     `json:"valueString,omitempty"`
    ValueRange           *Range                      `json:"valueRange,omitempty"`
    ValueRatio           *Ratio                      `json:"valueRatio,omitempty"`
    ValueSampledData     *SampledData                `json:"valueSampledData,omitempty"`
    ValueAttachment      *Attachment                 `json:"valueAttachment,omitempty"`
    ValueTime            *Time                       `json:"valueTime,omitempty"`
    ValueDateTime        *DateTime                   `json:"valueDateTime,omitempty"`
    ValuePeriod          *Period                     `json:"valuePeriod,omitempty"`
    DataAbsentReason     *CodeableConcept            `json:"dataAbsentReason,omitempty"`
    Interpretation       *CodeableConcept            `json:"interpretation,omitempty"`
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

// GoalTarget subResource
type GoalTarget struct {
    Measure               *CodeableConcept `json:"measure,omitempty"`
    DetailQuantity        *Quantity        `json:"detailQuantity,omitempty"`
    DetailRange           *Range           `json:"detailRange,omitempty"`
    DetailCodeableConcept *CodeableConcept `json:"detailCodeableConcept,omitempty"`
    DueDate               *Date            `json:"dueDate,omitempty"`
    DueDuration           *Duration        `json:"dueDuration,omitempty"`
}

// RiskAssessmentPrediction subResource
type RiskAssessmentPrediction struct {
    Outcome            *CodeableConcept `json:"outcome,omitempty"`
    ProbabilityDecimal *Decimal         `json:"probabilityDecimal,omitempty"`
    ProbabilityRange   *Range           `json:"probabilityRange,omitempty"`
    QualitativeRisk    *CodeableConcept `json:"qualitativeRisk,omitempty"`
    RelativeRisk       *Decimal         `json:"relativeRisk,omitempty"`
    WhenPeriod         *Period          `json:"whenPeriod,omitempty"`
    WhenRange          *Range           `json:"whenRange,omitempty"`
    Rationale          *String          `json:"rationale,omitempty"`
}

// ReferralRequestRequester subResource
type ReferralRequestRequester struct {
    Agent      *Reference `json:"agent,omitempty"`
    OnBehalfOf *Reference `json:"onBehalfOf,omitempty"`
}

// ParametersParameter subResource
type ParametersParameter struct {
    Name                 *String          `json:"name,omitempty"`
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
    ValueSchedule        interface{}      `json:"valueSchedule,omitempty"`
    ValueReference       *Reference       `json:"valueReference,omitempty"`
    Resource             interface{}      `json:"resource,omitempty"`
    Part                 interface{}      `json:"part,omitempty"`
}

// OperationOutcome sub-resource
type OperationOutcomeIssue struct {
    Severity    *Code            `json:"severity,omitempty"`
    Code        *Code            `json:"code,omitempty"`
    Details     *CodeableConcept `json:"details,omitempty"`
    Diagnostics *String          `json:"diagnostics,omitempty"`
    Location    []String         `json:"location,omitempty"`
    Expression  []String         `json:"expression,omitempty"`
}

// PractitionerQualification sub-resource
type PractitionerQualification struct {
    Identifier []Identifier     `json:"identifier,omitempty"`
    Code       *CodeableConcept `json:"code,omitempty"`
    Period     *Period          `json:"period,omitempty"`
    Issuer     *Reference       `json:"issuer,omitempty"`
}

// PractitionerRoleAvailableTime sub-resource
type PractitionerRoleAvailableTime struct {
    DaysOfWeek         []Code   `json:"daysOfWeek,omitempty"`
    AllDay             *Boolean `json:"allDay,omitempty"`
    AvailableStartTime *Time    `json:"availableStartTime,omitempty"`
    AvailableEndTime   *Time    `json:"availableEndTime,omitempty"`
}

// PractitionerRoleNotAvailable sub-resource
type PractitionerRoleNotAvailable struct {
    Description *String `json:"description,omitempty"`
    During      *Period `json:"during,omitempty"`
}
