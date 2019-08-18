package datatypes

// Boolean FHIR type
type Boolean bool

// Integer FHIR type
type Integer int32

// String FHIR type
type String string

// Decimal FHIR type
type Decimal float64

// URI FHIR type
type URI string

// URL FHIR type
type URL URI

// Canonical FHIR type
type Canonical URI

// Base64Binary FHIR type
type Base64Binary string

// Instant FHIR type
type Instant string

// Date FHIR type
type Date string

// DateTime FHIR type
type DateTime string

// Time FHIR type
type Time string

// Code FHIR type
type Code String

// OID FHIR type
type OID URI

// ID FHIR type
type ID String

// Markdown FHIR type
type Markdown String

// UnsignedInt FHIR type
type UnsignedInt Integer

// PositiveInt FHIR type
type PositiveInt Integer

// UUID FHIR type
type UUID URI
