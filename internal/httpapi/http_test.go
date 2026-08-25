package httpapi

import (
	"coldchain/internal/coldchain"
	"net/http/httptest"
	"testing"
)

func TestHTTPInvalidPark(t *testing.T) {
	a, e := coldchain.NewApp(t.TempDir() + "/x")
	if e != nil {
		t.Fatal(e)
	}
	defer a.Close()
	r := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/v1/parks/bad%20park/dispatches", nil)
	New(a.S).ServeHTTP(r, req)
	if r.Code != 400 {
		t.Fatalf("code %d", r.Code)
	}
}
