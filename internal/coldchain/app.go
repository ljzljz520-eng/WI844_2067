package coldchain

import (
	"coldchain/internal/domain"
	"coldchain/internal/parser"
	"coldchain/internal/service"
	"coldchain/internal/store"
)

type App struct{ S *service.Service }

func NewApp(path string) (*App, error) {
	st, e := store.Open(path)
	if e != nil {
		return nil, e
	}
	_ = st.PutPark(domain.Park{ID: "P1", Name: "Central", Timezone: "Asia/Shanghai", Active: true})
	_ = st.PutPark(domain.Park{ID: "P2", Name: "South", Timezone: "Asia/Shanghai", Active: true})
	_ = st.PutVehicle(domain.Vehicle{ID: "V1", Plate: "沪A-001", Zone: "chilled", CapacityKg: 1000, Available: true})
	return &App{S: service.New(st)}, nil
}
func (a *App) Close() error { return a.S.Store.Close() }
func (a *App) Create(park, zone string, weight int) (domain.DispatchOrder, error) {
	return a.S.CreateDispatch(parser.DispatchRequest{Park: park, Zone: zone, Weight: weight})
}
