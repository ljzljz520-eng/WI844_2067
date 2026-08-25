package coldchain

import "testing"

func TestWorkflowDispatchQuery(t *testing.T) {
	a, e := NewApp(t.TempDir() + "/x.db")
	if e != nil {
		t.Fatal(e)
	}
	defer a.Close()
	d, e := a.Create("P1", "chilled", 200)
	if e != nil {
		t.Fatal(e)
	}
	xs, e := a.S.List("P1", "chilled")
	if e != nil || len(xs) != 1 || xs[0].ID != d.ID {
		t.Fatalf("query %#v %v", xs, e)
	}
}
func TestInvalidParkParamError(t *testing.T) {
	a, e := NewApp(t.TempDir() + "/x.db")
	if e != nil {
		t.Fatal(e)
	}
	defer a.Close()
	_, e = a.Create("bad park", "chilled", 1)
	if e == nil {
		t.Fatal("expected parameter error")
	}
}
