package resources

import (
	d "github.com/monarko/fhirgo/STU3/datatypes"
	"github.com/monarko/fhirgo/schema"
)

// Encounter resource
type Encounter struct {
	Domain
	Identifier       []d.Identifier              `json:"identifier,omitempty"`
	Status           *d.Code                     `json:"status,omitempty"`
	StatusHistory    []d.EncounterStatusHistory  `json:"statusHistory,omitempty"`
	Class            *d.Coding                   `json:"class,omitempty"`
	ClassHistory     []d.EncounterClassHistory   `json:"classHistory,omitempty"`
	Type             []d.CodeableConcept         `json:"type,omitempty"`
	Priority         *d.CodeableConcept          `json:"priority,omitempty"`
	Subject          *d.Reference                `json:"subject,omitempty"`
	EpisodeOfCare    []d.Reference               `json:"episodeOfCare,omitempty"`
	IncomingReferral []d.Reference               `json:"incomingReferral,omitempty"`
	Participant      []d.EncounterParticipant    `json:"participant,omitempty"`
	Appointment      *d.Reference                `json:"appointment,omitempty"`
	Period           *d.Period                   `json:"period,omitempty"`
	Length           *d.Duration                 `json:"length,omitempty"`
	Reason           []d.CodeableConcept         `json:"reason,omitempty"`
	Diagnosis        []d.EncounterDiagnosis      `json:"diagnosis,omitempty"`
	Account          []d.Reference               `json:"account,omitempty"`
	Hospitalization  *d.EncounterHospitalization `json:"hospitalization,omitempty"`
	Location         []d.EncounterLocation       `json:"location,omitempty"`
	ServiceProvider  *d.Reference                `json:"serviceProvider,omitempty"`
	PartOf           *d.Reference                `json:"partOf,omitempty"`
}

// Validate returns a check against schema
func (e *Encounter) Validate() (bool, []error) {
	return schema.ValidateResource(*e, "3")
}
