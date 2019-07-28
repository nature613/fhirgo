package resources

import (
	d "github.com/monarko/fhirgo/datatypes"
)

// PatientContact subResource
type PatientContact struct {
	Relationship []d.CodeableConcept `json:"relationship"`
	Name         d.HumanName         `json:"name"`
	Telecom      []d.ContactPoint    `json:"telecom"`
	Address      d.Address           `json:"address"`
	Gender       d.Code              `json:"gender"`
	Organization d.Reference         `json:"organization"`
	Period       d.Period            `json:"period"`
}

// PatientCommunication subResource
type PatientCommunication struct {
	Language  d.CodeableConcept `json:"language"`
	Preferred d.Boolean         `json:"preferred"`
}

// PatientLink subResource
type PatientLink struct {
	Other d.Reference `json:"other"`
	Type  d.Code      `json:"type"`
}

// Patient resource
type Patient struct {
	ResourceType         string                 `json:"resourceType"`
	Identifier           []d.Identifier         `json:"identifier"`
	Active               d.Boolean              `json:"active"`
	Name                 []d.HumanName          `json:"name"`
	Telecom              []d.ContactPoint       `json:"telecom"`
	Gender               d.Code                 `json:"gender"`
	BirthDate            d.Date                 `json:"birthDate"`
	DeceasedBoolean      d.Boolean              `json:"deceasedBoolean"`
	DeceasedDateTime     d.DateTime             `json:"deceasedDateTime"`
	Address              []d.Address            `json:"address"`
	MaritalStatus        d.CodeableConcept      `json:"maritalStatus"`
	MultipleBirthBoolean d.Boolean              `json:"multipleBirthBoolean"`
	MultipleBirthInteger d.Integer              `json:"multipleBirthInteger"`
	Photo                []d.Attachment         `json:"photo"`
	Contact              []PatientContact       `json:"contact"`
	Communication        []PatientCommunication `json:"communication"`
	GeneralPractitioner  []d.Reference          `json:"generalPractitioner"`
	ManagingOrganization d.Reference            `json:"managingOrganization"`
	Link                 []PatientLink          `json:"link"`
}

// NewPatient returns a valid Patient resource
func NewPatient() (Patient, error) {
	p := Patient{ResourceType: "Patient"}

	return p, nil
}
