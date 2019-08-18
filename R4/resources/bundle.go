package resources

import (
	"encoding/json"

	"github.com/monarko/fhirgo/helpers"

	d "github.com/monarko/fhirgo/R4/datatypes"
	"github.com/monarko/fhirgo/schema"
)

// Bundle resource
type Bundle struct {
	Base
	Identifier *d.Identifier   `json:"identifier,omitempty"`
	Type       *d.Code         `json:"type,omitempty"`
	Timestamp  *d.Instant      `json:"timestamp,omitempty"`
	Total      *d.UnsignedInt  `json:"total,omitempty"`
	Link       []d.BundleLink  `json:"link,omitempty"`
	Entry      []d.BundleEntry `json:"entry,omitempty"`
	Signature  *d.Signature    `json:"signature,omitempty"`
}

// Validate returns a check against schema
func (b *Bundle) Validate() (bool, []error) {
	return schema.ValidateResource(*b)
}

// Transform Bundle
func (b *Bundle) Transform(resultSet ...string) (Bundle, error) {
	entries := make([]d.BundleEntry, 0)
	for _, entry := range b.Entry {
		e := &entry
		t, err := e.GetResourceType()
		if err != nil {
			return *b, err
		}
		bt, err := json.Marshal(e.Resource)
		if err != nil {
			return *b, err
		}
		if resultSet != nil && !helpers.SliceContainsString(resultSet, t) {
			continue
		}
		if t == "Patient" {
			p := &Patient{}
			if err := json.Unmarshal(bt, &p); err != nil {
				return *b, err
				// continue
			}
			entry.Resource = *p
			entries = append(entries, entry)
		} else if t == "Encounter" {
			p := &Encounter{}
			if err := json.Unmarshal(bt, &p); err != nil {
				return *b, err
				// continue
			}
			entry.Resource = *p
			entries = append(entries, entry)
		} else if t == "Observation" {
			p := &Observation{}
			if err := json.Unmarshal(bt, &p); err != nil {
				return *b, err
				// continue
			}
			entry.Resource = *p
			entries = append(entries, entry)
		} else {
			entries = append(entries, entry)
		}
	}
	b.Entry = entries
	total := d.UnsignedInt(len(entries))
	b.Total = &total

	return *b, nil
}
