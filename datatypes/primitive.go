package datatypes

import (
	"encoding/base64"
	"math"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrInvalidBooleanInput is returned when the given input is not a valid boolean
	ErrInvalidBooleanInput = errors.New("Invalid boolean value given as input")
	// ErrInvalidIntegerInput is returned when the given input is not a valid integer
	ErrInvalidIntegerInput = errors.New("Invalid integer value given as input")
	// ErrInvalidIntegerRange is returned when the given input is not in valid range
	ErrInvalidIntegerRange = errors.New("Integer value is out of range. Must be between -2,147,483,648 and 2,147,483,647")
	// ErrInvalidStringInput is returned when the given input is not a valid string
	ErrInvalidStringInput = errors.New("Invalid string value is given as input")
	// ErrInvalidStringLength is returned when the given string length is greater than 1024*1024
	ErrInvalidStringLength = errors.New("String length is out of range. Should be less than 1MB")
	// ErrInvalidDecimalInput is returned when the given input is not a valid float/decimal
	ErrInvalidDecimalInput = errors.New("Invalid decimal value given as input")
	// ErrInvalidDecimalRange is returned when decimal value is out of range
	ErrInvalidDecimalRange = errors.New("Decimal value is out of range")
	// ErrInvalidURI is returned when URI parsing failed
	ErrInvalidURI = errors.New("Invalid URI given as input")
	// ErrInvalidURL is returned when URL parsing failed
	ErrInvalidURL = errors.New("Invalid URL given as input")
	// ErrInvalidURLScheme is returned when protocol/scheme is empty
	ErrInvalidURLScheme = errors.New("A valid URL must have the protocol defined, it cannot be empty")
	// ErrInvalidURLHost is returned when host is empty
	ErrInvalidURLHost = errors.New("A valid host must be present")
	// ErrInvalidByteInput is returned when the input is not a valid byte slice
	ErrInvalidByteInput = errors.New("Invalid bytes as input")
	// ErrInvalidDateTimeInput is returned when the datetime format is not right
	ErrInvalidDateTimeInput = errors.New("A valid datetime format needed with timezone, e.g. 2015-02-07T13:28:17-05:00")
	// ErrInvalidDateInput is returned when the date format is not right
	ErrInvalidDateInput = errors.New("A valid date format needed with NO timezone, e.g. 2018, 1973-06, or 1905-08-23")
	// ErrInvalidTimeInput is returned when the time format is not right
	ErrInvalidTimeInput = errors.New("A valid time format needed with NO timezone, e.g. 13:28:17, 01:28:17PM")
	// ErrInvalidOIDInput is returned when the OID format is not right
	ErrInvalidOIDInput = errors.New("A valid OID is needed. It should start with `urn:oid:`, e.g. urn:oid:1.2.3.4.5")
	// ErrInvalidIDInput is returned when the ID format is not right
	ErrInvalidIDInput = errors.New("A valid ID required. 'A'..'Z', and 'a'..'z', numerals ('0'..'9'), '-' and '.', with a length limit of 64 characters")
	// ErrInvalidUnsignedIntRange is returned when the unsigned int is out of range
	ErrInvalidUnsignedIntRange = errors.New("Unsigned Integer value is out of range. Must be between 0 and 2,147,483,647")
	// ErrInvalidPositiveIntRange is returned when the
	ErrInvalidPositiveIntRange = errors.New("Positive Integer value is out of range. Must be between 1 and 2,147,483,647")
	// ErrInvalidUUIDInput is returned when the UUID is invalid
	ErrInvalidUUIDInput = errors.New("A valid UUID must be present, should start with `urn:uuid:`")
	// ErrInvalidUUIDFormat is returned when the UUID format is not right
	ErrInvalidUUIDFormat = errors.New("Invalid UUID format")
	// ErrInvalidUUIDGeneration is returned when the UUID generation failed
	ErrInvalidUUIDGeneration = errors.New("Unable to generate UUID")
)

