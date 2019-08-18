package schema

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"

	j "github.com/xeipuuv/gojsonschema"
)

// fhirSchema returns a valid FHIR schema
func fhirSchema() (j.JSONLoader, error) {
	schemaPath := os.Getenv("FHIR_SCHEMA_PATH")
	if len(strings.TrimSpace(schemaPath)) == 0 {
		return nil, errors.New("Environment variable FHIR_SCHEMA_PATH is not set")
	}

	fi, err := os.Stat(schemaPath)
	if err != nil {
		return nil, err
	}

	file := "file://"
	switch mode := fi.Mode(); {
	case mode.IsDir():
		file += filepath.Join(schemaPath, "fhir.schema.json")
	case mode.IsRegular():
		file += schemaPath
	}

	schemaLoader := j.NewReferenceLoader(file)

	return schemaLoader, nil
}

// ValidateResource against a valid schema
func ValidateResource(i interface{}) (bool, []error) {
	errs := make([]error, 0)
	valid := false
	r, err := json.Marshal(i)
	if err != nil {
		errs = append(errs, err)
		return valid, errs
	}

	schemaLoader, err := fhirSchema()
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
