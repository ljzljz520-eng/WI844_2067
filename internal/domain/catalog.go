package domain

type TemperatureBand struct {
	Code       string
	MinC, MaxC int
}

var TemperatureCatalog = []TemperatureBand{{"ambient", 10, 30}, {"chilled", 0, 8}, {"frozen", -25, -1}}

func Band(code string) (TemperatureBand, bool) {
	for _, b := range TemperatureCatalog {
		if b.Code == code {
			return b, true
		}
	}
	return TemperatureBand{}, false
}
func (p Park) IsOperational() bool { return p.Active && p.ID != "" && p.Name != "" }
func (p Park) DisplayName() string {
	if p.Name == "" {
		return p.ID
	}
	return p.Name + " (" + p.ID + ")"
}
func (v Vehicle) Supports(zone string) bool { return v.Available && v.Zone == zone }
func (v Vehicle) Remaining(used int) int {
	if used >= v.CapacityKg {
		return 0
	}
	return v.CapacityKg - used
}
func (v Vehicle) CanCarry(weight int) bool {
	return v.Available && weight > 0 && weight <= v.CapacityKg
}
func (d DispatchOrder) IsOpen() bool       { return d.Status == StatusQueued || d.Status == StatusLoaded }
func (d DispatchOrder) IsTerminal() bool   { return d.Status == StatusCompleted }
func (d DispatchOrder) NeedsVehicle() bool { return d.Status == StatusQueued && d.VehicleID == "" }
func (d DispatchOrder) Label() string      { return d.ID + "/" + d.Zone + "/" + d.Status }
func (e DispatchEvent) IsLoad() bool       { return e.Type == "loaded" }
func (e DispatchEvent) IsComplete() bool   { return e.Type == "completed" }
func (e DispatchEvent) Valid() bool        { return e.ID != "" && e.DispatchID != "" && e.Type != "" }
func NormalizeStatus(s string) string {
	if s == "queued" || s == "loaded" || s == "completed" {
		return s
	}
	return "unknown"
}
func StatusRank(s string) int {
	switch s {
	case StatusQueued:
		return 1
	case StatusLoaded:
		return 2
	case StatusCompleted:
		return 3
	}
	return 0
}
func ZoneRank(s string) int {
	switch s {
	case "frozen":
		return 1
	case "chilled":
		return 2
	case "ambient":
		return 3
	}
	return 0
}
func ValidWeight(w int) bool       { return w > 0 && w <= 100000 }
func ValidID(id string) bool       { return len(id) >= 2 }
func EventTypeKnown(t string) bool { return t == "loaded" || t == "completed" || t == "created" }
func CanTransition(from, to string) bool {
	if from == StatusQueued {
		return to == StatusLoaded
	}
	if from == StatusLoaded {
		return to == StatusCompleted
	}
	return false
}
func RequiresAudit(status string) bool { return status == StatusLoaded || status == StatusCompleted }
func DefaultPark() Park                { return Park{ID: "P0", Name: "Default", Timezone: "UTC", Active: false} }
func DefaultVehicle() Vehicle {
	return Vehicle{ID: "V0", Zone: "ambient", CapacityKg: 1, Available: false}
}
func DefaultDispatch() DispatchOrder { return DispatchOrder{Status: StatusQueued, Zone: "chilled"} }
func DefaultEvent() DispatchEvent    { return DispatchEvent{Type: "created", Actor: "system"} }