// Boolean FHIR type
type Boolean bool

// NewBoolean returns a valid boolean
func NewBoolean(value interface{}) (Boolean, error) {
	v, ok := value.(bool)
	if !ok {
		return false, ErrInvalidBooleanInput
	}

	return Boolean(v), nil
}

// Integer FHIR type
type Integer int32

// NewInteger returns a valid integer
func NewInteger(value interface{}) (Integer, error) {
	v, ok := value.(int32)
	if !ok {
		return 0, ErrInvalidIntegerInput
	}

	if v < -2147483648 || v > 2147483647 {
		return 0, ErrInvalidIntegerRange
	}

	return Integer(v), nil
}

// String FHIR type
type String string

// NewString returns a valid string
func NewString(value interface{}) (String, error) {
	v, ok := value.(string)
	if !ok {
		return "", ErrInvalidStringInput
	}
	if len(v) > 1024*1024 {
		return "", ErrInvalidStringLength
	}

	return String(v), nil
}

// Decimal FHIR type
type Decimal float64

// NewDecimal returns a valid Decimal
func NewDecimal(value interface{}) (Decimal, error) {
	v, ok := value.(string)
	if !ok {
		return 0.0, ErrInvalidStringInput
	}
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0.0, ErrInvalidDecimalInput
	}
	if f > math.MaxFloat64 {
		return 0.0, ErrInvalidDecimalRange
	}

	return Decimal(f), nil
}

// URI FHIR type
type URI string

// NewURI returns a valid URI
func NewURI(value interface{}) (URI, error) {
	v, ok := value.(string)
	if !ok {
		return "", ErrInvalidStringInput
	}
	_, err := url.ParseRequestURI(v)
	if err != nil {
		return "", ErrInvalidURI
	}

	return URI(v), nil
}

// URL FHIR type
type URL URI

// NewURL returns a valid URL
func NewURL(value interface{}) (URL, error) {
	v, ok := value.(string)
	if !ok {
		return "", ErrInvalidStringInput
	}
	u, err := url.Parse(v)
	if err != nil {
		return "", ErrInvalidURL
	}
	if strings.TrimSpace(u.Scheme) == "" {
		return "", ErrInvalidURLScheme
	}
	if strings.TrimSpace(u.Host) == "" {
		return "", ErrInvalidURLHost
	}

	return URL(v), nil
}

// Canonical FHIR type
type Canonical URI

// NewCanonical returns a valid URL
func NewCanonical(value interface{}) (Canonical, error) {
	v, err := NewURI(value)
	if err != nil {
		return "", err
	}

	return Canonical(v), nil
}

// Base64Binary FHIR type
type Base64Binary string

// NewBase64Binary returns a valid Base64Binary
func NewBase64Binary(value interface{}) (Base64Binary, error) {
	v, ok := value.([]byte)
	if !ok {
		return "", ErrInvalidByteInput
	}
	encoded := base64.StdEncoding.EncodeToString(v)

	return Base64Binary(encoded), nil
}

// Instant FHIR type
type Instant string

// NewInstant returns a valid Instant
func NewInstant(value interface{}) (Instant, error) {
	v, ok := value.(string)
	if !ok {
		return Instant(time.Now().Format(time.RFC3339)), ErrInvalidStringInput
	}
	t, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return Instant(time.Now().Format(time.RFC3339)), ErrInvalidDateTimeInput
	}
	return Instant(t.Format(time.RFC3339)), nil
}

// Date FHIR type
type Date string

// NewDate returns a valid Date
func NewDate(value interface{}) (Date, error) {
	v, ok := value.(string)
	if !ok {
		return Date(v), ErrInvalidStringInput
	}
	_, err := time.Parse("2006", v)
	if err == nil {
		return Date(v), nil
	}
	_, err = time.Parse("2006-01", v)
	if err == nil {
		return Date(v), nil
	}
	_, err = time.Parse("2006-01-02", v)
	if err == nil {
		return Date(v), nil
	}

	return Date(v), ErrInvalidDateInput
}

