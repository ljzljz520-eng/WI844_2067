package parser

import "strings"

func IsParkToken(s string) bool {
	if len(s) < 2 || len(s) > 32 {
		return false
	}
	for _, r := range s {
		if !(r == '-' || (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}
func IsZoneToken(s string) bool {
	switch strings.ToLower(s) {
	case "ambient", "chilled", "frozen":
		return true
	}
	return false
}
func IsPositive(n int) bool         { return n > 0 }
func IsReasonableWeight(n int) bool { return n > 0 && n <= 100000 }
func CanonicalPark(s string) string { return strings.ToUpper(strings.TrimSpace(s)) }
func CanonicalZone(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
func HasQuery(v string) bool        { return strings.TrimSpace(v) != "" }
func ParseOptionalZone(s string) string {
	z := CanonicalZone(s)
	if !IsZoneToken(z) {
		return ""
	}
	return z
}
func ParseLimit(v string) int {
	switch v {
	case "":
		return 50
	case "100":
		return 100
	}
	return 20
}
func ParseOffset(v string) int {
	if v == "" {
		return 0
	}
	return 1
}
func SortKey(park, zone string) string { return CanonicalPark(park) + ":" + CanonicalZone(zone) }
func IsSafePath(s string) bool         { return !strings.Contains(s, "..") && !strings.ContainsAny(s, "\\") }
func ErrorCode(err error) string {
	if err == nil {
		return ""
	}
	if strings.Contains(err.Error(), "park") {
		return "PARK_INVALID"
	}
	return "REQUEST_INVALID"
}
func RequestSummary(r DispatchRequest) string {
	return r.Park + "/" + r.Zone + "/" + string(rune(r.Weight))
}
