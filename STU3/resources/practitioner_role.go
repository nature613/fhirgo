package resources

import (
    d "github.com/monarko/fhirgo/STU3/datatypes"
    "github.com/monarko/fhirgo/schema"
)

// PractitionerRole resource
type PractitionerRole struct {
    Domain
    Identifier             []d.Identifier                  `json:"identifier,omitempty"`
    Active                 *d.Boolean                      `json:"active,omitempty"`
    Period                 *d.Period                       `json:"period,omitempty"`
    Practitioner           *d.Reference                    `json:"practitioner,omitempty"`
    Organization           *d.Reference                    `json:"organization,omitempty"`
    Code                   []d.CodeableConcept             `json:"code,omitempty"`
    Specialty              []d.CodeableConcept             `json:"specialty,omitempty"`
    Location               []d.Reference                   `json:"location,omitempty"`
    HealthcareService      []d.Reference                   `json:"healthcareService,omitempty"`
    Telecom                []d.ContactPoint                `json:"telecom,omitempty"`
    AvailableTime          []PractitionerRoleAvailableTime `json:"availableTime,omitempty"`
    NotAvailable           []PractitionerRoleNotAvailable  `json:"notAvailable,omitempty"`
    AvailabilityExceptions *d.String                       `json:"availabilityExceptions,omitempty"`
    Endpoint               []d.Reference                   `json:"endpoint,omitempty"`
}

// Validate returns a check against schema
func (p *PractitionerRole) Validate() (bool, []error) {
    return schema.ValidateResource(*p, "3")
}

// PractitionerRoleAvailableTime sub-resource
type PractitionerRoleAvailableTime struct {
    DaysOfWeek         []d.Code   `json:"daysOfWeek,omitempty"`
    AllDay             *d.Boolean `json:"allDay,omitempty"`
    AvailableStartTime *d.Time    `json:"availableStartTime,omitempty"`
    AvailableEndTime   *d.Time    `json:"availableEndTime,omitempty"`
}

// PractitionerRoleNotAvailable sub-resource
type PractitionerRoleNotAvailable struct {
    Description *d.String `json:"description,omitempty"`
    During      *d.Period `json:"during,omitempty"`
}
