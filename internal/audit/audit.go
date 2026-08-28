package audit

import (
	"coldchain/internal/domain"
	"sort"
)

func Timeline(events []domain.DispatchEvent, id string) []domain.DispatchEvent {
	out := []domain.DispatchEvent{}
	for _, e := range events {
		if e.DispatchID == id {
			out = append(out, e)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].At < out[j].At })
	return out
}
func Summary(events []domain.DispatchEvent) map[string]int {
	m := map[string]int{}
	for _, e := range events {
		m[e.Type]++
	}
	return m
}
func Latest(events []domain.DispatchEvent) domain.DispatchEvent {
	if len(events) == 0 {
		return domain.DispatchEvent{}
	}
	sort.Slice(events, func(i, j int) bool { return events[i].At > events[j].At })
	return events[0]
}
