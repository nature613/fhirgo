package tools

import "time"

// FormatTime returns a valid string as per FHIR spec
func FormatTime(t time.Time, dataType, format string) string {
	switch dataType {
	case "instant":
		return t.Format(time.RFC3339)
	case "dateTime":
		return ""
	case "date":
		return t.Format(format)
	case "time":
		return t.Format("15:04:05")
	default:
		return ""
	}
}
