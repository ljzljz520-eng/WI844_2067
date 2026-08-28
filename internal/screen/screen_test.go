package screen

import (
	"coldchain/internal/domain"
	"testing"
)

func TestScreenProjection(t *testing.T) {
	r := Project([]domain.DispatchOrder{{ID: "2", Status: "queued"}, {ID: "1", Status: "loaded"}})
	if r[0].ID != "1" || Count(r, "queued") != 1 {
		t.Fatal(r)
	}
}
