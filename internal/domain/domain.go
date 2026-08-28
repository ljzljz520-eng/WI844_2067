package domain

import "errors"

type Park struct {
	ID, Name, Timezone string
	Active             bool
}
type Vehicle struct {
	ID, Plate, Zone string
	CapacityKg      int
	Available       bool
}
type DispatchOrder struct {
	ID, ParkID, VehicleID, Zone, Status string
	WeightKg                            int
}
type DispatchEvent struct{ ID, DispatchID, Type, Actor, At string }

const (
	StatusQueued    = "queued"
	StatusLoaded    = "loaded"
	StatusCompleted = "completed"
)

var ErrInvalidPark = errors.New("invalid park parameter")
var ErrNotFound = errors.New("entity not found")
var ErrCapacity = errors.New("vehicle capacity or zone mismatch")

func (d DispatchOrder) CanLoad(v Vehicle) bool {
	return d.Status == StatusQueued && v.Available && v.Zone == d.Zone && v.CapacityKg >= d.WeightKg
}
func (d *DispatchOrder) Load(v Vehicle) error {
	if !d.CanLoad(v) {
		return ErrCapacity
	}
	d.VehicleID = v.ID
	d.Status = StatusLoaded
	return nil
}
func (d *DispatchOrder) Complete() error {
	if d.Status != StatusLoaded {
		return errors.New("dispatch is not loaded")
	}
	d.Status = StatusCompleted
	return nil
}
func ValidZone(z string) bool { return z == "ambient" || z == "chilled" || z == "frozen" }
