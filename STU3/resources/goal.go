package resources

import "github.com/monarko/fhirgo/schema"

// Goal resource
type Goal struct {
}

// Validate returns a check against schema
func (g *Goal) Validate() (bool, []error) {
	return schema.ValidateResource(*g, "3")
}
