package resources

import "github.com/monarko/fhirgo/schema"

// RiskAssessment resource
type RiskAssessment struct {
}

// Validate returns a check against schema
func (r *RiskAssessment) Validate() (bool, []error) {
	return schema.ValidateResource(*r, "3")
}
