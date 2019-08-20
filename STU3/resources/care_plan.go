package resources

import "github.com/monarko/fhirgo/schema"

// CarePlan resource
type CarePlan struct {
}

// Validate returns a check against schema
func (c *CarePlan) Validate() (bool, []error) {
	return schema.ValidateResource(*c, "3")
}
