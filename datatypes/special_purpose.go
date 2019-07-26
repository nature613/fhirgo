package datatypes

// Reference FHIR Type
type Reference struct {
	Reference  String     `json:"reference"`
	Type       URI        `json:"type"`
	Identifier Identifier `json:"identifier"`
	Display    String     `json:"display"`
}
