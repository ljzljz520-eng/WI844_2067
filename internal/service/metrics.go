package service

import "coldchain/internal/domain"

func eligible(v domain.Vehicle, d domain.DispatchOrder) bool {
	return v.Available && v.Zone == d.Zone && v.CapacityKg >= d.WeightKg
}
func choose(vs []domain.Vehicle, d domain.DispatchOrder) domain.Vehicle {
	best := domain.Vehicle{}
	for _, v := range vs {
		if eligible(v, d) && (best.ID == "" || v.CapacityKg < best.CapacityKg) {
			best = v
		}
	}
	return best
}
func totalWeight(ds []domain.DispatchOrder) int {
	n := 0
	for _, d := range ds {
		n += d.WeightKg
	}
	return n
}
func byStatus(ds []domain.DispatchOrder, st string) []domain.DispatchOrder {
	out := []domain.DispatchOrder{}
	for _, d := range ds {
		if d.Status == st {
			out = append(out, d)
		}
	}
	return out
}
func byZone(ds []domain.DispatchOrder, z string) []domain.DispatchOrder {
	out := []domain.DispatchOrder{}
	for _, d := range ds {
		if d.Zone == z {
			out = append(out, d)
		}
	}
	return out
}
func (s *Service) QueueMetrics() (map[string]int, error) {
	ds, e := s.Store.Dispatches()
	if e != nil {
		return nil, e
	}
	return map[string]int{"queued": len(byStatus(ds, domain.StatusQueued)), "loaded": len(byStatus(ds, domain.StatusLoaded)), "completed": len(byStatus(ds, domain.StatusCompleted)), "weight": totalWeight(ds)}, nil
}
func (s *Service) AvailableVehicles() ([]domain.Vehicle, error) {
	vs, e := s.Store.Vehicles()
	if e != nil {
		return nil, e
	}
	out := []domain.Vehicle{}
	for _, v := range vs {
		if v.Available {
			out = append(out, v)
		}
	}
	return out, nil
}
func (s *Service) MatchVehicle(d domain.DispatchOrder) (domain.Vehicle, error) {
	vs, e := s.Store.Vehicles()
	if e != nil {
		return domain.Vehicle{}, e
	}
	v := choose(vs, d)
	if v.ID == "" {
		return v, domain.ErrCapacity
	}
	return v, nil
}
func (s *Service) HasPark(id string) bool {
	xs, e := s.Store.Parks()
	if e != nil {
		return false
	}
	for _, p := range xs {
		if p.ID == id && p.Active {
			return true
		}
	}
	return false
}
