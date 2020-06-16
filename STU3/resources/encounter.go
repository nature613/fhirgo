package resources

import (
	d "github.com/monarko/fhirgo/STU3/datatypes"
	"github.com/monarko/fhirgo/schema"
)

// Encounter resource
type Encounter struct {
	Domain
	Identifier       []d.Identifier            `json:"identifier,omitempty"`
	Status           *d.Code                   `json:"status,omitempty"`
	StatusHistory    []EncounterStatusHistory  `json:"statusHistory,omitempty"`
	Class            *d.Coding                 `json:"class,omitempty"`
	ClassHistory     []EncounterClassHistory   `json:"classHistory,omitempty"`
	Type             []d.CodeableConcept       `json:"type,omitempty"`
	Priority         *d.CodeableConcept        `json:"priority,omitempty"`
	Subject          *d.Reference              `json:"subject,omitempty"`
	EpisodeOfCare    []d.Reference             `json:"episodeOfCare,omitempty"`
	IncomingReferral []d.Reference             `json:"incomingReferral,omitempty"`
	Participant      []EncounterParticipant    `json:"participant,omitempty"`
	Appointment      *d.Reference              `json:"appointment,omitempty"`
	Period           *d.Period                 `json:"period,omitempty"`
	Length           *d.Duration               `json:"length,omitempty"`
	Reason           []d.CodeableConcept       `json:"reason,omitempty"`
	Diagnosis        []EncounterDiagnosis      `json:"diagnosis,omitempty"`
	Account          []d.Reference             `json:"account,omitempty"`
	Hospitalization  *EncounterHospitalization `json:"hospitalization,omitempty"`
	Location         []EncounterLocation       `json:"location,omitempty"`
	ServiceProvider  *d.Reference              `json:"serviceProvider,omitempty"`
	PartOf           *d.Reference              `json:"partOf,omitempty"`
}

// Validate returns a check against schema
func (e *Encounter) Validate() (bool, []error) {
	return schema.ValidateResource(*e, "3")
}

// EncounterStatusHistory subResource
type EncounterStatusHistory struct {
	Status *d.Code   `json:"status,omitempty"`
	Period *d.Period `json:"period,omitempty"`
}

// EncounterClassHistory subResource
type EncounterClassHistory struct {
	Class  *d.Coding `json:"class,omitempty"`
	Period *d.Period `json:"period,omitempty"`
}

// EncounterParticipant subResource
type EncounterParticipant struct {
	Type       []d.CodeableConcept `json:"type,omitempty"`
	Period     *d.Period           `json:"period,omitempty"`
	Individual *d.Reference        `json:"individual,omitempty"`
}

// EncounterDiagnosis subResource
type EncounterDiagnosis struct {
	Condition *d.Reference       `json:"condition,omitempty"`
	Role      *d.CodeableConcept `json:"role,omitempty"`
	Rank      *d.PositiveInt     `json:"rank,omitempty"`
}

// EncounterHospitalization subResource
type EncounterHospitalization struct {
	PreAdmissionIdentifier *d.Identifier       `json:"preAdmissionIdentifier,omitempty"`
	Origin                 *d.Reference        `json:"origin,omitempty"`
	AdmitSource            *d.CodeableConcept  `json:"admitSource,omitempty"`
	ReAdmission            *d.CodeableConcept  `json:"reAdmission,omitempty"`
	DietPreference         []d.CodeableConcept `json:"dietPreference,omitempty"`
	SpecialCourtesy        []d.CodeableConcept `json:"specialCourtesy,omitempty"`
	SpecialArrangement     []d.CodeableConcept `json:"specialArrangement,omitempty"`
	Destination            *d.Reference        `json:"destination,omitempty"`
	DischargeDisposition   *d.CodeableConcept  `json:"dischargeDisposition,omitempty"`
}

// EncounterLocation subResource
type EncounterLocation struct {
	Location *d.Reference `json:"location,omitempty"`
	Status   *d.Code      `json:"status,omitempty"`
	Period   *d.Period    `json:"period,omitempty"`
}
