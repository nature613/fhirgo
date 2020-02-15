package resources

import (
    d "github.com/monarko/fhirgo/STU3/datatypes"
    "github.com/monarko/fhirgo/schema"
)

// Group resource
type Group struct {
    Domain
    Identifier     []d.Identifier        `json:"identifier,omitempty"`
    Active         *d.Boolean            `json:"active,omitempty"`
    Type           *d.Code               `json:"type,omitempty"`
    Actual         *d.Boolean            `json:"actual,omitempty"`
    Code           *d.CodeableConcept    `json:"code,omitempty"`
    Name           *d.String             `json:"name,omitempty"`
    Quantity       *d.UnsignedInt        `json:"quantity,omitempty"`
    Characteristic []GroupCharacteristic `json:"characteristic,omitempty"`
    Member         []GroupMember         `json:"member,omitempty"`
}

// Validate returns a check against schema
func (g *Group) Validate() (bool, []error) {
    return schema.ValidateResource(*g, "3")
}

// GroupCharacteristic sub-resource
type GroupCharacteristic struct {
    Code                 *d.CodeableConcept `json:"code,omitempty"`
    ValueCodeableConcept *d.CodeableConcept `json:"valueCodeableConcept,omitempty"`
    ValueBoolean         *d.Boolean         `json:"valueBoolean,omitempty"`
    ValueQuantity        *d.Quantity        `json:"valueQuantity,omitempty"`
    ValueRange           *d.Range           `json:"valueRange,omitempty"`
    Exclude              *d.Boolean         `json:"exclude,omitempty"`
    Period               *d.Period          `json:"period,omitempty"`
}

// GroupMember sub-resource
type GroupMember struct {
    Entity   *d.Reference `json:"entity,omitempty"`
    Period   *d.Period    `json:"period,omitempty"`
    Inactive *d.Boolean   `json:"inactive,omitempty"`
}
