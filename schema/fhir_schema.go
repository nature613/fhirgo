package schema

import (
	"encoding/json"
	"errors"

	j "github.com/xeipuuv/gojsonschema"
)

// fhirSchema returns a valid FHIR schema
func fhirSchema(version string) (j.JSONLoader, error) {
	var schema *[]byte
	if version == "3" {
		schema = &schemaSTU3
	} else if version == "4" {
		schema = &schemaR4
	} else {
		return nil, errors.New("invalid FHIR schema version")
	}

	schemaLoader := j.NewBytesLoader(*schema)

	return schemaLoader, nil
}

// ValidateResource against a valid schema
func ValidateResource(i interface{}, version string) (bool, []error) {
	errs := make([]error, 0)
	valid := false
	r, err := json.Marshal(i)
	if err != nil {
		errs = append(errs, err)
		return valid, errs
	}

	schemaLoader, err := fhirSchema(version)
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
