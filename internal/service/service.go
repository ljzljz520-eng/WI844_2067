package service

import (
	"coldchain/internal/domain"
	"coldchain/internal/parser"
	"coldchain/internal/store"
	"fmt"
	"time"
)

type Service struct {
	Store *store.Store
	now   func() time.Time
}

func New(s *store.Store) *Service { return &Service{Store: s, now: time.Now} }
func (s *Service) CreateDispatch(req parser.DispatchRequest) (domain.DispatchOrder, error) {
	parks, e := s.Store.Parks()
	if e != nil {
		return domain.DispatchOrder{}, e
	}
	found := false
	for _, p := range parks {
		if p.ID == req.Park {
			found = p.Active
		}
	}
	if !found {
		return domain.DispatchOrder{}, domain.ErrInvalidPark
	}
	if !domain.ValidZone(req.Zone) {
		return domain.DispatchOrder{}, fmt.Errorf("invalid zone")
	}
	d := domain.DispatchOrder{ID: fmt.Sprintf("D-%d", s.now().UnixNano()), ParkID: req.Park, Zone: req.Zone, WeightKg: req.Weight, Status: domain.StatusQueued}
	return d, s.Store.PutDispatch(d)
}
func (s *Service) List(park, zone string) ([]domain.DispatchOrder, error) {
	ds, e := s.Store.Dispatches()
	if e != nil {
		return nil, e
	}
	out := []domain.DispatchOrder{}
	for _, d := range ds {
		if d.ParkID == park && (zone == "" || d.Zone == zone) {
			out = append(out, d)
		}
	}
	return out, nil
}
func (s *Service) Load(id string) error {
	ds, e := s.Store.Dispatches()
	if e != nil {
		return e
	}
	vs, e := s.Store.Vehicles()
	if e != nil {
		return e
	}
	for _, d := range ds {
		if d.ID == id {
			for _, v := range vs {
				if d.CanLoad(v) {
					if e = d.Load(v); e == nil {
						_ = s.Store.PutDispatch(d)
						_ = s.Store.PutEvent(domain.DispatchEvent{ID: "E-load-" + id, DispatchID: id, Type: "loaded", Actor: "system", At: s.now().Format(time.RFC3339)})
						return nil
					}
				}
			}
			return domain.ErrCapacity
		}
	}
	return domain.ErrNotFound
}
func (s *Service) Complete(id string) error {
	ds, e := s.Store.Dispatches()
	if e != nil {
		return e
	}
	for _, d := range ds {
		if d.ID == id {
			if e = d.Complete(); e != nil {
				return e
			}
			if e = s.Store.PutDispatch(d); e != nil {
				return e
			}
			return s.Store.PutEvent(domain.DispatchEvent{ID: "E-complete-" + id, DispatchID: id, Type: "completed", Actor: "system", At: s.now().Format(time.RFC3339)})
		}
	}
	return domain.ErrNotFound
}
