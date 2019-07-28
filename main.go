package main

import (
	"encoding/json"
	"fmt"

	r "github.com/monarko/fhirgo/resources"
)

func main() {
	pat, err := r.NewPatient()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	j, err := json.Marshal(pat)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println(string(j))
}
