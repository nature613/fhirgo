package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	r "github.com/monarko/fhirgo/resources"
	"github.com/xeipuuv/gojsonschema"
)

func main() {
	pat, err := r.NewPatient()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	j, err := json.MarshalIndent(pat, "", "    ")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file := "file://" + filepath.Join(dir, "schema", "fhir.4.0.schema.json")

	schemaLoader := gojsonschema.NewReferenceLoader(file)
	docLoader := gojsonschema.NewBytesLoader(j)

	result, err := gojsonschema.Validate(schemaLoader, docLoader)
	if err != nil {
		fmt.Println("Result err:", err.Error())
	}

	if result.Valid() {
		fmt.Println("The document is valid")
		fmt.Println(string(j))
	} else {
		fmt.Println("The document is not valid. see errors:")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
