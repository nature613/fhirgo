package datatypes

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
