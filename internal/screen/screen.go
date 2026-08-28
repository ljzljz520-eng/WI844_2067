package screen

import (
	"coldchain/internal/domain"
	"sort"
)

type Row struct {
	ID, Zone, Status string
	Weight           int
}

func Project(ds []domain.DispatchOrder) []Row {
	rows := make([]Row, 0, len(ds))
	for _, d := range ds {
		rows = append(rows, Row{ID: d.ID, Zone: d.Zone, Status: d.Status, Weight: d.WeightKg})
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].ID < rows[j].ID })
	return rows
}
func Count(rows []Row, status string) int {
	n := 0
	for _, r := range rows {
		if r.Status == status {
			n++
		}
	}
	return n
}
