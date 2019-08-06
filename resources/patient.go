package resources

import (
	d "github.com/monarko/fhirgo/datatypes"
	"github.com/monarko/fhirgo/schema"
)

// Patient resource
type Patient struct {
	Domain
	Identifier           []d.Identifier           `json:"identifier,omitempty"`
	Active               *d.Boolean               `json:"active,omitempty"`
	Name                 []d.HumanName            `json:"name,omitempty"`
	Telecom              []d.ContactPoint         `json:"telecom,omitempty"`
	Gender               *d.Code                  `json:"gender,omitempty"`
	BirthDate            *d.Date                  `json:"birthDate,omitempty"`
	DeceasedBoolean      *d.Boolean               `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime     *d.DateTime              `json:"deceasedDateTime,omitempty"`
	Address              []d.Address              `json:"address,omitempty"`
	MaritalStatus        *d.CodeableConcept       `json:"maritalStatus,omitempty"`
	MultipleBirthBoolean *d.Boolean               `json:"multipleBirthBoolean,omitempty"`
	MultipleBirthInteger *d.Integer               `json:"multipleBirthInteger,omitempty"`
	Photo                []d.Attachment           `json:"photo,omitempty"`
	Contact              []d.PatientContact       `json:"contact,omitempty"`
	Communication        []d.PatientCommunication `json:"communication,omitempty"`
	GeneralPractitioner  []d.Reference            `json:"generalPractitioner,omitempty"`
	ManagingOrganization *d.Reference             `json:"managingOrganization,omitempty"`
	Link                 []d.PatientLink          `json:"link,omitempty"`
}

// NewPatient returns a valid Patient resource
func NewPatient(
	identifier []d.Identifier,
	active *d.Boolean,
	name []d.HumanName,
	telecom []d.ContactPoint,
	gender *d.Code,
	birthDate *d.Date,
	deceasedBoolean *d.Boolean,
	deceasedDateTime *d.DateTime,
	address []d.Address,
	maritalStatus *d.CodeableConcept,
	multipleBirthBoolean *d.Boolean,
	multipleBirthInteger *d.Integer,
	photo []d.Attachment,
	contact []d.PatientContact,
	communication []d.PatientCommunication,
	generalPractitioner []d.Reference,
	managingOrganization *d.Reference,
	link []d.PatientLink,
) Patient {
	p := Patient{}
	p.ResourceType = "Patient"
	p.Identifier = identifier
	p.Active = active
	p.Name = name
	p.Telecom = telecom
	p.Gender = gender
	p.BirthDate = birthDate
	p.DeceasedBoolean = deceasedBoolean
	p.DeceasedDateTime = deceasedDateTime
	p.Address = address
	p.MaritalStatus = maritalStatus
	p.MultipleBirthBoolean = multipleBirthBoolean
	p.MultipleBirthInteger = multipleBirthInteger
	p.Photo = photo
	p.Contact = contact
	p.Communication = communication
	p.GeneralPractitioner = generalPractitioner

	if managingOrganization != nil {
		p.ManagingOrganization = managingOrganization
	}
	p.Link = link

	return p
}

// Validate returns a check against schema
func (p *Patient) Validate() (bool, []error) {
	return schema.ValidateResource(*p, "4.0")
}
