package resources

import (
	d "github.com/monarko/fhirgo/STU3/datatypes"
	"github.com/monarko/fhirgo/schema"
)

// Location resource
type Location struct {
	Domain
	Identifier           []d.Identifier     `json:"identifier,omitempty"`
	Status               *d.Code            `json:"status,omitempty"`
	OperationalStatus    *d.Coding          `json:"operationalStatus,omitempty"`
	Name                 *d.String          `json:"name,omitempty"`
	Alias                []d.String         `json:"alias,omitempty"`
	Description          *d.String          `json:"description,omitempty"`
	Mode                 *d.Code            `json:"mode,omitempty"`
	Type                 *d.CodeableConcept `json:"type,omitempty"`
	Telecom              []d.ContactPoint   `json:"telecom,omitempty"`
	Address              *d.Address         `json:"address,omitempty"`
	PhysicalType         *d.CodeableConcept `json:"physicalType,omitempty"`
	Position             *LocationPosition  `json:"position,omitempty"`
	ManagingOrganization *d.Reference       `json:"managingOrganization,omitempty"`
	PartOf               *d.Reference       `json:"partOf,omitempty"`
	Endpoint             []d.Reference      `json:"endpoint,omitempty"`
}

// Validate returns a check against schema
func (l *Location) Validate() (bool, []error) {
	return schema.ValidateResource(*l, "3")
}

// LocationPosition sub-resource
type LocationPosition struct {
	Longitude *d.Decimal `json:"longitude,omitempty"`
	Latitude  *d.Decimal `json:"latitude,omitempty"`
	Altitude  *d.Decimal `json:"altitude,omitempty"`
}
