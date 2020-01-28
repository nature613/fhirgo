package resources

import (
    d "github.com/monarko/fhirgo/STU3/datatypes"
    "github.com/monarko/fhirgo/schema"
)

// Group resource
type Group struct {
    Domain
    Identifier     []d.Identifier          `json:"identifier,omitempty"`
    Active         *d.Boolean              `json:"active,omitempty"`
    Type           *d.Code                 `json:"type,omitempty"`
    Actual         *d.Boolean              `json:"actual,omitempty"`
    Code           *d.CodeableConcept      `json:"code,omitempty"`
    Name           *d.String               `json:"name,omitempty"`
    Quantity       *d.UnsignedInt          `json:"quantity,omitempty"`
    characteristic []d.GroupCharacteristic `json:"characteristic,omitempty"`
    Member         []d.GroupMember         `json:"member,omitempty"`
}

// Validate returns a check against schema
func (g *Group) Validate() (bool, []error) {
    return schema.ValidateResource(*g, "3")
}
