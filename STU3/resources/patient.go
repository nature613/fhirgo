package resources

import (
    d "github.com/monarko/fhirgo/STU3/datatypes"
    "github.com/monarko/fhirgo/schema"
)

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
    Animal               *PatientAnimal         `json:"animal,omitempty"`
    Communication        []PatientCommunication `json:"communication,omitempty"`
    GeneralPractitioner  []d.Reference          `json:"generalPractitioner,omitempty"`
    ManagingOrganization *d.Reference           `json:"managingOrganization,omitempty"`
    Link                 []PatientLink          `json:"link,omitempty"`
}

// Validate returns a check against schema
func (p *Patient) Validate() (bool, []error) {
    return schema.ValidateResource(*p, "3")
}

// PatientContact subResource
type PatientContact struct {
    Relationship []d.CodeableConcept `json:"relationship,omitempty"`
    Name         *d.HumanName        `json:"name,omitempty"`
    Telecom      []d.ContactPoint    `json:"telecom,omitempty"`
    Address      *d.Address          `json:"address,omitempty"`
    Gender       *d.Code             `json:"gender,omitempty"`
    Organization *d.Reference        `json:"organization,omitempty"`
    Period       *d.Period           `json:"period,omitempty"`
}

// PatientAnimal subResource
type PatientAnimal struct {
    Species      *d.CodeableConcept `json:"species,omitempty"`
    Breed        *d.CodeableConcept `json:"breed,omitempty"`
    GenderStatus *d.CodeableConcept `json:"genderStatus,omitempty"`
}

// PatientCommunication subResource
type PatientCommunication struct {
    Language  *d.CodeableConcept `json:"language,omitempty"`
    Preferred *d.Boolean         `json:"preferred,omitempty"`
}

// PatientLink subResource
type PatientLink struct {
    Other *d.Reference `json:"other,omitempty"`
    Type  *d.Code      `json:"type,omitempty"`
}
