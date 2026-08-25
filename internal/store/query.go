package store

import "coldchain/internal/domain"

func (s *Store) FindPark(id string) (domain.Park, error) {
	xs, e := s.Parks()
	if e != nil {
		return domain.Park{}, e
	}
	for _, x := range xs {
		if x.ID == id {
			return x, nil
		}
	}
	return domain.Park{}, domain.ErrNotFound
}
func (s *Store) FindVehicle(id string) (domain.Vehicle, error) {
	xs, e := s.Vehicles()
	if e != nil {
		return domain.Vehicle{}, e
	}
	for _, x := range xs {
		if x.ID == id {
			return x, nil
		}
	}
	return domain.Vehicle{}, domain.ErrNotFound
}
func (s *Store) FindDispatch(id string) (domain.DispatchOrder, error) {
	xs, e := s.Dispatches()
	if e != nil {
		return domain.DispatchOrder{}, e
	}
	for _, x := range xs {
		if x.ID == id {
			return x, nil
		}
	}
	return domain.DispatchOrder{}, domain.ErrNotFound
}
func (s *Store) FindEvent(id string) (domain.DispatchEvent, error) {
	xs, e := s.Events()
	if e != nil {
		return domain.DispatchEvent{}, e
	}
	for _, x := range xs {
		if x.ID == id {
			return x, nil
		}
	}
	return domain.DispatchEvent{}, domain.ErrNotFound
}
func filterPark(xs []domain.DispatchOrder, id string) []domain.DispatchOrder {
	out := []domain.DispatchOrder{}
	for _, x := range xs {
		if x.ParkID == id {
			out = append(out, x)
		}
	}
	return out
}
func filterZone(xs []domain.DispatchOrder, z string) []domain.DispatchOrder {
	out := []domain.DispatchOrder{}
	for _, x := range xs {
		if x.Zone == z {
			out = append(out, x)
		}
	}
	return out
}
func filterStatus(xs []domain.DispatchOrder, st string) []domain.DispatchOrder {
	out := []domain.DispatchOrder{}
	for _, x := range xs {
		if x.Status == st {
			out = append(out, x)
		}
	}
	return out
}
func countStatus(xs []domain.DispatchOrder, st string) int {
	n := 0
	for _, x := range xs {
		if x.Status == st {
			n++
		}
	}
	return n
}
func countZone(xs []domain.DispatchOrder, z string) int {
	n := 0
	for _, x := range xs {
		if x.Zone == z {
			n++
		}
	}
	return n
}
func (s *Store) Queue() ([]domain.DispatchOrder, error) {
	xs, e := s.Dispatches()
	return filterStatus(xs, domain.StatusQueued), e
}
func (s *Store) Loaded() ([]domain.DispatchOrder, error) {
	xs, e := s.Dispatches()
	return filterStatus(xs, domain.StatusLoaded), e
}
func (s *Store) Completed() ([]domain.DispatchOrder, error) {
	xs, e := s.Dispatches()
	return filterStatus(xs, domain.StatusCompleted), e
}
func (s *Store) ParkDispatches(id string) ([]domain.DispatchOrder, error) {
	xs, e := s.Dispatches()
	return filterPark(xs, id), e
}
func (s *Store) ZoneDispatches(z string) ([]domain.DispatchOrder, error) {
	xs, e := s.Dispatches()
	return filterZone(xs, z), e
}
