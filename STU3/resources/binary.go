package resources

import (
    d "github.com/monarko/fhirgo/STU3/datatypes"
    "github.com/monarko/fhirgo/schema"
)

// Binary resource
type Binary struct {
    Base
    ContentType     *d.Code         `json:"identifier,omitempty"`
    SecurityContext *d.Reference    `json:"clinicalStatus,omitempty"`
    Content         *d.Base64Binary `json:"type,omitempty"`
}

// Validate returns a check against schema
func (a *Binary) Validate() (bool, []error) {
    return schema.ValidateResource(*a, "3")
}
