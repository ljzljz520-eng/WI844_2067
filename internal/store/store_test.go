package store

import (
	"coldchain/internal/domain"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := t.TempDir() + "/db"
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.PutPark(domain.Park{ID: "P", Name: "Persist", Active: true}); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	xs, e := s.Parks()
	if e != nil || len(xs) != 1 || xs[0].Name != "Persist" {
		t.Fatalf("%v %#v", e, xs)
	}
}
