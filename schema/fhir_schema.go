package schema

import (
	"encoding/json"
	"errors"

	d "github.com/monarko/fhirgo/datatypes"
	j "github.com/xeipuuv/gojsonschema"
)

// FHIRSchema returns a valid FHIR schema
func FHIRSchema(path string) (j.JSONLoader, error) {
	file := "file://" + d.SchemaPath

	schemaLoader := j.NewReferenceLoader(file)

	return schemaLoader, nil
}

// ValidateResource against a valid schema
func ValidateResource(i interface{}, schemaFilePath string) (bool, []error) {
	errs := make([]error, 0)
	valid := false
	r, err := json.Marshal(i)
	if err != nil {
		errs = append(errs, err)
		return valid, errs
	}

	schemaLoader, err := FHIRSchema("4.0")
	if err != nil {
		errs = append(errs, err)
		return valid, errs
	}
	docLoader := j.NewBytesLoader(r)

	result, err := j.Validate(schemaLoader, docLoader)
	if err != nil {
		errs = append(errs, err)
		return valid, errs
	}

	if !result.Valid() {
		for _, desc := range result.Errors() {
			errs = append(errs, errors.New(desc.String()))
		}
	} else {
		valid = true
	}

	return valid, errs
}
