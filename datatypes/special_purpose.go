package datatypes

// Reference FHIR Type
type Reference struct {
	Reference  String     `json:"reference"`
	Type       URI        `json:"type"`
	Identifier Identifier `json:"identifier"`
	Display    String     `json:"display"`
}

// Meta FHIR Type
type Meta struct {
	VersionID   ID          `json:"versionId"`
	LastUpdated Instant     `json:"lastUpdated"`
	Source      URI         `json:"source"`
	Profile     []Canonical `json:"profile"`
	Security    []Coding    `json:"security"`
	Tag         []Coding    `json:"tag"`
}
