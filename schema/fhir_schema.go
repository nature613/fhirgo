package schema

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	j "github.com/xeipuuv/gojsonschema"
)

// FHIRSchema returns a valid FHIR schema
func FHIRSchema(version string) (j.JSONLoader, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file := "file://" + filepath.Join(dir, "schema", "fhir."+version+".schema.json")

	schemaLoader := j.NewReferenceLoader(file)

	return schemaLoader, nil
}

// ValidateResource against a valid schema
func ValidateResource(i interface{}, schemaVersion string) (bool, []error) {
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
