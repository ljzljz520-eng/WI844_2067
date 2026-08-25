package audit

import (
	"coldchain/internal/domain"
	"testing"
)

func TestAuditSummary(t *testing.T) {
	m := Summary([]domain.DispatchEvent{{Type: "loaded"}, {Type: "loaded"}, {Type: "completed"}})
	if m["loaded"] != 2 || m["completed"] != 1 {
		t.Fatal(m)
	}
}
