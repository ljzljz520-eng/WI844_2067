package coldchain

import "testing"

func TestWorkflowVehicleLoad(t *testing.T) {
	a, e := NewApp(t.TempDir() + "/x.db")
	if e != nil {
		t.Fatal(e)
	}
	defer a.Close()
	d, e := a.Create("P1", "chilled", 200)
	if e != nil {
		t.Fatal(e)
	}
	if e = a.S.Load(d.ID); e != nil {
		t.Fatal(e)
	}
}
