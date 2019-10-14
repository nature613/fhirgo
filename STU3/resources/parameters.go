package resources

import (
	d "github.com/monarko/fhirgo/STU3/datatypes"
	"github.com/monarko/fhirgo/schema"
)

// Parameters resource
type Parameters struct {
	Base
	Parameter []d.ParametersParameter `json:"parameter,omitempty"`
}

// Validate returns a check against schema
func (r *Parameters) Validate() (bool, []error) {
	return schema.ValidateResource(*r, "3")
}