// DateTime FHIR type
type DateTime string

// NewDateTime returns a valid DateTime
func NewDateTime(value interface{}) (DateTime, error) {
	dt, err := NewInstant(value)
	if err == nil {
		return DateTime(dt), nil
	}

	d, err := NewDate(value)
	if err == nil {
		return DateTime(d), nil
	}

	return "", err
}

// Time FHIR type
type Time string

// NewTime returns a valid Time
func NewTime(value interface{}) (Time, error) {
	v, ok := value.(string)
	if !ok {
		return Time(v), ErrInvalidStringInput
	}
	_, err := time.Parse("15:04:05", v)
	if err == nil {
		return Time(v), nil
	}
	t, err := time.Parse("03:04:05PM", v)
	if err == nil {
		return Time(t.Format("15:04:05")), nil
	}
	t, err = time.Parse("03:04:05 PM", v)
	if err == nil {
		return Time(t.Format("15:04:05")), nil
	}

	return Time(v), ErrInvalidTimeInput
}

// Code FHIR type
type Code String

// NewCode returns a valid Code
func NewCode(value interface{}) (Code, error) {
	v, err := NewString(value)
	if err != nil {
		return "", err
	}

	return Code(v), nil
}

// OID FHIR type
type OID URI

// NewOID returns a valid OID
func NewOID(value interface{}) (OID, error) {
	v, err := NewURI(value)
	if err != nil {
		return "", err
	}
	if !strings.HasPrefix(string(v), "urn:oid:") {
		return "", ErrInvalidOIDInput
	}

	return OID(v), nil
}

// ID FHIR type
type ID String

// NewID returns a valid ID
func NewID(value interface{}) (ID, error) {
	v, err := NewString(value)
	if err != nil {
		return "", err
	}
	var re = regexp.MustCompile(`[A-Za-z0-9\-\.]{1,64}`)
	if !re.MatchString(string(v)) {
		return "", ErrInvalidIDInput
	}
	return ID(v), nil
}

// Markdown FHIR type
type Markdown String

// NewMarkdown returns a valid Markdown
func NewMarkdown(value interface{}) (Markdown, error) {
	v, err := NewString(value)
	if err != nil {
		return "", err
	}

	return Markdown(v), nil
}

// UnsignedInt FHIR type
type UnsignedInt Integer

// NewUnsignedInt returns a valid UnsignedInt
func NewUnsignedInt(value interface{}) (UnsignedInt, error) {
	i, err := NewInteger(value)
	if err != nil {
		return 0, err
	}
	if i < 0 {
		return 0, ErrInvalidUnsignedIntRange
	}

	return UnsignedInt(i), nil
}

// PositiveInt FHIR type
type PositiveInt Integer

// NewPositiveInt returns a valid PositiveInt
func NewPositiveInt(value interface{}) (PositiveInt, error) {
	i, err := NewInteger(value)
	if err != nil {
		return 0, err
	}
	if i < 1 {
		return 0, ErrInvalidPositiveIntRange
	}

	return PositiveInt(i), nil
}

// UUID FHIR type
type UUID URI

// NewUUID returns a valid UUID
func NewUUID(value interface{}) (UUID, error) {
	if value != nil {
		v, err := NewURI(value)
		if err != nil {
			return "", err
		}
		if !strings.HasPrefix(string(v), "urn:uuid:") {
			return "", ErrInvalidUUIDInput
		}
		_, err = uuid.Parse(strings.TrimPrefix(string(v), "urn:uuid:"))
		if err != nil {
			return "", ErrInvalidUUIDFormat
		}
		return UUID(v), nil
	}

	u, err := uuid.NewUUID()
	if err != nil {
		return "", ErrInvalidUUIDGeneration
	}
	v := u.String()
	if !strings.HasPrefix(v, "urn:uuid:") {
		v = "urn:uuid:" + v
	}

	return UUID(v), nil
}
