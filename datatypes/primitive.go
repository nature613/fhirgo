package datatypes

import (
	"encoding/base64"
	"math"
	"net/url"
	"regexp"
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
	ErrInvalidIntegerRange = errors.New("Integer value is out of range")
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
		return "", ErrInvalidIntegerInput
	}
	//TODO: Error
	if len(v) > 1024*1024 {
		return "", ErrInvalidIntegerInput
	}

	return String(v), nil
}

// Decimal FHIR type
type Decimal float64

// NewDecimal returns a valid Decimal
func NewDecimal(value interface{}) (Decimal, error) {
	v, ok := value.(float64)
	if !ok {
		return 0.0, ErrInvalidIntegerInput
	}
	//TODO: Error
	if v > math.MaxFloat64 {
		return 0.0, ErrInvalidIntegerInput
	}

	return Decimal(v), nil
}

// URI FHIR type
type URI string

// NewURI returns a valid URI
func NewURI(value interface{}) (URI, error) {
	v, ok := value.(string)
	if !ok {
		return "", ErrInvalidIntegerInput
	}
	//TODO: Error
	_, err := url.ParseRequestURI(v)
	if err != nil {
		return "", ErrInvalidIntegerInput
	}

	return URI(v), nil
}

// URL FHIR type
type URL URI

// NewURL returns a valid URL
func NewURL(value interface{}) (URL, error) {
	v, ok := value.(string)
	if !ok {
		return "", ErrInvalidIntegerInput
	}
	//TODO: Error
	u, err := url.Parse(v)
	if err != nil {
		return "", ErrInvalidIntegerInput
	}
	if strings.TrimSpace(u.Scheme) == "" {
		return "", ErrInvalidIntegerInput
	}
	if strings.TrimSpace(u.Host) == "" {
		return "", ErrInvalidIntegerInput
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
		return "", ErrInvalidIntegerInput
	}
	//TODO: Error
	encoded := base64.StdEncoding.EncodeToString(v)

	return Base64Binary(encoded), nil
}

// Instant FHIR type
type Instant string

// NewInstant returns a valid Instant
func NewInstant(value interface{}) (Instant, error) {
	v, ok := value.(string)
	if !ok {
		return Instant(time.Now().Format(time.RFC3339)), ErrInvalidIntegerInput
	}
	//TODO: Error
	t, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return Instant(time.Now().Format(time.RFC3339)), ErrInvalidIntegerInput
	}
	return Instant(t.Format(time.RFC3339)), nil
}

// Date FHIR type
type Date string

// NewDate returns a valid Date
func NewDate(value interface{}) (Date, error) {
	v, ok := value.(string)
	if !ok {
		return Date(v), ErrInvalidIntegerInput
	}
	//TODO: Error
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

	return Date(v), ErrInvalidIntegerInput
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
		return Time(v), ErrInvalidIntegerInput
	}
	//TODO: Error
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

	return Time(v), ErrInvalidIntegerInput
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
	//TODO: Error
	if !strings.HasPrefix(string(v), "urn:oid:") {
		return "", err
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
	//TODO: Error
	if !re.MatchString(string(v)) {
		return "", err
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
	//TODO: Error
	if i < 0 {
		return 0, err
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
	//TODO: Error
	if i < 1 {
		return 0, err
	}

	return PositiveInt(i), nil
}

// UUID FHIR type
type UUID URI

// NewUUID returns a valid UUID
func NewUUID(value interface{}) (UUID, error) {
	//TODO: Error
	if value != nil {
		v, err := NewURI(value)
		if err != nil {
			return "", err
		}
		if !strings.HasPrefix(string(v), "urn:uuid:") {
			return "", err
		}
		_, err = uuid.Parse(strings.TrimPrefix(string(v), "urn:uuid:"))
		if err != nil {
			return "", err
		}
		return UUID(v), nil
	}

	u, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	v := u.String()
	if !strings.HasPrefix(v, "urn:uuid:") {
		v = "urn:uuid:" + v
	}

	return UUID(v), nil
}
