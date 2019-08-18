package datatypes

// Attachment FHIR type
type Attachment struct {
	Element
	ContentType *Code         `json:"contentType,omitempty"`
	Language    *Code         `json:"language,omitempty"`
	Data        *Base64Binary `json:"data,omitempty"`
	URL         *URL          `json:"url,omitempty"`
	Size        *UnsignedInt  `json:"size,omitempty"`
	Hash        *Base64Binary `json:"hash,omitempty"`
	Title       *String       `json:"title,omitempty"`
	Creation    *DateTime     `json:"creation,omitempty"`
}

// Coding FHIR type
type Coding struct {
	Element
	System       *URI     `json:"system,omitempty"`
	Version      *String  `json:"version,omitempty"`
	Code         *Code    `json:"code,omitempty"`
	Display      *String  `json:"display,omitempty"`
	UserSelected *Boolean `json:"userSelected,omitempty"`
}

// CodeableConcept FHIR Type
type CodeableConcept struct {
	Element
	Coding []Coding `json:"coding,omitempty"`
	Text   *String  `json:"text,omitempty"`
}

// Quantity FHIR Type
type Quantity struct {
	Element
	Value      *Decimal `json:"value,omitempty"`
	Comparator *Code    `json:"comparator,omitempty"`
	Unit       *String  `json:"unit,omitempty"`
	System     *URI     `json:"system,omitempty"`
	Code       *Code    `json:"code,omitempty"`
}

// Distance FHIR Type (Variations on Quantity)
type Distance Quantity

// Age FHIR Type (Variations on Quantity)
type Age Quantity

// Count FHIR Type (Variations on Quantity)
type Count Quantity

// Duration FHIR Type (Variations on Quantity)
type Duration Quantity

// MoneyQuantity FHIT Type  (Variations on Quantity)
type MoneyQuantity Quantity

// SimpleQuantity FHIT Type  (Variations on Quantity)
type SimpleQuantity Quantity

// Money FHIR Type
type Money struct {
	Element
	Value    *Decimal `json:"value,omitempty"`
	Currency *Code    `json:"currency,omitempty"`
}

// Range FHIR Type
type Range struct {
	Element
	Low  *Quantity `json:"low,omitempty"`
	High *Quantity `json:"high,omitempty"`
}

// Ratio FHIR Type
type Ratio struct {
	Element
	Numerator   *Quantity `json:"numerator,omitempty"`
	Denominator *Quantity `json:"denominator,omitempty"`
}

// Period FHIR Type
type Period struct {
	Element
	Start *DateTime `json:"start,omitempty"`
	End   *DateTime `json:"end,omitempty"`
}

// SampledData FHIR Type
type SampledData struct {
	Element
	Origin     *Quantity    `json:"origin,omitempty"`
	Period     *Decimal     `json:"period,omitempty"`
	Factor     *Decimal     `json:"factor,omitempty"`
	LowerLimit *Decimal     `json:"lowerLimit,omitempty"`
	UpperLimit *Decimal     `json:"upperLimit,omitempty"`
	Dimensions *PositiveInt `json:"dimensions,omitempty"`
	Data       *String      `json:"data,omitempty"`
}

// Identifier FHIR Type
type Identifier struct {
	Element
	Use    *Code            `json:"use,omitempty"`
	Type   *CodeableConcept `json:"type,omitempty"`
	System *URI             `json:"system,omitempty"`
	Value  *String          `json:"value,omitempty"`
	Period *Period          `json:"period,omitempty"`
	// Assigner *Reference       `json:"assigner,omitempty"`
}

// HumanName FHIR Type
type HumanName struct {
	Element
	Use    *Code    `json:"use,omitempty"`
	Text   *String  `json:"text,omitempty"`
	Family *String  `json:"family,omitempty"`
	Given  []String `json:"given,omitempty"`
	Prefix []String `json:"prefix,omitempty"`
	Suffix []String `json:"suffix,omitempty"`
	Period *Period  `json:"period,omitempty"`
}

// Address FHIR Type
type Address struct {
	Element
	Use        *Code    `json:"use,omitempty"`
	Type       *Code    `json:"type,omitempty"`
	Text       *String  `json:"text,omitempty"`
	Line       []String `json:"line,omitempty"`
	City       *String  `json:"city,omitempty"`
	District   *String  `json:"district,omitempty"`
	State      *String  `json:"state,omitempty"`
	PostalCode *String  `json:"postalCode,omitempty"`
	Country    *String  `json:"country,omitempty"`
	Period     *Period  `json:"period,omitempty"`
}

// ContactPoint FHIR Type
type ContactPoint struct {
	Element
	System *Code        `json:"system,omitempty"`
	Value  *String      `json:"value,omitempty"`
	Use    *Code        `json:"use,omitempty"`
	Rank   *PositiveInt `json:"rank,omitempty"`
	Period *Period      `json:"period,omitempty"`
}

// Repeat Type for Timing
type Repeat struct {
	BackboneElement
	BoundsDuration *Duration    `json:"boundsDuration,omitempty"`
	BoundsRange    *Range       `json:"boundsRange,omitempty"`
	BoundsPeriod   *Period      `json:"boundsPeriod,omitempty"`
	Count          *PositiveInt `json:"count,omitempty"`
	CountMax       *PositiveInt `json:"countMax,omitempty"`
	Duration       *Decimal     `json:"duration,omitempty"`
	DurationMax    *Decimal     `json:"durationMax,omitempty"`
	DurationUnit   *Code        `json:"durationUnit,omitempty"`
	Frequency      *PositiveInt `json:"frequency,omitempty"`
	FrequencyMax   *PositiveInt `json:"frequencyMax,omitempty"`
	Period         *Decimal     `json:"period,omitempty"`
	PeriodMax      *Decimal     `json:"periodMax,omitempty"`
	PeriodUnit     *Code        `json:"periodUnit,omitempty"`
	DayOfWeek      []Code       `json:"dayOfWeek,omitempty"`
	TimeOfDay      []Time       `json:"timeOfDay,omitempty"`
	When           []Code       `json:"when,omitempty"`
	Offset         *UnsignedInt `json:"offset,omitempty"`
}

// Timing FHIR Type
type Timing struct {
	Element
	Event  []DateTime       `json:"event,omitempty"`
	Repeat *Repeat          `json:"repeat,omitempty"`
	Code   *CodeableConcept `json:"code,omitempty"`
}

// Signature FHIR Type
type Signature struct {
	Element
	Type         []Coding      `json:"type,omitempty"`
	When         *Instant      `json:"when,omitempty"`
	Who          *Reference    `json:"who,omitempty"`
	OnBehalfOf   *Reference    `json:"onBehalfOf,omitempty"`
	TargetFormat *Code         `json:"targetFormat,omitempty"`
	SigFormat    *Code         `json:"sigFormat,omitempty"`
	Data         *Base64Binary `json:"data,omitempty"`
}

// Annotation FHIR Type
type Annotation struct {
	Element
	AuthorReference *Reference `json:"authorReference,omitempty"`
	AuthorString    *String    `json:"authorString,omitempty"`
	Time            *DateTime  `json:"time,omitempty"`
	Text            *Markdown  `json:"text,omitempty"`
}
