package resources

import (
	"encoding/json"

	"github.com/monarko/fhirgo/helpers"

	d "github.com/monarko/fhirgo/STU3/datatypes"
	"github.com/monarko/fhirgo/schema"
)

// Bundle resource
type Bundle struct {
	Base
	Identifier *d.Identifier  `json:"identifier,omitempty"`
	Type       *d.Code        `json:"type,omitempty"`
	Total      *d.UnsignedInt `json:"total,omitempty"`
	Link       []BundleLink   `json:"link,omitempty"`
	Entry      []BundleEntry  `json:"entry,omitempty"`
	Signature  *d.Signature   `json:"signature,omitempty"`
}

// Validate returns a check against schema
func (b *Bundle) Validate() (bool, []error) {
	return schema.ValidateResource(*b, "3")
}

// TODO: REMOVE THIS
// Transform Bundle
func (b *Bundle) Transform(resultSet ...string) (Bundle, error) {
	entries := make([]BundleEntry, 0)
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

// BundleLink subResource
type BundleLink struct {
	Relation *d.String `json:"relation,omitempty"`
	URL      *d.URI    `json:"url,omitempty"`
}

// BundleEntrySearch subResource
type BundleEntrySearch struct {
	Mode  *d.Code    `json:"mode,omitempty"`
	Score *d.Decimal `json:"score,omitempty"`
}

// BundleEntryRequest subResource
type BundleEntryRequest struct {
	Method          *d.Code    `json:"method,omitempty"`
	URL             *d.URI     `json:"url,omitempty"`
	IfNoneMatch     *d.String  `json:"ifNoneMatch,omitempty"`
	IfModifiedSince *d.Instant `json:"ifModifiedSince,omitempty"`
	IfMatch         *d.String  `json:"ifMatch,omitempty"`
	IfNoneExist     *d.String  `json:"ifNoneExist,omitempty"`
}

// BundleEntryResponse subResource
type BundleEntryResponse struct {
	Status       *d.String   `json:"status,omitempty"`
	Location     *d.URI      `json:"location,omitempty"`
	Etag         *d.String   `json:"etag,omitempty"`
	LastModified *d.Instant  `json:"lastModified,omitempty"`
	Outcome      interface{} `json:"outcome,omitempty"` // Resource
}

// BundleEntry subResource
type BundleEntry struct {
	Link     []BundleLink         `json:"link,omitempty"`
	FullURL  *d.URI               `json:"fullUrl,omitempty"`
	Resource interface{}          `json:"resource,omitempty"` // A resource in the bundle
	Search   *BundleEntrySearch   `json:"search,omitempty"`
	Request  *BundleEntryRequest  `json:"request,omitempty"`
	Response *BundleEntryResponse `json:"response,omitempty"`
}

// GetResourceType returns the resource type
func (b *BundleEntry) GetResourceType() (string, error) {
	m := make(map[string]interface{})
	bt, err := json.Marshal(b.Resource)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(bt, &m); err != nil {
		return "", err
	}
	return m["resourceType"].(string), nil
}
