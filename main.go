package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	d "github.com/monarko/fhirgo/datatypes"
	r "github.com/monarko/fhirgo/resources"
	"github.com/xeipuuv/gojsonschema"
)

func main() {
	var identifier []d.Identifier

	var active *d.Boolean

	var name []d.HumanName
	use := d.Code("official")
	family := d.String("Doe")
	given := []d.String{"John", "Michael"}
	nm := d.HumanName{Use: &use, Family: &family, Given: given}
	name = make([]d.HumanName, 0)
	name = append(name, nm)

	var telecom []d.ContactPoint

	var gender *d.Code
	gender = new(d.Code)
	*gender = "male"

	var birthDate *d.Date
	birthDate = new(d.Date)
	*birthDate = "1977-09-14"

	var deceasedBoolean *d.Boolean

	var deceasedDateTime *d.DateTime

	var address []d.Address

	var maritalStatus *d.CodeableConcept

	var multipleBirthBoolean *d.Boolean

	var multipleBirthInteger *d.Integer

	var photo []d.Attachment

	var contact []d.PatientContact

	var communication []d.PatientCommunication

	var generalPractitioner []d.Reference

	var managingOrganization *d.Reference

	var link []d.PatientLink

	pat, err := r.NewPatient(
		identifier,
		active,
		name,
		telecom,
		gender,
		birthDate,
		deceasedBoolean,
		deceasedDateTime,
		address,
		maritalStatus,
		multipleBirthBoolean,
		multipleBirthInteger,
		photo,
		contact,
		communication,
		generalPractitioner,
		managingOrganization,
		link,
	)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	j, err := json.MarshalIndent(pat, "", "    ")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println(string(j))

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
	} else {
		fmt.Println("The document is not valid. see errors:")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
