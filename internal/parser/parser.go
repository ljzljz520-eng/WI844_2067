package parser

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type DispatchRequest struct {
	Park, Zone string
	Weight     int
}

func ParsePark(raw string) (string, error) {
	p := strings.TrimSpace(raw)
	if p == "" || strings.ContainsAny(p, " /") {
		return "", fmt.Errorf("invalid park parameter: %w", fmt.Errorf("malformed"))
	}
	return p, nil
}
func ParseDispatch(q url.Values) (DispatchRequest, error) {
	p, e := ParsePark(q.Get("park"))
	if e != nil {
		return DispatchRequest{}, e
	}
	z := q.Get("zone")
	if z == "" {
		z = "chilled"
	}
	w, err := strconv.Atoi(q.Get("weight"))
	if err != nil || w <= 0 {
		return DispatchRequest{}, fmt.Errorf("invalid weight")
	}
	return DispatchRequest{Park: p, Zone: z, Weight: w}, nil
}
func NormalizeZone(z string) string { return strings.ToLower(strings.TrimSpace(z)) }
