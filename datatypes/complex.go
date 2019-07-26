package datatypes

// Attachment FHIR type
type Attachment struct {
	ContentType Code         `json:"contentType"`
	Language    Code         `json:"language"`
	Data        Base64Binary `json:"data"`
	URL         URL          `json:"url"`
	Size        UnsignedInt  `json:"size"`
	Hash        Base64Binary `json:"hash"`
	Title       String       `json:"title"`
	Creation    DateTime     `json:"creation"`
}

// Coding FHIR type
type Coding struct {
	System       URI     `json:"system"`
	Version      String  `json:"version"`
	Code         Code    `json:"code"`
	Display      String  `json:"display"`
	UserSelected Boolean `json:"userSelected"`
}

// CodeableConcept FHIR Type
type CodeableConcept struct {
	Coding []Coding `json:"coding"`
	Text   String   `json:"text"`
}

// Quantity FHIR Type
type Quantity struct {
	Value      Decimal `json:"value"`
	Comparator Code    `json:"comparator"`
	Unit       String  `json:"unit"`
	System     URI     `json:"system"`
	Code       Code    `json:"code"`
}

// Distance FHIR Type (Variations on Quantity)
type Distance Quantity

// Age FHIR Type (Variations on Quantity)
type Age Quantity

// Count FHIR Type (Variations on Quantity)
type Count Quantity

// Duration FHIR Type (Variations on Quantity)
type Duration Quantity

// Money FHIR Type
type Money struct {
	Value    Decimal `json:"value"`
	Currency Code    `json:"currency"`
}

// Range FHIR Type
type Range struct {
	Low  Quantity `json:"low"`
	High Quantity `json:"high"`
}

// Ratio FHIR Type
type Ratio struct {
	Numerator   Quantity `json:"numerator"`
	Denominator Quantity `json:"denominator"`
}

// Period FHIR Type
type Period struct {
	Start DateTime `json:"start"`
	End   DateTime `json:"end"`
}

// SampledData FHIR Type
type SampledData struct {
	Origin     Quantity    `json:"origin"`
	Period     Decimal     `json:"period"`
	Factor     Decimal     `json:"factor"`
	LowerLimit Decimal     `json:"lowerLimit"`
	UpperLimit Decimal     `json:"upperLimit"`
	Dimensions PositiveInt `json:"dimensions"`
	Data       String      `json:"data"`
}

// Identifier FHIR Type
type Identifier struct {
	Use    Code            `json:"use"`
	Type   CodeableConcept `json:"type"`
	System URI             `json:"system"`
	Value  String          `json:"value"`
	Period Period          `json:"period"`
	// Assigner Reference       `json:"assigner"`
}

// HumanName FHIR Type
type HumanName struct {
	Use    Code     `json:"use"`
	Text   String   `json:"text"`
	Family String   `json:"family"`
	Given  []String `json:"given"`
	Prefix []String `json:"prefix"`
	Suffix []String `json:"suffix"`
	Period Period   `json:"period"`
}

// Address FHIR Type
type Address struct {
	Use        Code     `json:"use"`
	Type       Code     `json:"type"`
	Text       String   `json:"text"`
	Line       []String `json:"line"`
	City       String   `json:"city"`
	District   String   `json:"district"`
	State      String   `json:"state"`
	PostalCode String   `json:"postalCode"`
	Country    String   `json:"country"`
	Period     Period   `json:"period"`
}

// ContactPoint FHIR Type
type ContactPoint struct {
	System Code        `json:"system"`
	Value  String      `json:"value"`
	Use    Code        `json:"use"`
	Rank   PositiveInt `json:"rank"`
	Period Period      `json:"period"`
}

// Repeat Type for Timing
type Repeat struct {
	BoundsDuration Duration    `json:"boundsDuration"`
	BoundsRange    Range       `json:"boundsRange"`
	BoundsPeriod   Period      `json:"boundsPeriod"`
	Count          PositiveInt `json:"count"`
	CountMax       PositiveInt `json:"countMax"`
	Duration       Decimal     `json:"duration"`
	DurationMax    Decimal     `json:"durationMax"`
	DurationUnit   Code        `json:"durationUnit"`
	Frequency      PositiveInt `json:"frequency"`
	FrequencyMax   PositiveInt `json:"frequencyMax"`
	Period         Decimal     `json:"period"`
	PeriodMax      Decimal     `json:"periodMax"`
	PeriodUnit     Code        `json:"periodUnit"`
	DayOfWeek      []Code      `json:"dayOfWeek"`
	TimeOfDay      []Time      `json:"timeOfDay"`
	When           []Code      `json:"when"`
	Offset         UnsignedInt `json:"offset"`
}

// Timing FHIR Type
type Timing struct {
	Event  []DateTime      `json:"event"`
	Repeat Repeat          `json:"repeat"`
	Code   CodeableConcept `json:"code"`
}

// Signature FHIR Type
type Signature struct {
	Type         []Coding     `json:"rype"`
	When         Instant      `json:"when"`
	Who          Reference    `json:"who"`
	OnBehalfOf   Reference    `json:"onBehalfOf"`
	TargetFormat Code         `json:"targetFormat"`
	SigFormat    Code         `json:"sigFormat"`
	Data         Base64Binary `json:"data"`
}

// Annotation FHIR Type
type Annotation struct {
	AuthorReference Reference `json:"authorReference"`
	AuthorString    String    `json:"authorString"`
	Time            DateTime  `json:"time"`
	Text            Markdown  `json:"text"`
}
