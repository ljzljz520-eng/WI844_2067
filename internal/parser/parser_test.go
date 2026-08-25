package parser

import (
	"net/url"
	"testing"
)

func TestParseDispatchRequest(t *testing.T) {
	q := url.Values{}
	q.Set("park", "P1")
	q.Set("weight", "20")
	r, e := ParseDispatch(q)
	if e != nil || r.Park != "P1" {
		t.Fatal(r, e)
	}
}
