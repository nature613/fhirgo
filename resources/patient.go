package resources

import (
	d "github.com/monarko/fhirgo/datatypes"
)

// PatientContact subResource
type PatientContact struct {
	Relationship []d.CodeableConcept `json:"relationship"`
	Name         *d.HumanName        `json:"name"`
	Telecom      []d.ContactPoint    `json:"telecom"`
	Address      *d.Address          `json:"address"`
	Gender       *d.Code             `json:"gender"`
	Organization *d.Reference        `json:"organization"`
	Period       *d.Period           `json:"period"`
}

// PatientCommunication subResource
type PatientCommunication struct {
	Language  *d.CodeableConcept `json:"language"`
	Preferred *d.Boolean         `json:"preferred"`
}

// PatientLink subResource
type PatientLink struct {
	Other *d.Reference `json:"other"`
	Type  *d.Code      `json:"type"`
}

// Patient resource
type Patient struct {
	Domain
	Identifier           []d.Identifier         `json:"identifier,omitempty"`
	Active               *d.Boolean             `json:"active,omitempty"`
	Name                 []d.HumanName          `json:"name,omitempty"`
	Telecom              []d.ContactPoint       `json:"telecom,omitempty"`
	Gender               *d.Code                `json:"gender,omitempty"`
	BirthDate            *d.Date                `json:"birthDate,omitempty"`
	DeceasedBoolean      *d.Boolean             `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime     *d.DateTime            `json:"deceasedDateTime,omitempty"`
	Address              []d.Address            `json:"address,omitempty"`
	MaritalStatus        *d.CodeableConcept     `json:"maritalStatus,omitempty"`
	MultipleBirthBoolean *d.Boolean             `json:"multipleBirthBoolean,omitempty"`
	MultipleBirthInteger *d.Integer             `json:"multipleBirthInteger,omitempty"`
	Photo                []d.Attachment         `json:"photo,omitempty"`
	Contact              []PatientContact       `json:"contact,omitempty"`
	Communication        []PatientCommunication `json:"communication,omitempty"`
	GeneralPractitioner  []d.Reference          `json:"generalPractitioner,omitempty"`
	ManagingOrganization *d.Reference           `json:"managingOrganization,omitempty"`
	Link                 []PatientLink          `json:"link,omitempty"`
}

// NewPatient returns a valid Patient resource
func NewPatient() (Patient, error) {
	d, err := NewDomain("Patient")
	if err != nil {
		return Patient{}, nil
	}
	p := Patient{}
	p.Domain = d

	return p, nil
}
